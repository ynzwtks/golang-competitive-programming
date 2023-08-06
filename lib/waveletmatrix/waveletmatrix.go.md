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
    path: lib/waveletmatrix/waveletmatrix_test.go
    title: lib/waveletmatrix/waveletmatrix_test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: lib/template/main.go
    title: lib/template/main.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix_test.go
    title: lib/waveletmatrix/waveletmatrix_test.go
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
    RuntimeError: bundler is not specified: lib/waveletmatrix/waveletmatrix.go\n"
  code: "package main\n\nimport (\n\t\"math\"\n\t\"sort\"\n)\n\ntype BitVector struct\
    \ {\n\tn    int      // \u914D\u5217\u306E\u9577\u3055\n\tdata []uint64 // bit\u60C5\
    \u5831\n\tsum  []uint32 // 0bit\u30681bit\u306E\u500B\u6570\u306E\u7D2F\u7A4D\u548C\
    \n}\n\nfunc NewBitVector(l []int, bit int) *BitVector {\n\tn := len(l)\n\tret\
    \ := &BitVector{}\n\tret.n = n\n\tret.data = make([]uint64, (n+63)/64) // \u30D3\
    \u30C3\u30C8\u30DE\u30C3\u30D7\u3068\u3057\u3066\u306Euint64\u914D\u5217\n\tret.sum\
    \ = make([]uint32, n)\n\tf := 1 << (bit - 1)\n\ts := uint32(0)\n\tfor i := 0;\
    \ i < n; i++ {\n\t\tif f&l[i] == f {\n\t\t\tret.data[i/64] |= 1 << (i % 64) //\
    \ 1\u3092\u683C\u7D0D\u3059\u308B\u30D3\u30C3\u30C8\u3092\u30BB\u30C3\u30C8\n\t\
    \t\ts++\n\t\t}\n\t\tret.sum[i] = s\n\t}\n\treturn ret\n}\n\nfunc (b *BitVector)\
    \ OneNum() int {\n\tif len(b.data) == 0 {\n\t\treturn 0\n\t}\n\treturn int(b.sum[b.n-1])\n\
    }\n\nfunc (b *BitVector) ZeroNum() int {\n\tif len(b.data) == 0 {\n\t\treturn\
    \ 0\n\t}\n\treturn b.n - b.OneNum()\n}\n\nfunc (b *BitVector) Rank(pos, bit int)\
    \ int {\n\tif pos == -1 {\n\t\treturn 0\n\t}\n\tif bit == 0 {\n\t\treturn pos\
    \ + 1 - int(b.sum[pos])\n\t} else {\n\t\treturn int(b.sum[pos])\n\t}\n}\nfunc\
    \ (b *BitVector) GetBit(pos int) int {\n\tif (b.data[pos/64]>>(pos%64))&1 == 1\
    \ {\n\t\treturn 1\n\t}\n\treturn 0\n}\nfunc (b *BitVector) Select(rank, bit int)\
    \ int {\n\tlow, high, mid := 0, len(b.sum)-1, 0\n\tfor low <= high {\n\t\tmid\
    \ = (low + high) / 2\n\t\tif b.Rank(mid, bit) >= rank {\n\t\t\thigh = mid - 1\n\
    \t\t} else {\n\t\t\tlow = mid + 1\n\t\t}\n\t}\n\tret := low\n\tif bit == 0 {\n\
    \t\tif 0 <= rank && rank <= b.ZeroNum() {\n\t\t\treturn ret\n\t\t}\n\t} else {\n\
    \t\tif 0 <= rank && rank <= b.OneNum() {\n\t\t\treturn ret\n\t\t}\n\t}\n\treturn\
    \ -1\n}\n\ntype WaveletMatrix struct {\n\twmd        []*BitVector\n\tsp      \
    \   map[int]int // \u5165\u529B\u3057\u305F\u30B9\u30E9\u30A4\u30B9\u306E\u4E2D\
    \u306B\u542B\u307E\u308C\u308B\u6570\u5024\u304C\u5DE6\u304B\u3089\u307F\u3066\
    \u521D\u3081\u3066\u51FA\u73FE\u3059\u308B\u4F4D\u7F6E(0-indexed)\n\tep      \
    \   []int       // Select\u3067\u306E\u6570\u5024\u5B58\u5728\u30C1\u30A7\u30C3\
    \u30AF\u7528\n\tn          int         // \u5165\u529B\u3057\u305F\u30B9\u30E9\
    \u30A4\u30B9\u306E\u9577\u3055\n\tbit        int         // \u5165\u529B\u3057\
    \u305F\u30B9\u30E9\u30A4\u30B9\u306B\u542B\u307E\u308C\u308B\u6700\u5927\u5024\
    \u306Ebit\u9577\n\tmaxElement int\n}\n\nfunc NewWaveletMatrix(a []int) *WaveletMatrix\
    \ {\n\tret := &WaveletMatrix{}\n\tsliceMax := func(x ...int) int {\n\t\tt := 0\n\
    \t\tfor i := 0; i < len(x); i++ {\n\t\t\tif t < x[i] {\n\t\t\t\tt = x[i]\n\t\t\
    \t}\n\t\t}\n\t\treturn t\n\t}\n\tret.maxElement = sliceMax(a...)\n\tt := 1\n\t\
    ret.bit = 1\n\tfor t < ret.maxElement {\n\t\tt *= 2\n\t\tret.bit++\n\t}\n\tret.wmd\
    \ = make([]*BitVector, ret.bit)\n\tret.n = len(a)\n\tret.sp = make(map[int]int)\n\
    \tw1 := make([]int, ret.n)\n\tw2 := make([]int, ret.n)\n\tcopy(w1, a)\n\tidx :=\
    \ 0\n\tp0, p1 := 0, 0\n\tfor i := ret.bit; i > 0; i-- {\n\t\tret.wmd[idx] = NewBitVector(w1,\
    \ i)\n\t\tp0, p1 = 0, ret.wmd[idx].ZeroNum()\n\t\tfor j := 0; j < ret.n; j++ {\n\
    \t\t\tif ret.wmd[idx].GetBit(j) == 0 {\n\t\t\t\tw2[p0] = w1[j]\n\t\t\t\tp0++\n\
    \t\t\t} else {\n\t\t\t\tw2[p1] = w1[j]\n\t\t\t\tp1++\n\t\t\t}\n\t\t}\n\t\tw1,\
    \ w2 = w2, w1\n\t\tidx++\n\t}\n\tfor i := 0; i < ret.n; i++ {\n\t\t_, ok := ret.sp[w1[i]]\n\
    \t\tif ok == false {\n\t\t\tret.sp[w1[i]] = i\n\t\t}\n\t}\n\tret.ep = w1\n\treturn\
    \ ret\n}\n\n// \u30A4\u30F3\u30C7\u30C3\u30AF\u30B9(0-indexed)\u306E\u5024\u3092\
    \u8FD4\u3059\nfunc (w *WaveletMatrix) Access(idx int) int {\n\tret := 0\n\tfor\
    \ i := 0; i < w.bit; i++ {\n\t\tb := w.wmd[i].GetBit(idx) // \u4FEE\u6B63: data\
    \ \u30D5\u30A3\u30FC\u30EB\u30C9\u3078\u306E\u76F4\u63A5\u30A2\u30AF\u30BB\u30B9\
    \u3092 GetBit \u306B\u5909\u66F4\n\t\tif b == 1 {\n\t\t\tret += 1 << (w.bit -\
    \ i - 1)\n\t\t}\n\t\tif b == 0 {\n\t\t\tidx = w.wmd[i].Rank(idx, 0) - 1\n\t\t\
    } else {\n\t\t\tidx = w.wmd[i].Rank(idx, 1) + w.wmd[i].ZeroNum() - 1\n\t\t}\n\t\
    }\n\treturn ret\n}\n\n// (count)\u756A\u76EE\u306E\u6570\u5024(num)\u306E\u4F4D\
    \u7F6E\u3092\u6C42\u3081\u308B\n// \u6570\u5024\u304C\u5B58\u5728\u3057\u306A\u3044\
    \u5834\u5408\u306Ffalse\u3092\u8FD4\u5374\nfunc (w *WaveletMatrix) Select(num,\
    \ count int) (int, bool) {\n\t_, ok := w.sp[num]\n\tif ok == false {\n\t\treturn\
    \ 0, false\n\t}\n\tpos := w.sp[num] + count - 1\n\tif len(w.ep) <= pos || w.ep[pos]\
    \ != num {\n\t\treturn 0, false\n\t}\n\tfor i := 0; i < w.bit; i++ {\n\t\tb :=\
    \ num >> i & 1\n\t\tif b == 0 {\n\t\t\tpos = w.wmd[w.bit-1-i].Select(pos+1, 0)\n\
    \t\t} else {\n\t\t\tz := w.wmd[w.bit-1-i].ZeroNum()\n\t\t\tpos = w.wmd[w.bit-1-i].Select(pos-z+1,\
    \ 1)\n\t\t}\n\t}\n\treturn pos, true\n}\n\n// \u6570\u5024n\u304C[0,idx]\u307E\
    \u3067\u306B\u51FA\u73FE\u3059\u308B\u56DE\u6570\u3092\u8FD4\u3059\n// \u5B58\u5728\
    \u3057\u306A\u3044\u5834\u5408\u306F0\nfunc (w *WaveletMatrix) Rank(n, idx int)\
    \ int {\n\t_, ok := w.sp[n]\n\tif ok == false {\n\t\treturn 0\n\t}\n\tfor i :=\
    \ 0; i < w.bit; i++ {\n\t\tb := (n >> (w.bit - i - 1)) & 1\n\t\tif b == 0 {\n\t\
    \t\tidx = w.wmd[i].Rank(idx, 0) - 1\n\t\t} else {\n\t\t\tidx = w.wmd[i].Rank(idx,\
    \ 1) + w.wmd[i].ZeroNum() - 1\n\t\t}\n\t}\n\treturn idx - w.sp[n] + 1\n}\n\n//\
    \ [l,r]\u306E\u4E2D\u3067n\u756A\u76EE\u306B\u5C0F\u3055\u3044\u5024\u3092\u8FD4\
    \u3059\nfunc (w *WaveletMatrix) KthSmall(l, r, n int) int {\n\tif l < 0 || r >=\
    \ w.n || l > r || r-l+1 < n {\n\t\treturn -1\n\t}\n\ta0, a1, b0, b1 := 0, 0, 0,\
    \ 0\n\t//a0:l\u3088\u308A\u524D\u306E0\u306E\u500B\u6570 a1:l\u3088\u308A\u524D\
    \u306E1\u306E\u500B\u6570\n\t//b0:[l,r]\u5185\u306E0\u306E\u500B\u6570 ab:[l,r]\u5185\
    \u306E1\u306E\u500B\u6570\n\tpos := n\n\tret := 0\n\tbit := 0\n\tfor i := 0; i\
    \ < w.bit; i++ {\n\t\ta0 = w.wmd[i].Rank(l, 0)\n\t\ta1 = w.wmd[i].Rank(l, 1)\n\
    \t\tb0 = w.wmd[i].Rank(r, 0) - a0\n\t\tb1 = w.wmd[i].Rank(r, 1) - a1\n\n\t\tif\
    \ w.wmd[i].GetBit(l) == 0 {\n\t\t\ta0--\n\t\t\tb0++\n\t\t} else {\n\t\t\ta1--\n\
    \t\t\tb1++\n\t\t}\n\t\tif b0 < pos {\n\t\t\tret += 1 << (w.bit - i - 1)\n\t\t\t\
    pos -= b0\n\t\t\tbit = 1\n\t\t} else {\n\t\t\tbit = 0\n\t\t}\n\t\tif bit == 0\
    \ {\n\t\t\tl = a0\n\t\t\tr = a0 + b0 - 1\n\t\t} else {\n\t\t\tl = a1 + w.wmd[i].ZeroNum()\n\
    \t\t\tr = a1 + b1 + w.wmd[i].ZeroNum() - 1\n\t\t}\n\t}\n\treturn ret\n}\nfunc\
    \ (w *WaveletMatrix) KthLarge(l, r, n int) int {\n\tk := r - l - n + 2 //n\u756A\
    \u76EE\u306B\u5927\u304D\u3044\u3092k\u756A\u76EE\u306B\u5C0F\u3055\u3044\u306B\
    \u5909\u63DB\n\treturn w.KthSmall(l, r, k)\n}\n\n// [l,r]\u306E\u4E2D\u3067\u306E\
    x\u306E\u51FA\u73FE\u56DE\u6570\u3092\u8FD4\u3059\nfunc (w *WaveletMatrix) Freq(l,\
    \ r, x int) int {\n\tt1 := w.Rank(x, r)\n\tt2 := 0\n\tif l != 0 {\n\t\tt2 = w.Rank(x,\
    \ l-1)\n\t}\n\treturn t1 - t2\n}\n\n// \u6307\u5B9A\u3057\u305Findex\u306E\u5024\
    \u304C[l,r]\u306E\u4E2D\u3067\u4F55\u756A\u76EE\u304B\u3092\u8FD4\u3059\nfunc\
    \ (w *WaveletMatrix) ConvertIndexToKth(l, r, index int, isPriorySmall bool) int\
    \ {\n\tif index < l || index > r {\n\t\treturn -math.MaxInt32\n\t}\n\tk := w.Access(index)\n\
    \tret := sort.Search(r-l+1, func(i int) bool {\n\t\tif isPriorySmall == true {\n\
    \t\t\tval := w.KthSmall(l, r, l+1)\n\t\t\treturn k <= val\n\t\t} else {\n\t\t\t\
    val := w.KthSmall(l, r, l+i)\n\t\t\treturn k >= val\n\t\t}\n\t})\n\treturn ret\
    \ + 1\n}\n\n// [l,r]\u306B\u3066low\u4EE5\u4E0A\u3001high\u4EE5\u4E0B\u306E\u8981\
    \u7D20\u6570\u3092\u30AB\u30A6\u30F3\u30C8\u3059\u308B\nfunc (w *WaveletMatrix)\
    \ RangeFreq(l, r, low, high int) int {\n\tif low == high {\n\t\treturn w.Freq(l,\
    \ r, low)\n\t}\n\treturn w.countLt(l, r, high+1) - w.countLt(l, r, low)\n}\n\n\
    // [l,r]\u3067v\u672A\u6E80\u306E\u8981\u7D20\u306E\u6570\u3092\u30AB\u30A6\u30F3\
    \u30C8\u3059\u308B\nfunc (w *WaveletMatrix) countLt(l, r, v int) int {\n\tif v\
    \ > w.maxElement {\n\t\treturn r - l\n\t}\n\tret := 0\n\tfor i := 0; i < w.bit;\
    \ i++ {\n\t\tb := (v >> (w.bit - i - 1)) & 1\n\t\tl0 := w.wmd[i].Rank(l, 0)\n\t\
    \tr0 := w.wmd[i].Rank(r, 0)\n\t\tif b == 1 {\n\t\t\tret += r0\n\t\t\tif l > 0\
    \ {\n\t\t\t\tret -= w.wmd[i].Rank(l-1, 0)\n\t\t\t}\n\t\t\tl += w.wmd[i].ZeroNum()\
    \ - l0\n\t\t\tr += w.wmd[i].ZeroNum() - r0\n\t\t} else {\n\t\t\tl = l0 - 1\n\t\
    \t\tr = r0 - 1\n\t\t}\n\t}\n\treturn ret\n}\n"
  dependsOn:
  - lib/template/main.go
  - lib/waveletmatrix/waveletmatrix_test.go
  - lib/waveletmatrix/range_kth_smallest.test.go
  isVerificationFile: false
  path: lib/waveletmatrix/waveletmatrix.go
  requiredBy:
  - lib/template/main.go
  - lib/waveletmatrix/waveletmatrix_test.go
  timestamp: '2023-08-06 21:18:43+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - lib/waveletmatrix/range_kth_smallest.test.go
documentation_of: lib/waveletmatrix/waveletmatrix.go
layout: document
redirect_from:
- /library/lib/waveletmatrix/waveletmatrix.go
- /library/lib/waveletmatrix/waveletmatrix.go.html
title: lib/waveletmatrix/waveletmatrix.go
---
