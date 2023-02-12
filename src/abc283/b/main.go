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

	var n int
	fmt.Fscan(in, &n)

	list := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		list[i] = a
	}

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var t, k int
		fmt.Fscan(in, &t, &k)

		if t == 1 {
			var x int
			fmt.Fscan(in, &x)
			list[k] = x
		} else {
			fmt.Fprintln(out, list[k])
		}
	}
}
