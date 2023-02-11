package template

import ()

/*
 * 最大公約数
 */
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a % b)
	} else {
		return a;
	}
}

/*
 * 複数の整数の最大公約数
 */
func ngcd(a []int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = gcd(a[i], res)
	}
	return res;
}
