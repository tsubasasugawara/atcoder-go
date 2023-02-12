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

	var n, t int
	fmt.Fscan(in, &n, &t)

	list := make([]int, n+1)
	fmt.Fscan(in, &list[1])
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &list[i])
		list[i] = list [i] + list[i-1]
	}

	t = t % list[len(list) - 1]

	i := 1
	time := 0
	for ; i <= n; i++ {
		if list[i] > t {
			ilen := list[i] - list[i-1]
			time = ilen - (list[i] - t)
			break
		}
	}

	fmt.Fprintf(out, "%d %d\n", i, time)
}
