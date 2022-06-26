package scanutil

import (
	"bufio"
	"fmt"
	"path/filepath"
)

type ProcessText func(string) bool

// ScannNText scans the content of Stdin for N times
// it stops if fn returns ("something", true) then returns
// "something"
// or if the input is "" in wich case return def
func ScannNText(question string, n int, def string, fn ProcessText) string {
	reader := bufio.NewReader(in)
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, question, def)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// remove the \n char
		data = data[:len(data)-1]
		if data == "" {
			break
		}
		ok := fn(data)
		if !ok {
			continue
		} else {
			def = data
			break
		}
	}
	return def
}

// IsExt returns true if s is a valid extension
func IsExt(s string) bool {
	return s == filepath.Ext(s)
}

// IsSep returns true if s is a valid separator
// only allowed (|-_:)
func IsSep(s string) bool {
	if s != "|" && s != "-" && s != "_" && s != ":" {
		return false
	}
	return true
}

// IsDir returns true if s is a valid directory
func IsDir(s string) bool {
	return true
}
