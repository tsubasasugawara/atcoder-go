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

	var s string
	fmt.Fscan(in, &s)

	ans := 0
	for _, v := range s {
		if v == 'v' {
			ans += 1
		} else {
			ans += 2
		}
	}
	fmt.Fprintln(out, ans)
}
