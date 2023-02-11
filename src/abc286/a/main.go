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
	var p, q, r, s int
	fmt.Fscan(in, &n, &p, &q, &r, &s)

	al := make([]int, 0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		al = append(al, a)
	}

	for i := 0; i <= q - p; i++ {
		al[(p + i) - 1], al[(r + i) - 1] = al[(r + i) - 1], al[(p + i) - 1]
	}

	fmt.Fprintf(out, "%d", al[0])
	for i := 1; i < len(al); i++ {
		fmt.Fprintf(out, " %d", al[i])
	}
	fmt.Fprintln(out)
}
