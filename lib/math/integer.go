package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
}

//　因数分解
func factorization(n int) map[int]int {
	m := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n = n / i
		}
	}
	if n != 0 && n != 1 {
		m[n]++
	}
	return m
}

/// 約数のリスト
func DivisorList(n int) []int {
	var l []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			l = append(l, i)
			if i != n/i {
				l = append(l, n/i)
			}
		}
	}
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	return l
}

//素数リスト
func PrimeList(n int) []int {
	if n < 2 {
		return nil
	} else if n == 2 {
		return []int{2}
	}
	l := make([]int, n+1)
	primes := make([]int, 0, n)
	for i := 4; i <= n; i += 2 {
		l[i] = 1
	}
	for i := 3; i <= n; i += 2 {
		if l[i] == 1 {
			continue
		}
		for j := i * 2; j <= n; j += i {
			l[j] = 1
		}
	}
	primes = append(primes, 2)
	for i := 3; i <= n; i += 2 {
		if l[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}

//素数判定
func isPrime(n int) bool {

	c := int(math.Sqrt(float64(n)))
	for i := 2; i <= c; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 拡張ユークリッド互助法
func exgcd(a, b int) (int, int) {
	x1, y1, m1 := 1, 0, a
	x2, y2, m2 := 0, 1, b
	for b != 0 {
		t := a / b
		if m1 >= m2 {
			x1 -= (x2 * t)
			y1 -= (y2 * t)
			m1 -= (m2 * t)
		} else {
			x2 -= (x1 * t)
			y2 -= (y1 * t)
			m2 -= (m1 * t)
		}
		a, b = b, a%b
		if m1 == 1 {
			return x1, y1
		}
	}
	return x2, y2
}

// コンビネーションのリスト
func CombinationList(n, mod int) []int {
	ret := make([]int, n+1)
	ret[0], ret[n] = 1, 1
	if n == 1 {
		return ret
	}
	ret[1], ret[n-1] = n, n
	p := n
	for i := 2; i <= n-2; i++ {
		t := (p * (n - i + 1)) % mod
		p = (t * PowMod(i, mod-2, mod)) % mod
		ret[i] = p
	}
	return ret
}

// コンビネーション(オーバーフロー対策)
func Combination(n, r int) int {
	r = min(n-r, r)
	d := make([]int, r)

	for i := 0; i < r; i++ {
		d[i] = n - i
	}
	for i := 2; i <= r; i++ {
		ti := i
		for j := 0; j < r; j++ {
			g := gcd(d[j], ti)
			if g != 1 {
				ti /= g
				d[j] /= g
			}
			if ti == 1 {
				break
			}
		}
	}
	ret := 1
	for i := 0; i < r; i++ {
		ret *= d[i]
	}
	return ret
}

// コンビネーションの剰余
func CombinationMod(a, b, mod int) int {
	t1, t2 := 1, 1
	for i := 2; i <= b; i++ {
		t2 = (i * t2) % mod
	}
	gyaku := PowMod(t2, mod-2, mod)

	for i := a - b + 1; i <= a; i++ {
		t1 = (i * t1) % mod
	}
	return ((t1 * gyaku) % mod)
}

// べき乗の剰余
func PowMod(a, b, mod int) int {
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = (p * p) % mod
			q = q / 2
		} else {
			t = (t * p) % mod
			p = (p * p) % mod
			q = (q - 1) / 2
		}
	}
	return (t * p) % mod
}

// べき乗
func pow(a, b int) int {
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = (p * p)
			q = q / 2
		} else {
			t = (t * p)
			p = (p * p)
			q = (q - 1) / 2
		}
	}
	return (t * p)
}

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

// 複数のLCM
func LcmList(x []int) int {
	n := len(x)
	if n == 1 {
		return x[0]
	}
	ans := lcm(x[0], x[1])
	for i := 2; i < n; i++ {
		ans = lcm(ans, x[i])
	}
	return ans
}

// 最小公倍数
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

// 最大公約数
func gcd(a, b int) int {
	c := 1
	for c != 0 {
		c = a % b
		a, b = b, c
	}
	return a
}

// ２分探索
func binsearch(l []int, key int) int {
	left := 0
	right := len(l)
	for left < right {
		mid := (left + right) / 2
		if l[mid] == key {
			return mid
		} else if key > l[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

// リュカの定理(nCrのMod計算)
func Lucas(N, R, M int) int {
	// M=3のみに対応。ルーブの中のnCrがn=2の決め打ちになっている。
	// nCrのn<rの場合は0
	n := DecimalToBaseN(N, M)
	r := DecimalToBaseN(R, M)
	p := 1
	for j := 0; j < len(r); j++ {
		if r[j] > n[j] {
			p = 0
			break
		} else if n[j] == 2 && r[j] == 1 {
			p *= 2
		} else {
			p *= 1
		}
	}
	return p % M
}

// N進数変換
func DecimalToBaseN(decimal, base int) []int {
	ret := make([]int, 0)
	t := decimal
	if t == 0 {
		return []int{0}
	}
	for t != 0 {
		num := t % base
		t = t / base
		ret = append(ret, num)
	}
	return ret
}
