package scanutil

import (
	"bufio"
	"fmt"
	"os"
)

func ScannNDate(question string, n int, def string) string {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Print(question)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// remove the \n char
		data = data[:len(data)-1]
		if data != "" {
			def = data
		}
		break
	}
	return def
}
