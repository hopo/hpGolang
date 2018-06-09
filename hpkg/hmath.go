// **********
// hpkg hmath
// **********

package hpkg

import (
// "sort"
)

// Factorial
func Fac(x int) int {
	if x == 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return x * Fac(x-1)
}

// Fibonacci
func Fibo(lim int) (flist []int) {
	if lim < 1 {
		return
	}
	a, b := 0, 1
	flist = []int{0}
	for i := 0; i < lim-1; i++ {
		flist = append(flist, b)
		a, b = b, a+b
	}
	return
}

// Denominator
func Denominator(n int) (rslt []int) {
	var is []int
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			is = append(is, i)
		}
	}
	rslt = is
	return
}

// Greater Common Denominator
func Gcd(a, b int) (rslt int) {
	sa := Denominator(a)
	sb := Denominator(b)
	for _, va := range sa {
		for _, vb := range sb {
			if va == vb {
				rslt = va
			}
		}
	}
	return
}

// Lower Common Multiple
func Lcm(a, b int) (rslt int) {
	return a * b / Gcd(a, b) // maybe need edit
}

// GCD & LCM
func Gcdlcm(a, b int) (rslt []int) {
	x := Gcd(a, b)
	y := Lcm(a, b)
	return []int{x, y}
}

// math power n**p
func Pow(n, p int) int {
	ret := 1
	for i := 0; i < p; i++ {
		ret *= n
	}
	return ret
}
