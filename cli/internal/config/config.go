package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

const (
	// storageDir is the root folder to store cfg, indexes, and files given the storagePath
	storageDir = ".starship"

	// configFileName is the name of the starship cli config file
	configFileName = "starship.conf"
)

// errStarshipAPINotSet is returned from NewConfigManager if the API url does not exist in config
var errStarshipAPINotSet = errors.New("the starship API to connect to is not set")

// Cfg represents the config for the starship cli
type Cfg struct {
	APIURL string `toml:"api"`
}

// NewConfigManager instantiates a new instance of Cfg or loads it if it already exists
func NewConfigManager() (*Cfg, error) {
	var cfg Cfg

	// create a folder at storage dir if it does not exist
	u, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "could not get user running app")
	}

	starshipPath := filepath.Join(u.HomeDir, storageDir)
	err = os.MkdirAll(starshipPath, os.ModePerm)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create config path at %v", starshipPath)
	}

	// create a config file if it does not exist
	configPath := filepath.Join(u.HomeDir, storageDir, configFileName)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Print(" Provide the full url for the starship API that you want to connect to: ")

		// ask for api to connect to
		var input string
		fmt.Scanln(&input)
		cfg.APIURL = input

		c, err := CreateNewConfigFile(&cfg, configPath)
		if err != nil {
			return nil, err
		}

		cfg = *c
	}

	// read config file
	c, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.Errorf("could not read config file at path %v", configPath)
	}

	// decode config file to struct
	if _, err := toml.Decode(string(c), &cfg); err != nil {
		return nil, errors.Errorf("could not decode config file at path %v. is it in a valid format?", configPath)
	}

	// validate api url in config file
	if err := validateAPIURL(cfg.APIURL); err != nil {
		fmt.Print(" Starship API URL retrieved from config is invalid. \n Provide the full url for the starship API that you want to connect to: ")

		// ask for api to connect to
		var input string
		fmt.Scanln(&input)
		cfg.APIURL = input

		c, err := CreateNewConfigFile(&cfg, configPath)
		if err != nil {
			return nil, err
		}

		cfg = *c
	}

	return &cfg, nil
}

// CreateNewConfigFile creates a new config file at the specified path and validates it
func CreateNewConfigFile(cfg *Cfg, configPath string) (*Cfg, error) {
	// validate api url provided
	if err := validateAPIURL(cfg.APIURL); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(cfg); err != nil {
		return nil, errors.Wrap(err, "could not encode config to toml")
	}

	err := ioutil.WriteFile(configPath, buf.Bytes(), os.ModePerm)
	if err != nil {
		return nil, errors.Wrapf(err, "could not write config file at path %v", configPath)
	}

	return cfg, nil
}

func validateAPIURL(u string) error {
	_, err := url.Parse(u)
	if err != nil {
		return errors.Wrap(err, "url provided is not valid, protocol and hostname are required")
	}

	if strings.Index(u, "https://") == -1 && strings.Index(u, "http://") == -1 {
		return errors.New("url provided is not valid, protocol and hostname are required")
	}

	return nil
}
