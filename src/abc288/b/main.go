package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

type Name struct {
	Name string
}

type Names []Name

func (n Names) Len() int {
	return len(n)
}

func (n Names) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Names) Less(i, j int) bool {
	return n[i].Name < n[j].Name
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	var names Names
	for i := 0; i < n; i++ {
		var name string
		fmt.Fscan(in, &name)
		if i < k {
			names = append(names, Name{name})
		}
	}

	sort.Sort(names)

	for i := 0; i < k; i++ {
		fmt.Fprintln(out, names[i].Name)
	}
}
