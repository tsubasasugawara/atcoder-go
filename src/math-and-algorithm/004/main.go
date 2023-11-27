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

	var a1, a2, a3 int
	fmt.Fscan(in, &a1, &a2, &a3)

	fmt.Fprintln(out, a1 * a2 * a3)
}
