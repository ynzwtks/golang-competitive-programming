package waveletmatrix

import (
	"math/rand"
	"sort"
	"testing"
)

func TestNewWaveletMatrix(t *testing.T) {
	testCases := [][]int{
		{},
		{1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for _, testCase := range testCases {
		_ = NewWaveletMatrix(testCase)
	}
}

func TestAccess(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	wm := NewWaveletMatrix(data)

	for i, v := range data {
		if wm.Access(i) != v {
			t.Errorf("Access(%d) returned %d, expected %d", i, wm.Access(i), v)
		}
	}
}

func TestSelect(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, 3}
	wm := NewWaveletMatrix(data)

	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			pos, ok := wm.Select(i, j)
			if !ok || data[pos] != i {
				t.Errorf("Select(%d, %d) returned %d, expected %d", i, j, pos, (i-1)*i+j-1)
			}
		}
	}

	_, ok := wm.Select(4, 1)
	if ok {
		t.Error("Select(4, 1) should return false")
	}
}

func TestRank(t *testing.T) {
	data := []int{1, 2, 1, 2, 1, 2}
	wm := NewWaveletMatrix(data)

	type query struct {
		n, idx, want int
	}

	queries := []query{
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 2},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 1},
		{2, 3, 2},
		{2, 4, 2},
		{2, 5, 3},
	}

	for _, q := range queries {
		got := wm.Rank(q.n, q.idx)
		if got != q.want {
			t.Errorf("Rank(%d, %d) returned %d, expected %d", q.n, q.idx, got, q.want)
		}
	}
}

func TestQuantile(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, 3}
	wm := NewWaveletMatrix(data)

	quantiles := []int{1, 2, 2, 3, 3, 3}
	for i, q := range quantiles {
		if wm.Quantile(0, 5, i+1) != q {
			t.Errorf("Quantile(0, 5, %d) returned %d, expected %d", i+1, wm.Quantile(0, 5, i+1), q)
		}
	}
}
func TestQuantile2(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, 3}
	wm := NewWaveletMatrix(data)

	quantiles2 := []int{3, 3, 3, 2, 2, 1}
	for i, q := range quantiles2 {
		if wm.Quantile2(0, 5, i+1) != q {
			t.Errorf("Quantile2(0, 5, %d) returned %d, expected %d", i+1, wm.Quantile2(0, 5, i+1), q)
		}
	}
}

func TestWaveletMatrixRandomData(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	wm := NewWaveletMatrix(data)
	sortedData := make([]int, 1000)
	copy(sortedData, data)
	sort.Ints(sortedData)

	for i := 0; i < 1000; i++ {
		q := wm.Quantile(0, 999, i+1)
		if q != sortedData[i] {
			t.Errorf("Quantile(0, 999, %d) returned %d, expected %d", i+1, q, sortedData[i])
		}

		q2 := wm.Quantile2(0, 999, i+1)
		if q2 != sortedData[999-i] {
			t.Errorf("Quantile2(0, 999, %d) returned %d, expected %d", i+1, q2, sortedData[999-i])
		}
	}
}
func TestRangeFreq(t *testing.T) {
	data := []int{0, 1, 2, 3, 2, 2, 1}
	wm := NewWaveletMatrix(data)

	tests := []struct {
		l, r, x int
		want    int
	}{
		{l: 0, r: 6, x: 2, want: 3},
		{l: 1, r: 6, x: 2, want: 3},
		{l: 2, r: 6, x: 2, want: 3},
		{l: 3, r: 6, x: 2, want: 2},
		{l: 0, r: 6, x: 0, want: 1},
		{l: 0, r: 6, x: 1, want: 2},
		{l: 0, r: 6, x: 3, want: 1},
		{l: 4, r: 4, x: 5, want: 0},
	}

	for _, test := range tests {
		result := wm.RangeFreq(test.l, test.r, test.x)
		if result != test.want {
			t.Errorf("RangeFreq(%d, %d, %d) = %d; want %d", test.l, test.r, test.x, result, test.want)
		}
	}
}
