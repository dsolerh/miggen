package scanutil

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type ProcessText func(string) (string, bool)

func ScannNText(question string, n int, def string, fn ProcessText) string {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Print(question)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// remove the \n char
		data, ok := fn(data[:len(data)-1])
		if !ok {
			continue
		} else {
			def = data
			break
		}
	}
	return def
}

func IsExt(s string) (string, bool) {
	if s != filepath.Ext(s) {
		return "", false
	}
	return s, true
}

func IsSep(s string) (string, bool) {
	if s != "|" && s != "-" && s != "_" && s != ":" {
		return "", false
	}
	return s, true
}
