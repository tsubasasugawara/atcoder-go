package template

import ()

//  最大公約数
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a % b)
	} else {
		return a;
	}
}

// 複数の整数の最大公約数
func ngcd(a []int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = gcd(a[i], res)
	}
	return res;
}

// 最小公倍数
func lcm(a, b int) {
	return a / gcd(a, b) * b;
}

// 複数の整数の最小公倍数
func nlcm(a []int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = lcm(res, a[i])
	}
	return res
}

/* --------------------------------UnionFind-------------------------------- */
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
/* ------------------------------------------------------------------------- */

// 素数判定
func isPrime(n int) bool {
	if (n < 2) {
		return false
	}
	if (n == 2) {
		return true
	}
	if (n % 2 == 0) {
		return false
	}
	for i := 3; i * i < n; i += 2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// 桁和
func digsum(n int) int {
	res := 0
	for n > 0 {
		res += res % 10
		n /= 10
	}
	return res
}

// 約数全列挙
func enum_div(n int) []int {
	var ret []int
	for i := 1; i * i <= n; i++ {
		if n % i == 0 {
			ret = append(ret, i)
			if (i != 1 && i * i != n) {
				ret = append(ret, n / i)
			}
		}
	}
	return ret
}


/* --------------------------------スタック-------------------------------- */
type Stack[T any]  []T

func newStack[T any]() *Stack[T] {
	s := make(Stack[T], 0)
	return &s
}

func (s *Stack[T]) push(v T) {
	(*s) = append((*s), v)
}

func (s *Stack[T]) pop() T {
	n := len(*s) - 1
	res := (*s)[n]
	(*s) = (*s)[:n]
	return res
}

func (s *Stack[T]) isEmpty() bool {
	return len(*s) == 0
}
/* ------------------------------------------------------------------------ */

/* --------------------------------キュー-------------------------------- */
type Queue[T any] []T

func newQueue[T any]() *Queue[T] {
	q := make(Queue[T], 0)
	return &q
}

func (q *Queue[T]) push(v T) {
	(*q) = append((*q), v)
}

func (q *Queue[T]) pop() T {
	res := (*q)[0]
	(*q) = (*q)[1:]
	return res
}

func (q *Queue[T]) isEmpty() bool {
	return len(*q) == 0
}
/* ---------------------------------------------------------------------- */

/* --------------------------------幅優先探索-------------------------------- */
// 各座標の深さ情報
type Corr struct {
	x int
	y int
	depth int
}
func bfs(grid [][]int, q Queue[Corr]) int {
	isPassed := make([][]int, len(grid)) 
	for i := 0; i < len(isPassed); i++ {
		isPassed[i] = make([]int, len(grid[i]))
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	for !q.isEmpty() {
		now := q.push()
		for i := 0; i < 4; i++ {
			nx := now.x + dx[i]
			ny := now.y + dy[i]

			if nx >= len(grid[0]) || nx < 0 {
				continue
			}
			if ny >= len(grid) || ny < 0 {
				continue
			}
			if isPassed[ny][nx] {
				continue
			}

			isPassed[ny][nx] = true;
			next := new(Corr)
			next.x = nx
			next.y = ny
			next.depth = now.depth + 1
			q.push(next)
		}
	}
}
	}
/* -------------------------------------------------------------------------- */
