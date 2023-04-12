package waveletmatrix

//**********************************************************************
// Sample
//**********************************************************************
/*


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
	n    int      // 配列の長さ
	data []uint64 // bit情報
	sum  []uint32 // 0bitと1bitの個数の累積和
}

func NewBitVector(l []int, bit int) *BitVector {
	n := len(l)
	ret := &BitVector{}
	ret.n = n
	ret.data = make([]uint64, (n+63)/64) // ビットマップとしてのuint64配列
	ret.sum = make([]uint32, n)
	f := 1 << (bit - 1)
	s := uint32(0)
	for i := 0; i < n; i++ {
		if f&l[i] == f {
			ret.data[i/64] |= 1 << (i % 64) // 1を格納するビットをセット
			s++
		}
		ret.sum[i] = s
	}
	return ret
}

func (b *BitVector) OneNum() int {
	if len(b.data) == 0 {
		return 0
	}
	return int(b.sum[b.n-1])
}

func (b *BitVector) ZeroNum() int {
	if len(b.data) == 0 {
		return 0
	}
	return b.n - b.OneNum()
}

func (b *BitVector) Rank(pos, bit int) int {
	if bit == 0 {
		return pos + 1 - int(b.sum[pos])
	} else {
		return int(b.sum[pos])
	}
}
func (b *BitVector) GetBit(pos int) int {
	if (b.data[pos/64]>>(pos%64))&1 == 1 {
		return 1
	}
	return 0
}
func (b *BitVector) Select(rank, bit int) int {
	low, high, mid := 0, len(b.sum)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if b.Rank(mid, bit) >= rank {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	ret := low
	if bit == 0 {
		if 0 <= rank && rank <= b.ZeroNum() {
			return ret
		}
	} else {
		if 0 <= rank && rank <= b.OneNum() {
			return ret
		}
	}
	return -1
}

//**********************************************************************
// WaveletMatrix
//**********************************************************************
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
			if ret.wmd[idx].GetBit(j) == 0 {
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
		b := w.wmd[i].GetBit(idx) // 修正: data フィールドへの直接アクセスを GetBit に変更
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
	if len(w.ep) <= pos || w.ep[pos] != num {
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

		if w.wmd[i].GetBit(l) == 0 {
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
func (w *WaveletMatrix) Quantile2(l, r, n int) int {
	k := r - l - n + 2 //n番目に大きいをk番目に小さいに変換
	return w.Quantile(l, r, k)
}

//[l,r]の中でのxの出現回数を返す
func (w *WaveletMatrix) RangeFreq(l, r, x int) int {
	t1 := w.Rank(x, r)
	t2 := 0
	if l != 0 {
		t2 = w.Rank(x, l-1)
	}
	return t1 - t2
}
