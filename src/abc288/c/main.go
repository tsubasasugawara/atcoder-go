package main

import (
	"fmt"
	"os"
	"bufio"
)

type UnionFind struct {
	par []int
}

func (u UnionFind) newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

func (u UnionFind) root(x int) int {
	if u.par[x] < 0 {
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

	u.par[yr] += u.par[xr]
	u.par[xr] = u.par[yr]
}

func (u UnionFind) same(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	}
	return false
}

func (u UnionFind) size(x int) int {
	return -u.par[u.root(x)]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	uni := newUnionFind(n)

	var a, b int
	cnt := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		if uni.root(a) == uni.root(b) {
			cnt += 1
		}
		uni.unite(a, b)
	}
	
	fmt.Fprintln(out, cnt)
}
