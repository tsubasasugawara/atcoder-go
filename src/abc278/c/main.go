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

	var n, q int
	fmt.Fscan(in, &n, &q)

	following := make(map[string]bool)

	for i := 0; i < q; i++ {
		var t int
		var a, b string
		fmt.Fscan(in, &t, &a, &b)

		switch t {
		case 1:
			following[a + " " + b] = true
			break
		case 2:
			following[a+" "+b] = false
			break
		case 3:
			if following[a+" "+b] && following[b+" "+a] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
 }
