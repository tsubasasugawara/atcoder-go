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

	var s, t, x int
	fmt.Fscan(in, &s, &t, &x)

	if s > t {
		t += 24
		if x <= 12 {
			x += 24
		}
	}
	

	if s <= x && x < t {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
