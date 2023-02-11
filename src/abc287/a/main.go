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

	cnt := 0
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		if s == "Against" {
			cnt += 1
		} else {
			cnt +=-1 
		}
	}

	if cnt > 0 {
		fmt.Fprintln(out, "No")
	} else {
		fmt.Fprintln(out, "Yes")
	}
}
