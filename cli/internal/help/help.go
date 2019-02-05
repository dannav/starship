package help

import "fmt"

var helpText = `
Starship is your team's personal search engine.

Usage:

	starship command [arguments]

The commands are:

	index     store a document in the search engine
	search    search all documents for relevant information

Use "starship help [command]" for more information about a command.
`

var indexText = `
Usage: starship [-p path to store in search index] index [path to file]

The Index command allows you to store a document in your teams search index for
retrieval using the search command.

Think of the starship search index as a logical directory of where files are stored.

The -p flag defines where you would like to store the document on starship. All document
paths are unique and if the -p flag is not set the path will default to the root of
the search index.

Example: starship -p company index ./path/to/employee-handbook.pdf
`

var searchText = `
Usage: starship search [query]

The search command allows you to search all documents indexed with starship for the given
text provided. Search does more than just look for matching keywords when returning results.
Starship makes an attempt to understand what you typed and may return documents that are
similar in meaning.

For example using the search query in the example below (paid time off), starship might return
a list of holidays and hours found in documents indexed previously.

By default no search command will return empty results. Starship will try to return the best
results it can find based on documents previously indexed by you.

When searching for documents you will be presented with commands to view the next result (n),
previous result (p), quit (q), or download the file to your current directory.

Example: starship search "paid time off"
`

// ShowHelp writes helpText to the screen
func ShowHelp() {
	fmt.Printf("%v\n", helpText)
}

// ShowIndex writes the index help text to the screen
func ShowIndex() {
	fmt.Printf("%v\n", indexText)
}

// ShowSearch writes the search help text to the screen
func ShowSearch() {
	fmt.Printf("%v\n", searchText)
}
