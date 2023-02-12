package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var list []string
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		list = append(list, s)
	}

	for i := len(list) - 1; i >= 0; i-- {
		fmt.Fprintln(out, list[i])
	}
}
