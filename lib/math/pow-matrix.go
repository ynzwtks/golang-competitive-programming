package main

// 正方行列の掛け算
func MulMatrix(m1, m2 [][]int, mod int) [][]int {
	size := len(m1)
	ret := make([][]int, size)
	for i := 0; i < size; i++ {
		ret[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				ret[i][k] += m1[i][j] * m2[j][k]
				ret[i][k] %= mod
			}
		}
	}
	return ret
}

// 正方行列の冪乗
func PowMatrix(m [][]int, n int, mod int) [][]int {
	size := len(m)
	tm := make([][]int, size)
	for i := 0; i < size; i++ {
		tm[i] = make([]int, size)
		tm[i][i] = 1
	}
	for n >= 2 {
		if n%2 == 0 {
			m = MulMatrix(m, m, mod)
			n = n / 2
		} else {
			tm = MulMatrix(tm, m, mod)
			m = MulMatrix(m, m, mod)
			n = (n - 1) / 2
		}
	}
	m = MulMatrix(tm, m, mod)
	return m
}
