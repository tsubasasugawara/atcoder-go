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

	ss := make([]string, 1000) 
	for i := 0; i < n; i++ {
		var s string 
		fmt.Fscan(in, &s)
		ss[i] = s
	}

	tt := make([]string, 1000)
	for i := 0; i < m; i++ {
		var t string 
		fmt.Fscan(in, &t)
		tt[i] = t
	}

	cnt := 0
	for i := 0; i < n; i++ {
		list := strings.Split(ss[i], "")
		suffix := list[3] + list[4] + list[5]
		match := false
		for j := 0; j < m; j++ {
			if suffix == tt[j] {
				match = true
			}
		}

		if match {
			cnt += 1
		}
	}

	fmt.Fprintln(out, cnt)
}
