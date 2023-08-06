---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: lib/template/main.go
    title: lib/template/main.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/range_kth_smallest.test.go
    title: lib/waveletmatrix/range_kth_smallest.test.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix.go
    title: lib/waveletmatrix/waveletmatrix.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: lib/template/main.go
    title: lib/template/main.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix.go
    title: lib/waveletmatrix/waveletmatrix.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/range_kth_smallest.test.go
    title: lib/waveletmatrix/range_kth_smallest.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/opt/hostedtoolcache/Python/3.11.4/x64/lib/python3.11/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n          \
    \         ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n\
    \  File \"/opt/hostedtoolcache/Python/3.11.4/x64/lib/python3.11/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: lib/waveletmatrix/waveletmatrix_test.go\n"
  code: "package main\n\nimport (\n\t\"math/rand\"\n\t\"sort\"\n\t\"testing\"\n)\n\
    \nfunc TestNewWaveletMatrix(t *testing.T) {\n\ttestCases := [][]int{\n\t\t{},\n\
    \t\t{1, 1, 1, 1},\n\t\t{1, 2, 3, 4, 5},\n\t\t{5, 4, 3, 2, 1},\n\t}\n\n\tfor _,\
    \ testCase := range testCases {\n\t\t_ = NewWaveletMatrix(testCase)\n\t}\n}\n\n\
    func TestAccess(t *testing.T) {\n\tdata := []int{1, 2, 3, 4, 5}\n\twm := NewWaveletMatrix(data)\n\
    \n\tfor i, v := range data {\n\t\tif wm.Access(i) != v {\n\t\t\tt.Errorf(\"Access(%d)\
    \ returned %d, expected %d\", i, wm.Access(i), v)\n\t\t}\n\t}\n}\n\nfunc TestSelect(t\
    \ *testing.T) {\n\tdata := []int{1, 2, 2, 3, 3, 3}\n\twm := NewWaveletMatrix(data)\n\
    \n\tfor i := 1; i <= 3; i++ {\n\t\tfor j := 1; j <= i; j++ {\n\t\t\tpos, ok :=\
    \ wm.Select(i, j)\n\t\t\tif !ok || data[pos] != i {\n\t\t\t\tt.Errorf(\"Select(%d,\
    \ %d) returned %d, expected %d\", i, j, pos, (i-1)*i+j-1)\n\t\t\t}\n\t\t}\n\t\
    }\n\n\t_, ok := wm.Select(4, 1)\n\tif ok {\n\t\tt.Error(\"Select(4, 1) should\
    \ return false\")\n\t}\n}\n\nfunc TestRank(t *testing.T) {\n\tdata := []int{1,\
    \ 2, 1, 2, 1, 2}\n\twm := NewWaveletMatrix(data)\n\n\ttype query struct {\n\t\t\
    n, idx, want int\n\t}\n\n\tqueries := []query{\n\t\t{1, 0, 1},\n\t\t{1, 1, 1},\n\
    \t\t{1, 2, 2},\n\t\t{1, 3, 2},\n\t\t{2, 0, 0},\n\t\t{2, 1, 1},\n\t\t{2, 2, 1},\n\
    \t\t{2, 3, 2},\n\t\t{2, 4, 2},\n\t\t{2, 5, 3},\n\t}\n\n\tfor _, q := range queries\
    \ {\n\t\tgot := wm.Rank(q.n, q.idx)\n\t\tif got != q.want {\n\t\t\tt.Errorf(\"\
    Rank(%d, %d) returned %d, expected %d\", q.n, q.idx, got, q.want)\n\t\t}\n\t}\n\
    }\n\nfunc TestWaveletMatrix_KthSmall(t *testing.T) {\n\tdata := []int{1, 2, 2,\
    \ 3, 3, 3}\n\twm := NewWaveletMatrix(data)\n\n\tquantiles := []int{1, 2, 2, 3,\
    \ 3, 3}\n\tfor i, q := range quantiles {\n\t\tif wm.KthSmall(0, 5, i+1) != q {\n\
    \t\t\tt.Errorf(\"Quantile(0, 5, %d) returned %d, expected %d\", i+1, wm.KthSmall(0,\
    \ 5, i+1), q)\n\t\t}\n\t}\n}\nfunc TestWaveletMatrix_KthLarge(t *testing.T) {\n\
    \tdata := []int{1, 2, 2, 3, 3, 3}\n\twm := NewWaveletMatrix(data)\n\n\tquantiles2\
    \ := []int{3, 3, 3, 2, 2, 1}\n\tfor i, q := range quantiles2 {\n\t\tif wm.KthLarge(0,\
    \ 5, i+1) != q {\n\t\t\tt.Errorf(\"Quantile2(0, 5, %d) returned %d, expected %d\"\
    , i+1, wm.KthLarge(0, 5, i+1), q)\n\t\t}\n\t}\n}\n\nfunc TestWaveletMatrixRandomData(t\
    \ *testing.T) {\n\tdata := make([]int, 1000)\n\tfor i := range data {\n\t\tdata[i]\
    \ = rand.Intn(1000)\n\t}\n\n\twm := NewWaveletMatrix(data)\n\tsortedData := make([]int,\
    \ 1000)\n\tcopy(sortedData, data)\n\tsort.Ints(sortedData)\n\n\tfor i := 0; i\
    \ < 1000; i++ {\n\t\tq := wm.KthSmall(0, 999, i+1)\n\t\tif q != sortedData[i]\
    \ {\n\t\t\tt.Errorf(\"Quantile(0, 999, %d) returned %d, expected %d\", i+1, q,\
    \ sortedData[i])\n\t\t}\n\n\t\tq2 := wm.KthLarge(0, 999, i+1)\n\t\tif q2 != sortedData[999-i]\
    \ {\n\t\t\tt.Errorf(\"Quantile2(0, 999, %d) returned %d, expected %d\", i+1, q2,\
    \ sortedData[999-i])\n\t\t}\n\t}\n}\nfunc TestFreq(t *testing.T) {\n\tdata :=\
    \ []int{0, 1, 2, 3, 2, 2, 1}\n\twm := NewWaveletMatrix(data)\n\n\ttests := []struct\
    \ {\n\t\tl, r, x int\n\t\twant    int\n\t}{\n\t\t{l: 0, r: 6, x: 2, want: 3},\n\
    \t\t{l: 1, r: 6, x: 2, want: 3},\n\t\t{l: 2, r: 6, x: 2, want: 3},\n\t\t{l: 3,\
    \ r: 6, x: 2, want: 2},\n\t\t{l: 0, r: 6, x: 0, want: 1},\n\t\t{l: 0, r: 6, x:\
    \ 1, want: 2},\n\t\t{l: 0, r: 6, x: 3, want: 1},\n\t\t{l: 4, r: 4, x: 5, want:\
    \ 0},\n\t}\n\n\tfor _, test := range tests {\n\t\tresult := wm.Freq(test.l, test.r,\
    \ test.x)\n\t\tif result != test.want {\n\t\t\tt.Errorf(\"Freq(%d, %d, %d) = %d;\
    \ want %d\", test.l, test.r, test.x, result, test.want)\n\t\t}\n\t}\n}\nfunc TestRangeFreq(t\
    \ *testing.T) {\n\tdata := []int{0, 1, 2, 3, 2, 2, 1}\n\twm := NewWaveletMatrix(data)\n\
    \n\ttests := []struct {\n\t\tl, r, low, high int\n\t\twant            int\n\t\
    }{\n\t\t{l: 0, r: 6, low: 2, high: 2, want: 3},\n\t\t{l: 1, r: 6, low: 2, high:\
    \ 2, want: 3},\n\t\t{l: 2, r: 6, low: 2, high: 2, want: 3},\n\t\t{l: 3, r: 6,\
    \ low: 2, high: 2, want: 2},\n\t\t{l: 0, r: 6, low: 0, high: 0, want: 1},\n\t\t\
    {l: 0, r: 6, low: 1, high: 1, want: 2},\n\t\t{l: 0, r: 6, low: 3, high: 3, want:\
    \ 1},\n\t\t{l: 4, r: 4, low: 5, high: 5, want: 0},\n\t}\n\n\tfor _, test := range\
    \ tests {\n\t\tresult := wm.RangeFreq(test.l, test.r, test.low, test.high)\n\t\
    \tif result != test.want {\n\t\t\tt.Errorf(\"RangeFreq(%d, %d, %d, %d) = %d; want\
    \ %d\", test.l, test.r, test.low, test.high, result, test.want)\n\t\t}\n\t}\n\
    }\n"
  dependsOn:
  - lib/template/main.go
  - lib/waveletmatrix/waveletmatrix.go
  - lib/waveletmatrix/range_kth_smallest.test.go
  isVerificationFile: false
  path: lib/waveletmatrix/waveletmatrix_test.go
  requiredBy:
  - lib/template/main.go
  - lib/waveletmatrix/waveletmatrix.go
  timestamp: '2023-08-06 21:18:43+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - lib/waveletmatrix/range_kth_smallest.test.go
documentation_of: lib/waveletmatrix/waveletmatrix_test.go
layout: document
redirect_from:
- /library/lib/waveletmatrix/waveletmatrix_test.go
- /library/lib/waveletmatrix/waveletmatrix_test.go.html
title: lib/waveletmatrix/waveletmatrix_test.go
---
