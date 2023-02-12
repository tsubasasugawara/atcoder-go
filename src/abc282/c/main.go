package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	list := strings.Split(s, "")

	q := 0
	for i, v := range list {
		if v == "\"" {
			q = 1 - q
			continue
		}
		if q == 0 && v == "," {
			list[i] = "."
		}
	}

	for _, v := range list {
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
