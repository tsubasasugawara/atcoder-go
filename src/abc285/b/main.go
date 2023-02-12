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

	for i := 1; i <= n - 1; i++ {
		for l := n - i; l >= 0; l--  {
			if l == 0 {
				fmt.Fprintln(out, l)
				break
			}

			ok := true
			for k := 0; k < l; k++ {
				if list[k] == list[k+i] {
					ok = false
					break	
				}
			}
			if ok {
				fmt.Fprintln(out, l)
				break
			}
		}
	}
}
