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

	var a, b int
	fmt.Fscan(in, &a, &b)

	if int(b / 2) == a {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
