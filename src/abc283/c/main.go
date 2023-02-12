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

	var s string
	fmt.Fscan(in, &s)

	list := strings.Split(s, "")
	ans := 0

	for i := 0; i < len(list); i++ {
		if i < len(list) - 1 && list[i] == "0" && list[i+1] == "0" {
			i += 1
		}
		ans += 1
	}

	fmt.Fprintln(out, ans)
}
