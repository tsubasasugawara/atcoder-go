package main

import (
	"fmt"
	"os"
	"bufio"
)

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	k := 1
	for n > 1 {
		if n%2 != 0 {
			k = k * x
			x = x * x
			n = (n - 1) / 2
		} else {
			x = x * x
			n = n / 2
		}
	}
	return k * x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int
	fmt.Fscan(in, &a, &b)

	fmt.Fprintln(out, pow(a, b))
}
