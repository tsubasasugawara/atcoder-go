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

	var n, m int
	fmt.Fscan(in, &n, &m)

	list := make([][]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		list[i] = strings.Split(s, "")
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ok := true
			for k := 0; k < m; k++ {
				if list[i][k] == "x" && list[j][k] == "x" {
					ok = false
				}
			}
			if ok {
				cnt += 1
			}
		}
	}

	fmt.Fprintln(out, cnt)
}
