package main

import (
	"fmt"
	"os"
	"bufio"
)

func pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	
	list := []rune(s)
	var ans int = 0
	for i, _ := range list {
		ans += pow(26, len(list) - (i + 1)) * int(list[i] - 'A' + 1)
	}

	fmt.Fprintln(out, ans)
}
