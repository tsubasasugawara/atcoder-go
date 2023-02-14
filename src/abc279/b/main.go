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

	var s, t string
	fmt.Fscan(in, &s, &t)

	sl := strings.Split(s, "")
	tl := strings.Split(t, "")

	if len(tl) > len(sl) {
		fmt.Fprintln(out, "No")
		return
	}

	if len(tl) == len(sl) && s != t {
		fmt.Fprintln(out, "No")
		return
	}

	for i := 0; i < len(sl) - len(tl) + 1; {
		ok := true
		j := 0
		for ; j < len(tl); j++ {
			if tl[j] != sl[i + j] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Fprintln(out, "Yes")
			return
		}
		i += j + 1
	}

	fmt.Fprintln(out, "No")
}
