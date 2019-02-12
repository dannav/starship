## Using The Starship CLI

The Starship CLI requires an existing Starship instance to connect to. If you're interested in
setting up your own instance of Starship visit the section on [running Starship](/getting-started/running-starship).
Otherwise, when first attempting to `index` or `search`, the CLI will prompt you for the URL of
a Starship instance to connect to. This information will be stored in the Starship config located on your
machine at `~/.starship/starship.conf`.

### USAGE

	starship command [arguments]

### THE COMMANDS ARE

```
help      show the starship cli help text

index     store a document in the search engine

search    search all documents for relevant information
```

---

If you want to get more information on how to use a command you can type:

`starship help [command]`


For instance, to learn more about indexing a file from the command line:

`starship help index`

## Indexing Files

The Index command allows you to store a document in your teams search index for
retrieval using the search command.

### USAGE

`starship [-p path to store in search index (optional)] index [path to file]`

### EXAMPLE

`starship -p company index ./path/to/employee-handbook.pdf`

Think of the Starship search index as a logical directory of where files are stored.

The `-p` flag defines where you would like to store the document on Starship. All document
paths are unique and if the `-p` flag is not set the path will default to the root of
the search index.

If a file already exists at the path defined, Starship will prompt you if you want to overwrite
that file with the version you are going to index. If you choose not to overwrite the file the
program will exit and you'll have the opportunity to define a new path.

If you always want to overwrite files that exist when indexing and you do not want to receive a
prompt you can pass the `-y` flag to force a yes selection.

### EXAMPLE

`starship -y -p company index ./path/to/employee-handbook.pdf`

## Searching

### USAGE

`starship search [query]`

### EXAMPLE

`starship search paid time off`

The search command allows you to search all documents indexed with Starship for the given
text provided. Search does more than just look for matching keywords when returning results.
Starship uses machine learning to understand what you typed, and may return documents that are
similar in meaning.

For example using the search query in the example below (paid time off), Starship might return
a list of company holidays and hours found in documents indexed previously.

By default no search query will return empty results, unless no documents have been indexed.
Starship will try to return the best results it can find based on documents previously
indexed by you or your team.

When searching for documents you will be presented with commands to view the next result (n),
previous result (p), quit (q), or download the file to your current directory (d).