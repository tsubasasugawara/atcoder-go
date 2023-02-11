package main

import (
	"fmt"
	"os"
	"bufio"
)

type UnionFind struct {
	par []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = i
	}
	return u
}

func (u UnionFind) root(x int) int {
	if u.par[x] == x {
		return x
	}

	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}

	u.par[xr] = u.par[yr]
}

func (u UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var n, m int
	fmt.Fscan(in, &n, &m)

	if m != n - 1 {
		fmt.Fprintln(out, "No")
		return
	}

	uni := newUnionFind(n + 1)
	ok := true
	cnts := make([]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if uni.same(u, v) {
			ok = false
		}
		uni.unite(u, v)

		cnts[u] += 1
		cnts[v] += 1
		if cnts[u] > 2 || cnts [v] > 2 {
			fmt.Fprintln(out, "No")
			return
		}
	}

	if ok {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
