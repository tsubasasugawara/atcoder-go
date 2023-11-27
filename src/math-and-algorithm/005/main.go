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

	sum := 0
	var a int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		sum += a % 100
		sum %= 100
	}

	fmt.Fprintln(out, sum)
}
