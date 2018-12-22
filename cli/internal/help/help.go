package help

import "fmt"

var helpText = `
Stuff is your team's personal search engine.

Usage:

	stuff command [arguments]

The commands are:

	index     store a document in the search engine
	search    search all documents for relevant information
	download  download a file with the given document id to the current directory

Use "stuff help [command]" for more information about a command.
`

// ShowHelp writes helpText to the screen
func ShowHelp() {
	fmt.Printf("%v\n", helpText)
}
