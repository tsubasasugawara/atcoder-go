package main

import (
	"fmt"
	"os"
	"bufio"
)

func isMissJudge(h int, m int) bool {
	return (h % 10) >= 6 || ((m / 10) + (h / 10) * 10) >= 24
}

func next(h int, m int) (int, int) {
	if m == 59 {
		return (h + 1) % 24, 0
	} else {
		return h, m + 1 % 24
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, m int
	fmt.Fscan(in, &h, &m)

	for isMissJudge(h, m) {
		h, m = next(h, m)
	}

	fmt.Fprintf(out, "%d %d\n", h, m)
}
