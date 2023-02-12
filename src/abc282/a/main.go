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

	var k int
	fmt.Fscan(in, &k)

	for i := 1; i <= k; i++ {
		fmt.Fprintf(out, "%c", 'A' - 1 + i)
	}
	fmt.Fprintln(out)
}
