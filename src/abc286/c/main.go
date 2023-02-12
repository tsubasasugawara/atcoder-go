package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a, b int
	var s string
	fmt.Fscan(in, &n, &a, &b, &s)

	ans := 1<<63 - 1
	list := strings.Split(s, "")
	for i := 0; i < n; i++ {
		tmp := a * i
		for j := 0; j < n/2; j++ {
			if list[j] != list[n - 1 - j] {
				tmp += b
			}
		}
		ans = int(math.Min(float64(ans), float64(tmp)))
		list = append(list[1:], list[0])
	}
	fmt.Fprintln(out, ans)
}
