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

	for i := 0; i < len(sl); i++ {
		if sl[i] != tl[i] {
			fmt.Fprintln(out, i + 1)
			return
		}
	}

	fmt.Fprintln(out, len(tl))
}
