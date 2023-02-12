package main

import (
	"fmt"
	"os"
	"bufio"
)

func isAlphabet(c int32) bool {
	return 'A' <= c && c <= 'Z'
}

func isNum(c int32) bool {
	return '0' <= c && c <= '9'
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	list := []rune(s)
	for i := 0; i < len(list); i++ {
		if !isAlphabet(list[i]) {
			fmt.Fprintln(out, "No")
			return
		}
		if len(list[i+1:]) < 6 || list[i+1] == '0' {
			fmt.Fprintln(out, "No")
			return
		}
		for j := 1; j <= 6; j++ {
			if !isNum(list[i + j]) {
				fmt.Fprintln(out, "No")
				return
			}
		}
		if len(list[i+7:]) == 0 ||  !isAlphabet(list[i + 7]) {
			fmt.Fprintln(out, "No")
			return
		}
		i += 7
	}
	fmt.Fprintln(out, "Yes")
}
