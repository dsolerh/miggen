package scanutil

import (
	"io"
	"os"
)

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout
