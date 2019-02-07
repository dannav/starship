package main

import (
	"flag"
	"net/http"
	"strings"
	"time"

	"github.com/dannav/starship/cli/internal/command"
	"github.com/dannav/starship/cli/internal/config"
	"github.com/dannav/starship/cli/internal/engine"
	"github.com/dannav/starship/cli/internal/help"
	"github.com/dannav/starship/cli/internal/log"
	"github.com/dannav/starship/cli/internal/uerrors"
)

// general cfg
var (
	// client is an http client with a timeout duration of 1 min
	client = &http.Client{
		Timeout: time.Second * 60,
	}

	// cmd is the command that is going to be run
	cmd string

	// args are any arguments passed through the cli for cmd
	args []string
)

// cli flags
var (
	// showHelp is a flag to determine whether help text should be shown
	showHelp bool

	// indexPath is a flag that denotes what path to index content at for a document
	indexPath string

	// forceYes is a flag that forces a yes input for any prompts
	forceYes bool
)

// init runs first, it parses any flags and arguments for this cli app
func init() {
	flag.BoolVar(&showHelp, "h", false, "show help text")
	flag.BoolVar(&showHelp, "help", false, "show help text")
	flag.BoolVar(&forceYes, "y", false, "force yes on prompt")
	flag.BoolVar(&forceYes, "yes", false, "force yes on prompt")
	flag.StringVar(&indexPath, "p", "/", "path to store document in index")
	flag.Parse()

	args = flag.Args()
}

func main() {
	// if help flags set or starship is run without any commands show help text
	if showHelp || len(args) == 0 {
		help.ShowHelp()
		return
	}

	// parse command from arguments
	args = args[1:]
	cmd = strings.ToLower(flag.Arg(0))

	// check to see if the starship API is up and functioning
	e := engine.NewEngine(client)

	// load or set config if we need it for the command we are running
	if cmd == "index" || cmd == "search" {
		cfg, err := config.NewConfigManager()
		if err != nil {
			log.Error(err.Error())
			return
		}
		e.APIEndpoint = cfg.APIURL

		ready, err := e.Ready()
		if ready != true {
			if err != nil {
				log.Error(err.Error())
			}

			log.Error(uerrors.APINotAvailable)
			return
		}
	}

	// cli main logic
	switch cmd {
	case "help":
		var subcommand string
		if len(args) > 0 {
			subcommand = args[0]
		}

		switch subcommand {
		case "index":
			help.ShowIndex()
			return
		case "search":
			help.ShowSearch()
			return
		}

		help.ShowHelp()
		return
	case "index":
		// index file passed in argument
		if err := command.Index(args, indexPath, e, forceYes); err != nil {
			log.Error(err.Error())
			return
		}
	case "search":
		// perform search on words passed as arguments
		if err := command.Search(args, e); err != nil {
			log.Error(err.Error())
			return
		}
	default:
		help.ShowHelp()
		return
	}
}
