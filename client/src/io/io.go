package io

import (
	"bufio"
	"fmt"
	"os"
)

func Prompt(str string) string {
	var input string

	fmt.Printf("%s", str)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	return input
}
