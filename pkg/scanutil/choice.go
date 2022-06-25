package scanutil

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type ProcessChoice func(s string) (bool, error)

func ScannNChoice(question string, n int, fn ProcessChoice) bool {
	reader := bufio.NewReader(os.Stdin)
	result := false
	for i := 0; i < n; i++ {
		fmt.Print(question)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// remove the \n char
		result, err = fn(data[:len(data)-1])
		if err != nil {
			continue
		} else {
			break
		}
	}
	return result
}

func YesOrNo(data string) (bool, error) {
	if data == "" || data == "y" || data == "yes" {
		return true, nil
	}
	if data == "n" || data == "no" {
		return false, nil
	}
	return false, errors.New("invalid choice")
}
