// Package uerrors are all errors that are to be displayed to a user
package uerrors

const (
	// IndexPathNotProvided represents and error where a path is not provided as an argument to the index command
	IndexPathNotProvided = "Please provide the path to a file to index."

	// IndexFileDoesNotExist represents an error where the file does not exist at the path provided for the index command
	IndexFileDoesNotExist = "The file at the path provided does not exist."

	// SearchNoTextGiven represents an error where no search text was given for the search command
	SearchNoTextGiven = "Please provide some text that you would like to perform a search with."

	// SearchRequestFailed represents an error where a request to the search api failed
	SearchRequestFailed = "There was an issue performing your search."

	// SearchNoResults represents an error where there were no search results
	SearchNoResults = "No search results were found that matched your query."

	// APINotAvailable represents an error where the API is currently down
	APINotAvailable = "The starship API is currently experiencing maintenance and will be back shortly."
)
