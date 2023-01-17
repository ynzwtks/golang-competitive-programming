package waveletmatrix

//**********************************************************************
// Sample
//**********************************************************************
/*
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4, 8, 8, 4, 2, 1}
	wm := NewWaveletMatrix(a)
	fmt.Println(wm.Quantile2(0, 3, 2)) //4
	for i := 0; i < wm.bit; i++ {
		fmt.Println(wm.wmd[i])
	}
}
*/
//**********************************************************************
// BitVector
//**********************************************************************
type BitVector struct {
	n    int    //配列の長さ
	data []byte //bit情報
	sum0 []int  //0bitの個数の累積和
	sum1 []int  //1bitの個数の累積和
}

func NewBitVector(l []int, bit int) *BitVector {
	n := len(l)
	ret := &BitVector{}
	ret.n = n
	ret.data = make([]byte, n)
	ret.sum0 = make([]int, n)
	ret.sum1 = make([]int, n)
	f := 1 << (bit - 1)
	s0 := 0
	s1 := 0
	for i := 0; i < n; i++ {
		if f&l[i] == f {
			ret.data[i] = 1
			s1++
		} else {
			s0++
		}
		ret.sum0[i] = s0
		ret.sum1[i] = s1
	}
	return ret
}

// 1の個数の合計を返す
func (b *BitVector) OneNum() int {
	return b.sum1[b.n-1]
}

// 0の個数の合計を返す
func (b *BitVector) ZeroNum() int {
	return b.sum0[b.n-1]
}

// [0...pos]までのbit(0 or 1)の個数を返す
// posは0-indexed
func (b *BitVector) Rank(pos, bit int) int {
	if bit == 0 {
		return b.sum0[pos]
	} else {
		return b.sum1[pos]
	}
}

// bit(0 or 1)がrank回出現する位置を返す
func (b *BitVector) Select(rank, bit int) int {
	f := func(array []int, target int) int {
		low, high, mid := 0, len(array)-1, 0
		for low <= high {
			mid = (low + high) / 2
			if array[mid] >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return low
	}
	ret := -1
	if bit == 0 {
		if 0 <= rank && rank <= b.ZeroNum() {
			ret = f(b.sum0, rank)
		}
	} else {
		if 0 <= rank && rank <= b.OneNum() {
			ret = f(b.sum1, rank)
		}
	}
	return ret
}

//**********************************************************************
// WaveletMatrix
//**********************************************************************
//参考にしたもの：https://miti-7.hatenablog.com/entry/2018/04/28/152259

type WaveletMatrix struct {
	wmd []*BitVector
	sp  map[int]int // 入力したスライスの中に含まれる数値が左からみて初めて出現する位置(0-indexed)
	ep  []int       // Selectでの数値存在チェック用
	n   int         // 入力したスライスの長さ
	bit int         // 入力したスライスに含まれる最大値のbit長
}

func NewWaveletMatrix(a []int) *WaveletMatrix {
	ret := &WaveletMatrix{}
	f := func(x ...int) int {
		ret := 0
		for i := 0; i < len(x); i++ {
			if ret < x[i] {
				ret = x[i]
			}
		}
		return ret
	}
	ma := f(a...)
	t := 1
	ret.bit = 1
	for t < ma {
		t *= 2
		ret.bit++
	}
	ret.wmd = make([]*BitVector, ret.bit)
	ret.n = len(a)
	ret.sp = make(map[int]int)
	w1 := make([]int, ret.n)
	w2 := make([]int, ret.n)
	copy(w1, a)
	idx := 0
	p0, p1 := 0, 0
	for i := ret.bit; i > 0; i-- {
		ret.wmd[idx] = NewBitVector(w1, i)
		p0, p1 = 0, ret.wmd[idx].ZeroNum()

		for j := 0; j < ret.n; j++ {
			if ret.wmd[idx].data[j] == 0 {
				w2[p0] = w1[j]
				p0++
			} else {
				w2[p1] = w1[j]
				p1++
			}
		}
		w1, w2 = w2, w1
		idx++
	}
	for i := 0; i < ret.n; i++ {
		_, ok := ret.sp[w1[i]]
		if ok == false {
			ret.sp[w1[i]] = i
		}
	}
	ret.ep = w1
	return ret
}

// インデックス(0-indexed)の値を返す
func (w *WaveletMatrix) Access(idx int) int {
	ret := 0
	for i := 0; i < w.bit; i++ {
		b := w.wmd[i].data[idx]
		if b == 1 {
			ret += 1 << (w.bit - i - 1)
		}
		if b == 0 {
			idx = w.wmd[i].Rank(idx, 0) - 1
		} else {
			idx = w.wmd[i].Rank(idx, 1) + w.wmd[i].ZeroNum() - 1
		}
	}
	return ret
}

// (count)番目の数値(num)の位置を求める
// 数値が存在しない場合はfalseを返却
func (w *WaveletMatrix) Select(num, count int) (int, bool) {
	_, ok := w.sp[num]
	if ok == false {
		return 0, false
	}
	pos := w.sp[num] + count - 1
	if w.ep[pos] != num {
		return 0, false
	}
	for i := 0; i < w.bit; i++ {
		b := num >> i & 1
		if b == 0 {
			pos = w.wmd[w.bit-1-i].Select(pos+1, 0)
		} else {
			z := w.wmd[w.bit-1-i].ZeroNum()
			pos = w.wmd[w.bit-1-i].Select(pos-z+1, 1)
		}
	}
	return pos, true
}

// 数値nが[0,idx]までに出現する回数を返す
// 存在しない場合は0
func (w *WaveletMatrix) Rank(n, idx int) int {
	_, ok := w.sp[n]
	if ok == false {
		return 0
	}
	for i := 0; i < w.bit; i++ {
		b := (n >> (w.bit - i - 1)) & 1
		if b == 0 {
			idx = w.wmd[i].Rank(idx, 0) - 1
		} else {
			idx = w.wmd[i].Rank(idx, 1) + w.wmd[i].ZeroNum() - 1
		}
	}
	return idx - w.sp[n] + 1
}

//[l,r]の中でn番目に小さい値を返す
func (w *WaveletMatrix) Quantile(l, r, n int) int {
	if l < 0 || r >= w.n || l > r || r-l+1 < n {
		return -1
	}
	a0, a1, b0, b1 := 0, 0, 0, 0
	//a0:lより前の0の個数 a1:lより前の1の個数
	//b0:[l,r]内の0の個数 a1:[l,r]内の1の個数
	pos := n
	ret := 0
	bit := 0
	for i := 0; i < w.bit; i++ {
		a0 = w.wmd[i].Rank(l, 0)
		a1 = w.wmd[i].Rank(l, 1)
		b0 = w.wmd[i].Rank(r, 0) - a0
		b1 = w.wmd[i].Rank(r, 1) - a1

		if w.wmd[i].data[l] == 0 {
			a0--
			b0++
		} else {
			a1--
			b1++
		}
		if b0 < pos {
			ret += 1 << (w.bit - i - 1)
			pos -= b0
			bit = 1
		} else {
			bit = 0
		}
		if bit == 0 {
			l = a0
			r = a0 + b0 - 1
		} else {
			l = a1 + w.wmd[i].ZeroNum()
			r = a1 + b1 + w.wmd[i].ZeroNum() - 1
		}
	}
	return ret
}

//[l,r]の中でn番目に大きい値を返す
func (w *WaveletMatrix) Quantile2(l, r, n int) int {
	k := r - l - n + 2 //n番目に大きいをk番目に小さいに変換
	return w.Quantile(l, r, k)
}
