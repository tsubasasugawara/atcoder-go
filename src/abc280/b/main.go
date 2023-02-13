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

	list := make([]int, n + 1)
	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		list[i] = s
	}

	ans := make([]int, n)
	sum := make([]int, n)
	
	ans[0] = list[0]
	sum[0] = list[0]
	for i := 1; i < n; i++ {
		ans[i] = list[i] - sum[i-1]
		sum[i] = sum[i-1] + ans[i]
	}

	fmt.Fprintf(out, "%d", ans[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, " %d",  ans[i])
	}
	fmt.Fprintln(out)
}
