package output

// text templates
const (
	search = `
{{print "Name:     \033[1;36m"}}{{.DocumentName}}{{print "\033[0m" }}
{{print "Path:     \033[1;33m"}}{{.Path}}{{print "\033[0m" }}

{{.Text}}

{{print "commands: \033[1;32m[n] next, [p] previous, [d] download file, [q] quit\033[0m"}}
`
)
