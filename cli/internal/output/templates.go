package output

// text templates
const (
	search = `
{{print "Document: \033[1;33m"}}{{.DocumentID}}{{print "\033[0m"}}
{{print "Name:     \033[1;36m"}}{{.DocumentName}}{{print "\033[0m" }}

{{.Text}}

{{.DownloadURL}}

{{print "commands: \033[1;32m[n] next, [p] previous, [d] download file, [q] quit\033[0m"}}
`
)
