package log

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Error writes an error to stderr
func Error(s string) {
	red := color.New(color.FgRed).SprintFunc()

	fmt.Fprintf(os.Stderr, "%v %v\n", red("ERROR:"), s)
}
