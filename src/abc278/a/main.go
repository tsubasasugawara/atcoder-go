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

	var n, k int
	fmt.Fscan(in, &n, &k)

	list := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		list[i] = a
	}

	if n <= k {
		list = make([]int, n)
	} else {
		list = append(list[k:], make([]int, k)...)
	}

	fmt.Fprint(out, list[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, " %d", list[i])
	}
	fmt.Fprintln(out)
}
