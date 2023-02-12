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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		ans := 0
		var a int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			if a % 2 != 0 {
				ans += 1
			}
		}
		fmt.Fprintln(out, ans)
	}
}
