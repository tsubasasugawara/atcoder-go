package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"reflect"
	"strings"
)

func trans(list []string) []string {
	xl := len(list[0])
	yl := len(list)

	l := make([][]string, yl)
	for i := 0; i < yl; i++ {
		l[i] = strings.Split(list[i], "")
	}

	t := make([][]string, xl)
	for i := 0; i < xl; i++ {
		t[i] = make([]string, yl)
	}

	for i := 0; i < yl; i++ {
		for j := 0; j < xl; j++ {
			t[j][i] = l[i][j]
		}
	}

	res := make([]string, xl)
	for i := 0; i < xl; i++ {
		res[i] = strings.Join(t[i], "")
	}

	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	sl := make([]string, h)
	tl := make([]string, h)

	for i := 0; i < h; i++ {
		fmt.Fscan(in, &sl[i])
	}

	for i := 0; i < h; i++ {
		fmt.Fscan(in, &tl[i])
	}

	sl = trans(sl)
	tl = trans(tl)

	sort.Strings(sl)
	sort.Strings(tl)

	if reflect.DeepEqual(sl, tl) {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
