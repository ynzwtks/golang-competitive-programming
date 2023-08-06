---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/range_kth_smallest.test.go
    title: lib/waveletmatrix/range_kth_smallest.test.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/static_range_frequency.test.go
    title: lib/waveletmatrix/static_range_frequency.test.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix.go
    title: lib/waveletmatrix/waveletmatrix.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix_test.go
    title: lib/waveletmatrix/waveletmatrix_test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix.go
    title: lib/waveletmatrix/waveletmatrix.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/waveletmatrix_test.go
    title: lib/waveletmatrix/waveletmatrix_test.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/range_kth_smallest.test.go
    title: lib/waveletmatrix/range_kth_smallest.test.go
  - icon: ':heavy_check_mark:'
    path: lib/waveletmatrix/static_range_frequency.test.go
    title: lib/waveletmatrix/static_range_frequency.test.go
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
    RuntimeError: bundler is not specified: lib/template/main.go\n"
  code: "package main\n\nimport (\n\t\"bufio\"\n\t\"container/list\"\n\t\"fmt\"\n\t\
    \"math\"\n\t\"math/big\"\n\t\"math/bits\"\n\t\"os\"\n\t\"reflect\"\n\t\"sort\"\
    \n\t\"strconv\"\n\t\"strings\"\n)\n\nconst (\n\tBUFSIZE                = 10000000\n\
    \tMOD2305843009213693951 = 2305843009213693951\n\tMOD1000000007          = 1000000007\n\
    \tMOD998244353           = 998244353\n\tMOD                    = MOD998244353\n\
    \tINF                    = 1 << 60\n)\n\nvar rdr *bufio.Scanner\nvar wtr *bufio.Writer\n\
    var p10 []int\n\nfunc solve() {\n\n}\nfunc flush() { wtr.Flush() }\nfunc main()\
    \ {\n\tdefer flush()\n\trdr = bufio.NewScanner(os.Stdin)\n\trdr.Buffer(make([]byte,\
    \ 4096), math.MaxInt64)\n\trdr.Split(bufio.ScanWords)\n\twtr = bufio.NewWriterSize(os.Stdout,\
    \ BUFSIZE)\n\tp10 = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000,\
    \ 100000000, 1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000,\
    \ 100000000000000, 1000000000000000, 10000000000000000, 100000000000000000, 1000000000000000000}\n\
    \tsolve()\n\n}\n\n// \u69CB\u9020\u4F53\u5B9A\u7FA9\ntype Data struct{ start,\
    \ end, val int }\ntype Pair struct{ a, b int }\ntype Item struct{ key, num int\
    \ }\ntype Point struct{ y, x float64 }\ntype Graph struct {\n\tn     int\n\tedges\
    \ [][]Edge\n}\ntype Edge struct {\n\tfrom, to, w int\n}\ntype FenwickTree struct\
    \ {\n\tdata []int\n\tlen  int\n}\ntype Bit2D struct {\n\tdata       [][]int\n\t\
    lenX, lenY int\n}\ntype UnionFind struct {\n\troot  []int\n\tsize  []int\n\tgroup\
    \ int\n}\ntype Deque struct {\n\tdeq *list.List\n\tcur *list.Element\n\tsum int\n\
    }\ntype DataDeque struct {\n\tdeq *list.List\n\tcur *list.Element\n}\ntype IntPQ\
    \ struct {\n\td                []int\n\tprioritySmallest bool\n}\ntype ItemPQ\
    \ struct {\n\td                []Item\n\tprioritySmallest bool\n}\ntype Frac struct\
    \ {\n\ttop    int\n\tbottom int\n}\ntype SegInt struct {\n\tdata []int\n\te  \
    \  int\n\top   func(int, int) int\n\tlen  int\n}\ntype BitSets []big.Int\n\n//\
    \ \u51FA\u529B\nfunc out(a ...interface{}) {\n\tstr := fmt.Sprintln(a...)\n\t\
    _, err := wtr.WriteString(str)\n\tif err != nil {\n\t\treturn\n\t}\n\treturn\n\
    }\nfunc outf(f float64) {\n\toutfmt(\"%0.9f\\n\", f)\n}\nfunc outn(x int) {\n\t\
    outfmt(\"%d \", x)\n}\nfunc outfn(f float64) {\n\toutfmt(\"%0.9f \", f)\n}\nfunc\
    \ outfmt(format string, a ...interface{}) (int, error) {\n\tstr := fmt.Sprintf(format,\
    \ a...)\n\treturn wtr.WriteString(str)\n}\nfunc outListX(l []int) {\n\ts := make([]string,\
    \ 0, len(l))\n\tfor i := 0; i < len(l); i++ {\n\t\ts = append(s, strconv.Itoa(l[i]))\n\
    \t}\n\twrite(strings.Join(s, \" \"))\n}\nfunc outListY(l []int) {\n\ts := make([]string,\
    \ 0, len(l))\n\tfor i := 0; i < len(l); i++ {\n\t\ts = append(s, strconv.Itoa(l[i]))\n\
    \t}\n\twrite(strings.Join(s, \"\\n\"))\n}\nfunc outYN(x bool) {\n\tif x == true\
    \ {\n\t\tout(\"Yes\")\n\t} else {\n\t\tout(\"No\")\n\t}\n}\nfunc write(s string)\
    \ {\n\t_, err := wtr.WriteString(s)\n\tif err != nil {\n\t\treturn\n\t}\n\t_,\
    \ err = wtr.WriteString(\"\\n\")\n\tif err != nil {\n\t\treturn\n\t}\n\treturn\n\
    }\nfunc outGrid(g [][]byte) {\n\tfor i := 0; i < len(g); i++ {\n\t\tout(string(g[i]))\n\
    \t}\n}\n\n// \u5165\u529B\nfunc rs() string { rdr.Scan(); return rdr.Text() }\n\
    func ri() int {\n\trdr.Scan()\n\ti, e := strconv.Atoi(rdr.Text())\n\tif e != nil\
    \ {\n\t\tpanic(e)\n\t}\n\treturn i\n}\nfunc ri2() (int, int) {\n\ta := ri()\n\t\
    b := ri()\n\treturn a, b\n}\nfunc ri3() (int, int, int) {\n\ta := ri()\n\tb :=\
    \ ri()\n\tc := ri()\n\treturn a, b, c\n}\nfunc ri4() (int, int, int, int) {\n\t\
    a := ri()\n\tb := ri()\n\tc := ri()\n\td := ri()\n\treturn a, b, c, d\n}\nfunc\
    \ ris(n int) []int {\n\tres := make([]int, n)\n\tfor i := 0; i < n; i++ {\n\t\t\
    res[i] = ri()\n\t}\n\treturn res\n}\nfunc riy(n int) []int {\n\tret := make([]int,\
    \ 0)\n\tfor i := 0; i < n; i++ {\n\t\tret = append(ret, ri())\n\t}\n\treturn ret\n\
    }\nfunc risapp(a []int, n int) []int {\n\tfor i := 0; i < n; i++ {\n\t\ta = append(a,\
    \ ri())\n\t}\n\treturn a\n}\nfunc ris2(n int) ([]int, []int) {\n\tres := make([]int,\
    \ n)\n\tres2 := make([]int, n)\n\tfor i := 0; i < n; i++ {\n\t\tres[i] = ri()\n\
    \t\tres2[i] = ri()\n\t}\n\treturn res, res2\n}\nfunc rf() float64 {\n\tf, e :=\
    \ strconv.ParseFloat(rs(), 64)\n\tif e != nil {\n\t\tpanic(e)\n\t}\n\treturn f\n\
    }\nfunc rfs(n int) []float64 {\n\tres := make([]float64, n)\n\tfor i := 0; i <\
    \ n; i++ {\n\t\tres[i] = rf()\n\t}\n\treturn res\n}\nfunc rss(n int) []string\
    \ {\n\tres := make([]string, n)\n\tfor i := 0; i < n; i++ {\n\t\tres[i] = rs()\n\
    \t}\n\treturn res\n}\nfunc readIntLines(h, w int) [][]int {\n\tl := make([][]int,\
    \ h)\n\tfor i := 0; i < h; i++ {\n\t\tl[i] = ris(w)\n\t}\n\treturn l\n}\nfunc\
    \ readGrid(h int) [][]byte {\n\tg := make([][]byte, h)\n\tfor i := 0; i < h; i++\
    \ {\n\t\tg[i] = []byte(rs())\n\t}\n\treturn g\n}\nfunc readEdges(m int, readWeight\
    \ bool) []Edge {\n\tret := make([]Edge, m)\n\tfor i := 0; i < m; i++ {\n\t\ta,\
    \ b := ri2()\n\t\tret[i].from = a\n\t\tret[i].to = b\n\t\tret[i].w = 1\n\t\tif\
    \ readWeight == true {\n\t\t\tret[i].w = ri()\n\t\t}\n\t}\n\treturn ret\n}\nfunc\
    \ readPairs(m int) []Pair {\n\tret := make([]Pair, m)\n\tfor i := 0; i < m; i++\
    \ {\n\t\ta, b := ri2()\n\t\tret[i] = Pair{a, b}\n\t}\n\treturn ret\n}\n\n// \u6C4E\
    \u7528\u95A2\u6570\nfunc abs(a int) int {\n\tif a < 0 {\n\t\treturn -a\n\t} else\
    \ {\n\t\treturn a\n\t}\n}\nfunc max(i int, j int) int {\n\tif i > j {\n\t\treturn\
    \ i\n\t}\n\treturn j\n}\nfunc maxf(i float64, j float64) float64 {\n\tif i > j\
    \ {\n\t\treturn i\n\t}\n\treturn j\n}\nfunc min(i int, j int) int {\n\tif i <\
    \ j {\n\t\treturn i\n\t}\n\treturn j\n}\nfunc minf(i float64, j float64) float64\
    \ {\n\tif i < j {\n\t\treturn i\n\t}\n\treturn j\n}\nfunc chmin(a *int, b int)\
    \ {\n\tif *a > b {\n\t\t*a = b\n\t}\n}\nfunc chmax(a *int, b int) {\n\tif *a <\
    \ b {\n\t\t*a = b\n\t}\n}\nfunc isInRange(x, low, high int) bool {\n\treturn low\
    \ <= x && x <= high\n}\nfunc isOutRange(x, low, high int) bool {\n\treturn !isInRange(x,\
    \ low, high)\n}\nfunc cs(a []int) []int {\n\tret := make([]int, len(a))\n\tcopy(ret,\
    \ a)\n\treturn ret\n}\nfunc maxSlice(l []int) int {\n\tret := math.MinInt64\n\t\
    for i := 0; i < len(l); i++ {\n\t\tif ret < l[i] {\n\t\t\tret = l[i]\n\t\t}\n\t\
    }\n\treturn ret\n}\nfunc minSlice(l []int) int {\n\tret := math.MaxInt64\n\tfor\
    \ i := 0; i < len(l); i++ {\n\t\tif ret > l[i] {\n\t\t\tret = l[i]\n\t\t}\n\t\
    }\n\treturn ret\n}\nfunc maxSliceIdx(l []int) int {\n\tcur := math.MinInt64\n\t\
    idx := 0\n\tfor i := 0; i < len(l); i++ {\n\t\tif cur < l[i] {\n\t\t\tcur = l[i]\n\
    \t\t\tidx = i\n\t\t}\n\t}\n\treturn idx\n}\nfunc minSliceIdx(l []int) int {\n\t\
    cur := math.MaxInt64\n\tidx := 0\n\tfor i := 0; i < len(l); i++ {\n\t\tif cur\
    \ > l[i] {\n\t\t\tcur = l[i]\n\t\t\tidx = i\n\t\t}\n\t}\n\treturn idx\n}\nfunc\
    \ maxMapKey(m map[int]int) int {\n\ttm, key := -INF, 0\n\tfor i, v := range m\
    \ {\n\t\tif v > tm {\n\t\t\ttm = v\n\t\t\tkey = i\n\t\t}\n\t}\n\treturn key\n\
    }\nfunc minMapKey(m map[int]int) int {\n\ttm, key := INF, 0\n\tfor i, v := range\
    \ m {\n\t\tif v < tm {\n\t\t\ttm = v\n\t\t\tkey = i\n\t\t}\n\t}\n\treturn key\n\
    }\nfunc maxMapVal(m map[int]int) int {\n\tt := -INF\n\tfor _, v := range m {\n\
    \t\tchmax(&t, v)\n\t}\n\treturn t\n}\nfunc minMapVal(m map[int]int) int {\n\t\
    t := INF\n\tfor _, v := range m {\n\t\tchmin(&t, v)\n\t}\n\treturn t\n}\nfunc\
    \ maxvar(x ...int) int { return maxSlice(x) }\nfunc minvar(x ...int) int { return\
    \ minSlice(x) }\nfunc condInt(a bool, b, c int) int {\n\tif a == true {\n\t\t\
    return b\n\t} else {\n\t\treturn c\n\t}\n}\nfunc condBool(a, b, c bool) bool {\n\
    \tif a == true {\n\t\treturn b\n\t} else {\n\t\treturn c\n\t}\n}\nfunc condStr(a\
    \ bool, b, c string) string {\n\tif a == true {\n\t\treturn b\n\t} else {\n\t\t\
    return c\n\t}\n}\nfunc divCeil(a, b int) int { return (a + b - 1) / b }\n\n//\
    \ Grid\u64CD\u4F5C\nfunc gridRangeChecker(aMin, aMax, bMin, bMax int) func(Pair)\
    \ bool {\n\tf := func(p Pair) bool {\n\t\treturn isInRange(p.a, aMin, aMax) &&\
    \ isInRange(p.b, bMin, bMax)\n\t}\n\treturn f\n}\nfunc createGrid(h, w int, ch\
    \ byte) [][]byte {\n\tret := make([][]byte, h)\n\tfor i := 0; i < h; i++ {\n\t\
    \tret[i] = make([]byte, w)\n\t\tif ch == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tfor j\
    \ := 0; j < w; j++ {\n\t\t\tret[i][j] = ch\n\t\t}\n\t}\n\treturn ret\n}\nfunc\
    \ rotateGrid(g [][]byte, rotation int, flipY, flipX bool) [][]byte {\n\th := len(g)\n\
    \tw := len(g[0])\n\tvar ret [][]byte\n\tif rotation == 1 || rotation == 3 {\n\t\
    \tret = createGrid(w, h, 0)\n\t} else {\n\t\tret = createGrid(h, w, 0)\n\t}\n\t\
    conv := gridPosConv(h, w, rotation, flipY, flipX)\n\tfor i := 0; i < h; i++ {\n\
    \t\tfor j := 0; j < w; j++ {\n\t\t\tch := g[i][j]\n\t\t\tni, nj := conv(i, j)\n\
    \t\t\tret[ni][nj] = ch\n\t\t}\n\t}\n\treturn ret\n}\nfunc gridPosConv(maxY, maxX,\
    \ rotation int, flipY, flipX bool) func(int, int) (int, int) {\n\treturn func(y,\
    \ x int) (int, int) {\n\t\tif flipY {\n\t\t\ty = maxY - 1 - y\n\t\t}\n\t\tif flipX\
    \ {\n\t\t\tx = maxX - 1 - x\n\t\t}\n\t\tswitch rotation % 4 {\n\t\tcase 1:\n\t\
    \t\ty, x = x, maxY-1-y\n\t\tcase 2:\n\t\t\ty, x = maxY-1-y, maxX-1-x\n\t\tcase\
    \ 3:\n\t\t\ty, x = maxX-1-x, y\n\t\t}\n\t\treturn y, x\n\t}\n}\nfunc clipGrid(grid\
    \ [][]byte, y, x, height, width int) [][]byte {\n\tif y < 0 || x < 0 || y+height\
    \ > len(grid) || x+width > len(grid[0]) {\n\t\tpanic(\"The provided range is out\
    \ of grid bounds.\")\n\t}\n\tclipped := make([][]byte, height)\n\tfor i := range\
    \ clipped {\n\t\tclipped[i] = make([]byte, width)\n\t}\n\tfor i := 0; i < height;\
    \ i++ {\n\t\tfor j := 0; j < width; j++ {\n\t\t\tclipped[i][j] = grid[y+i][x+j]\n\
    \t\t}\n\t}\n\treturn clipped\n}\nfunc pasteGrid(original, edited [][]byte, y,\
    \ x int) {\n\tif y < 0 || x < 0 || y+len(edited) > len(original) || x+len(edited[0])\
    \ > len(original[0]) {\n\t\tpanic(\"The provided range is out of original grid\
    \ bounds.\")\n\t}\n\tfor i := range edited {\n\t\tfor j := range edited[i] {\n\
    \t\t\toriginal[y+i][x+j] = edited[i][j]\n\t\t}\n\t}\n}\nfunc flipSubGrid(original\
    \ [][]byte, y, x, height, width int, flipY, flipX bool) {\n\textracted := clipGrid(original,\
    \ y, x, height, width)\n\tflipped := rotateGrid(extracted, 0, flipY, flipX)\n\t\
    pasteGrid(original, flipped, y, x)\n}\nfunc isSubGridMatch(grid [][]byte, y, x\
    \ int, subgrid [][]byte, eval func(byte, byte) bool) bool {\n\t//eval\u4F8B:\n\
    \t//func(a,b byte)bool{return  a == b || a == '?' || b == '?'}\n\tif y+len(subgrid)\
    \ > len(grid) || x+len(subgrid[0]) > len(grid[0]) {\n\t\treturn false\n\t}\n\t\
    for i := range subgrid {\n\t\tfor j := range subgrid[i] {\n\t\t\tif !eval(grid[y+i][x+j],\
    \ subgrid[i][j]) {\n\t\t\t\treturn false\n\t\t\t}\n\t\t}\n\t}\n\treturn true\n\
    }\nfunc strToGrid(s []string) [][]byte {\n\tret := make([][]byte, len(s))\n\t\
    for i := 0; i < len(s); i++ {\n\t\tret[i] = []byte(s[i])\n\t}\n\treturn ret\n\
    }\nfunc addGridBanpei(g [][]byte, c byte) [][]byte {\n\thead := make([]byte, len(g[0])+2)\n\
    \ttail := make([]byte, len(g[0])+2)\n\tfor i := 0; i < len(g[0])+2; i++ {\n\t\t\
    head[i] = c\n\t\ttail[i] = c\n\t}\n\tret := make([][]byte, len(g)+2)\n\tret[0]\
    \ = head\n\tret[len(g)+1] = tail\n\tfor i := 1; i <= len(g); i++ {\n\t\tt := g[i-1]\n\
    \t\tt = append([]byte{c}, t...)\n\t\tt = append(t, c)\n\t\tret[i] = t\n\t}\n\t\
    return ret\n}\n\n// Mod\nfunc modAdd(a, b int) int { return (a + b) % MOD }\n\
    func modSub(a, b int) int { return (((a - b) % MOD) + MOD) % MOD }\nfunc modMul(a,\
    \ b int) int { return (a * b) % MOD }\nfunc modInv(x int) int {\n\ta, _ := exgcd(x,\
    \ MOD)\n\treturn (a + MOD) % MOD\n}\nfunc sumMod(a []int) int {\n\tret := 0\n\t\
    for i := 0; i < len(a); i++ {\n\t\tret = modAdd(ret, a[i])\n\t}\n\treturn ret\n\
    }\n\n// 2\u5206\u63A2\u7D22\nfunc UpperBound(array []int, target int) int {\n\t\
    low, high, mid := 0, len(array)-1, 0\n\tfor low <= high {\n\t\tmid = (low + high)\
    \ / 2\n\t\tif array[mid] > target {\n\t\t\thigh = mid - 1\n\t\t} else {\n\t\t\t\
    low = mid + 1\n\t\t}\n\t}\n\treturn low\n}\nfunc LowerBound(array []int, target\
    \ int) int {\n\tlow, high, mid := 0, len(array)-1, 0\n\tfor low <= high {\n\t\t\
    mid = (low + high) / 2\n\t\tif array[mid] >= target {\n\t\t\thigh = mid - 1\n\t\
    \t} else {\n\t\t\tlow = mid + 1\n\t\t}\n\t}\n\treturn low\n}\nfunc rangeCount(a\
    \ []int, nmin, nmax int) int {\n\treturn UpperBound(a, nmax) - LowerBound(a, nmin)\n\
    }\nfunc searchNearestNum(a []int, x int) int {\n\tk := LowerBound(a, x)\n\tif\
    \ k == len(a) {\n\t\treturn a[k-1]\n\t}\n\tif k == 0 {\n\t\treturn a[0]\n\t}\n\
    \tif abs(a[k]-x) > abs(a[k-1]-x) {\n\t\treturn a[k-1]\n\t} else {\n\t\treturn\
    \ a[k]\n\t}\n}\nfunc countIntervalLen(a []int, x, lowerLimit, higherLimit int)\
    \ int {\n\tidx := UpperBound(a, x)\n\tlow := lowerLimit\n\thigh := higherLimit\n\
    \tif idx != 0 {\n\t\tlow = a[idx-1]\n\t}\n\tif idx != len(a) {\n\t\thigh = a[idx]\n\
    \t}\n\treturn high - low - 1\n} // \u6607\u9806\u30BD\u30FC\u30C8\u6E08\u307F\u306E\
    []int\u306B\u3064\u3044\u3066\u6307\u5B9A\u3057\u305F\u6570\u5024\u306E\u524D\u5F8C\
    \u8981\u7D20\u306E\u533A\u9593\u3092\u6C42\u3081\u308B\nfunc searchPairsRange(p\
    \ []Pair, l, r int) (int, int) {\n\tret := sort.Search(len(p), func(i int) bool\
    \ {\n\t\treturn l < p[i].a\n\t})\n\tif ret != 0 && l < p[ret-1].b {\n\t\tret--\n\
    \t}\n\tret2 := sort.Search(len(p), func(i int) bool {\n\t\treturn r < p[i].b\n\
    \t})\n\tif ret2 != len(p) && p[ret2].a > r {\n\t\tret2--\n\t}\n\tchmax(&ret, 0)\n\
    \tchmax(&ret2, 0)\n\tchmin(&ret, len(p)-1)\n\tchmin(&ret2, len(p)-1)\n\tchmin(&ret,\
    \ ret2)\n\treturn ret, ret2\n} // \u6607\u9806\u306E\u91CD\u8907\u306E\u306A\u3044\
    \u533A\u9593([]Pair)\u306B\u3064\u3044\u3066\u533A\u9593[l,r]\u304C\u542B\u307E\
    \u308C\u308Bindex\u306E\u7BC4\u56F2\u3092\u8FD4\u3059\n\n// \u30B9\u30E9\u30A4\
    \u30B9\u30FBmap\u64CD\u4F5C\u7CFB\nfunc sumIntSlice(l []int) int {\n\ts := 0\n\
    \tfor i := 0; i < len(l); i++ {\n\t\ts += l[i]\n\t}\n\treturn s\n}\nfunc slice2map(l\
    \ []int) map[int]int {\n\tm := make(map[int]int)\n\tfor i := 0; i < len(l); i++\
    \ {\n\t\tm[l[i]]++\n\t}\n\treturn m\n}\nfunc keys(m map[int]int) []int {\n\tret\
    \ := make([]int, 0, len(m))\n\tfor i := range m {\n\t\tret = append(ret, i)\n\t\
    }\n\treturn ret\n}\nfunc values(m map[int]int) []int {\n\tret := make([]int, 0,\
    \ len(m))\n\tfor _, v := range m {\n\t\tret = append(ret, v)\n\t}\n\treturn ret\n\
    }\nfunc reverseSort(a []int) []int {\n\tsort.Slice(a, func(i, j int) bool { return\
    \ a[i] > a[j] })\n\treturn a\n}\nfunc reverseList(x []int) []int {\n\tret := make([]int,\
    \ 0, len(x))\n\tfor i := len(x) - 1; i >= 0; i-- {\n\t\tret = append(ret, x[i])\n\
    \t}\n\treturn ret\n}\nfunc reverseBytes(x []byte) []byte {\n\tret := make([]byte,\
    \ 0, len(x))\n\tfor i := len(x) - 1; i >= 0; i-- {\n\t\tret = append(ret, x[i])\n\
    \t}\n\treturn ret\n}\nfunc haveSameKeys(map1, map2 interface{}) bool {\n\tval1,\
    \ val2 := reflect.ValueOf(map1), reflect.ValueOf(map2)\n\tif val1.Kind() != reflect.Map\
    \ || val2.Kind() != reflect.Map {\n\t\tpanic(\"Both parameters must be maps\"\
    )\n\t}\n\tif val1.Len() != val2.Len() {\n\t\treturn false\n\t}\n\tfor _, key :=\
    \ range val1.MapKeys() {\n\t\tif val2.MapIndex(key).Kind() == reflect.Invalid\
    \ {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\nfunc intSliceCnt(a []int,\
    \ maxRange int) []int {\n\tret := intSlice(maxRange, 0)\n\tfor i := 0; i < len(a);\
    \ i++ {\n\t\tret[a[i]]++\n\t}\n\treturn ret\n}\nfunc defaultIntMap(d int) (map[int]int,\
    \ func(int) int) {\n\tdefaultVal := d\n\tm := make(map[int]int)\n\taccessor :=\
    \ func(x int) int {\n\t\tval, ok := m[x]\n\t\tif ok == true {\n\t\t\treturn val\n\
    \t\t} else {\n\t\t\treturn defaultVal\n\t\t}\n\t}\n\treturn m, accessor\n}\nfunc\
    \ transpose(matrix [][]int) [][]int {\n\tif len(matrix) == 0 || len(matrix[0])\
    \ == 0 {\n\t\treturn matrix\n\t}\n\trows := len(matrix)\n\tcols := len(matrix[0])\n\
    \ttransposeMatrix := make([][]int, cols)\n\tfor i := 0; i < cols; i++ {\n\t\t\
    transposeMatrix[i] = make([]int, rows)\n\t\tfor j := 0; j < rows; j++ {\n\t\t\t\
    transposeMatrix[i][j] = matrix[j][i]\n\t\t}\n\t}\n\treturn transposeMatrix\n}\n\
    func removeElements(s []int, start, end int) []int {\n\tif start < 0 || start\
    \ >= len(s) || end < 0 || end >= len(s) || start > end {\n\t\treturn nil\n\t}\n\
    \treturn append(s[:start], s[end+1:]...)\n} // \u30B9\u30E9\u30A4\u30B9\u306E\
    [begin,end]\u306E\u533A\u9593\u3092\u524A\u9664\u3059\u308B\nfunc insertSliceAt(index\
    \ int, a []int, b []int) []int {\n\tif index < 0 || index > len(a) {\n\t\tout(\"\
    Error: index out of range\")\n\t\treturn nil\n\t}\n\ta = append(a, make([]int,\
    \ len(b))...)\n\tcopy(a[index+len(b):], a[index:])\n\tcopy(a[index:], b)\n\treturn\
    \ a\n} // \u30B9\u30E9\u30A4\u30B9a\u306Eindex\u306E\u4F4D\u7F6E\u306Bb\u3092\u633F\
    \u5165\u3059\u308B\nfunc intSlice(a, value int) []int {\n\tret := make([]int,\
    \ a)\n\tif value == 0 {\n\t\treturn ret\n\t}\n\tfor i := 0; i < a; i++ {\n\t\t\
    ret[i] = value\n\t}\n\treturn ret\n}\nfunc intSlice2(a, b, value int) [][]int\
    \ {\n\tret := make([][]int, a)\n\tfor i := 0; i < a; i++ {\n\t\tret[i] = make([]int,\
    \ b)\n\t\tif value == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tfor j := 0; j < b; j++ {\n\
    \t\t\tret[i][j] = value\n\t\t}\n\t}\n\treturn ret\n}\nfunc intSlice3(a, b, c,\
    \ value int) [][][]int {\n\tret := make([][][]int, a)\n\tfor i := 0; i < a; i++\
    \ {\n\t\tret[i] = make([][]int, b)\n\t\tfor j := 0; j < b; j++ {\n\t\t\tret[i][j]\
    \ = make([]int, c)\n\t\t\tif value == 0 {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tfor\
    \ k := 0; k < c; k++ {\n\t\t\t\tret[i][j][k] = value\n\t\t\t}\n\t\t}\n\t}\n\t\
    return ret\n}\nfunc floatSlice(a int, value float64) []float64 {\n\tret := make([]float64,\
    \ a)\n\tif value == 0 {\n\t\treturn ret\n\t}\n\tfor i := 0; i < a; i++ {\n\t\t\
    ret[i] = value\n\t}\n\treturn ret\n}\nfunc floatSlice2(a, b int, value float64)\
    \ [][]float64 {\n\tret := make([][]float64, a)\n\tfor i := 0; i < a; i++ {\n\t\
    \tret[i] = make([]float64, b)\n\t\tif value == 0 {\n\t\t\tcontinue\n\t\t}\n\t\t\
    for j := 0; j < b; j++ {\n\t\t\tret[i][j] = value\n\t\t}\n\t}\n\treturn ret\n\
    }\nfunc boolSlice(a int) []bool {\n\tret := make([]bool, a)\n\treturn ret\n}\n\
    func boolSlice2(a, b int) [][]bool {\n\tret := make([][]bool, a)\n\tfor i := 0;\
    \ i < a; i++ {\n\t\tret[i] = make([]bool, b)\n\t}\n\treturn ret\n}\nfunc rangeSlice(begin,\
    \ end int) []int {\n\tret := make([]int, 0, end-begin+1)\n\tfor i := begin; i\
    \ <= end; i++ {\n\t\tret = append(ret, i)\n\t}\n\treturn ret\n}\nfunc rotate90(a\
    \ [][]int) [][]int {\n\tw, h := len(a), len(a[0])\n\tres := make([][]int, h)\n\
    \tfor i := 0; i < h; i++ {\n\t\tres[i] = make([]int, w)\n\t\tfor j := 0; j < w;\
    \ j++ {\n\t\t\tres[i][j] = a[j][h-i-1]\n\t\t}\n\t}\n\treturn res\n}\nfunc intsToBytes(a\
    \ []int, add int) []byte {\n\tret := make([]byte, len(a))\n\tfor i := 0; i < len(a);\
    \ i++ {\n\t\tret[i] = byte(a[i] + add)\n\t}\n\treturn ret\n}\nfunc bytesToInts(a\
    \ []byte, add int) []int {\n\tret := make([]int, len(a))\n\tfor i := 0; i < len(a);\
    \ i++ {\n\t\tret[i] = int(a[i]) + add\n\t}\n\treturn ret\n}\nfunc stringIndexer()\
    \ (func(string) int, func(int) string) {\n\tcur := 0\n\tm := make(map[string]int)\n\
    \tr := make(map[int]string)\n\tf := func(s string) int {\n\t\tv := m[s]\n\t\t\
    if v == 0 {\n\t\t\tcur++\n\t\t\tm[s] = cur\n\t\t\tr[cur] = s\n\t\t\treturn cur\n\
    \t\t}\n\t\treturn v\n\t}\n\tf2 := func(x int) string {\n\t\treturn r[x]\n\t}\n\
    \treturn f, f2\n} // \u6587\u5B57\u5217\u3092index\u306B\u5909\u63DB\nfunc compressCoordinate(arr\
    \ []int) ([]int, map[int]int, map[int]int) {\n\tarrCopy := append([]int(nil),\
    \ arr...)\n\tsort.Ints(arrCopy)\n\tm := make(map[int]int)\n\trev := make(map[int]int)\n\
    \trank := 0\n\tfor _, a := range arrCopy {\n\t\tif _, ok := m[a]; !ok {\n\t\t\t\
    m[a] = rank\n\t\t\trank++\n\t\t}\n\t}\n\tresult := make([]int, len(arr))\n\tfor\
    \ i, a := range arr {\n\t\tresult[i] = m[a]\n\t\trev[i] = a\n\t}\n\treturn result,\
    \ m, rev\n} // \u5EA7\u6A19\u5727\u7E2E\n\n// Pair\u64CD\u4F5C\nfunc dedupPair(p\
    \ []Pair, KeyIsA bool, f func(int, int) int) []Pair {\n\tt := make(map[int]int)\n\
    \tret := make([]Pair, 0)\n\tif f == nil {\n\t\tfor k := range pairs2map(p) {\n\
    \t\t\tret = append(ret, k)\n\t\t}\n\t\treturn ret\n\t}\n\tfor i := 0; i < len(p);\
    \ i++ {\n\t\tif KeyIsA == true {\n\t\t\tval, ok := t[p[i].a]\n\t\t\tif ok == true\
    \ {\n\t\t\t\tt[p[i].a] = f(val, p[i].b)\n\t\t\t} else {\n\t\t\t\tt[p[i].a] = p[i].b\n\
    \t\t\t}\n\t\t} else {\n\t\t\tval, ok := t[p[i].b]\n\t\t\tif ok == true {\n\t\t\
    \t\tt[p[i].b] = f(val, p[i].a)\n\t\t\t} else {\n\t\t\t\tt[p[i].b] = p[i].a\n\t\
    \t\t}\n\t\t}\n\t}\n\tfor k, v := range t {\n\t\tif KeyIsA == true {\n\t\t\tret\
    \ = append(ret, Pair{k, v})\n\t\t} else {\n\t\t\tret = append(ret, Pair{v, k})\n\
    \t\t}\n\t}\n\treturn ret\n}\nfunc sortPairs(p []Pair, sortByA, prioritySmallestKey1,\
    \ prioritySmallestKey2 bool) {\n\tf := []func(int, int) bool{func(a, b int) bool\
    \ { return a < b }, func(a, b int) bool { return a > b }}\n\tf1 := condInt(prioritySmallestKey1,\
    \ 0, 1)\n\tf2 := condInt(prioritySmallestKey2, 0, 1)\n\tif sortByA == true {\n\
    \t\tsort.Slice(p, func(i, j int) bool { return condBool(p[i].a == p[j].a, f[f2](p[i].b,\
    \ p[j].b), f[f1](p[i].a, p[j].a)) })\n\t} else {\n\t\tsort.Slice(p, func(i, j\
    \ int) bool { return condBool(p[i].b == p[j].b, f[f2](p[i].a, p[j].a), f[f1](p[i].b,\
    \ p[j].b)) })\n\t}\n}\nfunc pairs2map(p []Pair) map[Pair]int {\n\tret := make(map[Pair]int)\n\
    \tfor i := 0; i < len(p); i++ {\n\t\tret[p[i]]++\n\t}\n\treturn ret\n}\nfunc map2pair(m\
    \ map[int]int) []Pair {\n\tt := make([]Pair, 0, len(m))\n\tfor i, v := range m\
    \ {\n\t\tt = append(t, Pair{i, v})\n\t}\n\treturn t\n}\nfunc splitPair(p []Pair)\
    \ ([]int, []int) {\n\tres := make([]int, len(p))\n\tres2 := make([]int, len(p))\n\
    \tfor i, v := range p {\n\t\tres[i] = v.a\n\t\tres2[i] = v.b\n\t}\n\treturn res,\
    \ res2\n}\nfunc slice2Pairs(a, b []int) []Pair {\n\tif len(a) != len(b) {\n\t\t\
    return nil\n\t}\n\tret := make([]Pair, 0, len(a))\n\tfor i := 0; i < len(a); i++\
    \ {\n\t\tret = append(ret, Pair{a[i], b[i]})\n\t}\n\treturn ret\n}\n\n// \u6587\
    \u5B57\u5217\nfunc findChar(s string, p byte) ([]int, []int) {\n\tret := make([]int,\
    \ 0)\n\tret2 := make([]int, 0)\n\tfor i := 0; i < len(s); i++ {\n\t\tif s[i] ==\
    \ p {\n\t\t\tret = append(ret, i)\n\t\t\tret2 = append(ret2, len(s)-i-1)\n\t\t\
    }\n\t}\n\treturn ret, ret2\n} // \u5DE6\u3068\u53F3\u304B\u3089\u6587\u5B57\u304C\
    \u51FA\u73FE\u3059\u308B\u4F4D\u7F6E\u3092\u8FD4\u3059\nfunc repeat(s string,\
    \ n int) string {\n\treturn strings.Repeat(s, n)\n}\n\n// \u30B0\u30E9\u30D5\n\
    func NewGraph(n int) *Graph {\n\tg := &Graph{}\n\tg.n = n\n\tg.edges = make([][]Edge,\
    \ g.n)\n\treturn g\n}\nfunc (g *Graph) AddEdge(a, b, c int) {\n\tvar t Edge\n\t\
    t.from, t.to, t.w = a, b, c\n\tg.edges[a] = append(g.edges[a], t)\n}\nfunc (g\
    \ *Graph) ReadSimpleUndirectedGraph(m int) {\n\tfor i := 0; i < m; i++ {\n\t\t\
    a, b := ri2()\n\t\tg.AddEdge(a, b, 1)\n\t\tg.AddEdge(b, a, 1)\n\t}\n}\nfunc (g\
    \ *Graph) ReadWightedUndirectedGraph(m int) {\n\tfor i := 0; i < m; i++ {\n\t\t\
    a, b, c := ri3()\n\t\tg.AddEdge(a, b, c)\n\t\tg.AddEdge(b, a, c)\n\t}\n}\nfunc\
    \ (g *Graph) ReadSimpleDirectedGraph(m int) {\n\tfor i := 0; i < m; i++ {\n\t\t\
    a, b := ri2()\n\t\tg.AddEdge(a, b, 1)\n\t}\n}\nfunc (g *Graph) ReadWightedDirectedGraph(m\
    \ int) {\n\tfor i := 0; i < m; i++ {\n\t\ta, b, c := ri3()\n\t\tg.AddEdge(a, b,\
    \ c)\n\t}\n}\nfunc (g *Graph) PairToEdge(p []Pair, undirected bool) {\n\tfor i\
    \ := 0; i < len(p); i++ {\n\t\tg.AddEdge(p[i].a, p[i].b, 1)\n\t\tif undirected\
    \ == true {\n\t\t\tg.AddEdge(p[i].b, p[i].a, 1)\n\t\t}\n\t}\n}\nfunc (g *Graph)\
    \ SortEdgeByNode(prioritySmallest bool) {\n\tfor v := 0; v < len(g.edges); v++\
    \ {\n\t\tif prioritySmallest == true {\n\t\t\tsort.Slice(g.edges[v], func(i, j\
    \ int) bool {\n\t\t\t\treturn g.edges[v][i].to < g.edges[v][j].to\n\t\t\t})\n\t\
    \t} else {\n\t\t\tsort.Slice(g.edges[v], func(i, j int) bool {\n\t\t\t\treturn\
    \ g.edges[v][i].to > g.edges[v][j].to\n\t\t\t})\n\t\t}\n\t}\n}\nfunc readMatrixEdges(n,\
    \ m int, readCost, directed bool) [][]int {\n\tret := intSlice2(n, n, 0)\n\ta,\
    \ b, c := 0, 0, 1\n\tfor i := 0; i < m; i++ {\n\t\tif readCost == true {\n\t\t\
    \ta, b, c = ri3()\n\t\t} else {\n\t\t\ta, b = ri2()\n\t\t}\n\t\tret[a][b] = c\n\
    \t\tif directed == false {\n\t\t\tret[b][a] = c\n\t\t}\n\t}\n\treturn ret\n} //\
    \ \u96A3\u63A5\u884C\u5217\u306E\u8AAD\u307F\u8FBC\u307F\nfunc restorePath(p []int,\
    \ begin, end int) []int {\n\tret := make([]int, 0)\n\tcur := end\n\tret = append(ret,\
    \ cur)\n\tfor cur != -1 && cur != begin {\n\t\tret = append(ret, p[cur])\n\t\t\
    cur = p[cur]\n\t}\n\tret = reverseList(ret)\n\treturn ret\n} // \u6700\u77ED\u30D1\
    \u30B9\u5FA9\u5143(BFS,\u30C0\u30A4\u30AF\u30B9\u30C8\u30E9,\u30D9\u30EB\u30DE\
    \u30F3\u30D5\u30A9\u30FC\u30C9\u5171\u901A)\nfunc searchGrid(g [][]byte, s string)\
    \ [][]Pair {\n\tret := make([][]Pair, len(s))\n\tfor i := 0; i < len(g); i++ {\n\
    \t\tfor j := 0; j < len(g[0]); j++ {\n\t\t\tfor k := 0; k < len(s); k++ {\n\t\t\
    \t\tif g[i][j] == s[k] {\n\t\t\t\t\tret[k] = append(ret[k], Pair{i, j})\n\t\t\t\
    \t}\n\t\t\t}\n\t\t}\n\t}\n\treturn ret\n} //Grid\u4E0A\u306E\u5BFE\u8C61\u306E\
    \u5EA7\u6A19\u3092\u63A2\u7D22\n\n// BIT\nfunc NewFenwickTree(n int) *FenwickTree\
    \ {\n\tret := &FenwickTree{}\n\tret.data = make([]int, n)\n\tret.len = n\n\treturn\
    \ ret\n}\nfunc (f *FenwickTree) Add(p, x int) {\n\tif p < 0 || p >= f.len {\n\t\
    \tpanic(\"worong range\")\n\t\treturn\n\t}\n\tp += 1\n\tfor p <= f.len {\n\t\t\
    f.data[p-1] += x\n\t\tp += p & -p\n\t}\n}\nfunc (f *FenwickTree) Replace(p, x\
    \ int) {\n\tif p < 0 || p >= f.len {\n\t\tpanic(\"wrong range\")\n\t\treturn\n\
    \t}\n\toldValue := f.Sum(p, p+1)\n\tdiff := x - oldValue\n\tf.Add(p, diff)\n}\n\
    func (f *FenwickTree) Sum(l, r int) int {\n\tif l < 0 || l >= f.len || r < 0 ||\
    \ r > f.len || l > r {\n\t\treturn 0\n\t}\n\tsum := func(r int) int {\n\t\ts :=\
    \ 0\n\t\tfor r > 0 {\n\t\t\ts += f.data[r-1]\n\t\t\tr -= r & -r\n\t\t}\n\t\treturn\
    \ s\n\t}\n\treturn sum(r) - sum(l)\n} // \u533A\u9593[l,r)\u306E\u5408\u8A08\n\
    func (f *FenwickTree) ModSum(l, r, mod int) int {\n\tif l < 0 || l >= f.len ||\
    \ r < 0 || r > f.len || l > r {\n\t\treturn 0\n\t}\n\tsum := func(r int) int {\n\
    \t\ts := 0\n\t\tfor r > 0 {\n\t\t\ts += f.data[r-1]\n\t\t\ts %= mod\n\t\t\tr -=\
    \ r & -r\n\t\t}\n\t\treturn s % mod\n\t}\n\treturn modSub(sum(r), sum(l))\n} //\
    \ \u533A\u9593[l,r)\u306E\u5408\u8A08(MOD)\nfunc (f *FenwickTree) Set(l []int)\
    \ {\n\tn := len(l)\n\tif n != f.len {\n\t\tpanic(\"worong length\")\n\t\treturn\n\
    \t}\n\tfor i, v := range l {\n\t\tf.Add(i, v)\n\t}\n}\nfunc (f *FenwickTree) LowerBound(l,\
    \ k int) (int, int, bool) {\n\tif f.Sum(l, f.len) < k {\n\t\treturn f.len, 0,\
    \ false\n\t}\n\tb := f.Sum(0, l)\n\tidx := sort.Search(f.len, func(i int) bool\
    \ {\n\t\tt := f.Sum(0, i)\n\t\treturn t-b >= k\n\t})\n\treturn idx, f.Sum(l, idx),\
    \ true\n}\nfunc (f *FenwickTree) RangeCount(idx, nmax int) int {\n\tret1, _, ok\
    \ := f.LowerBound(idx, nmax+1)\n\tret2, _, _ := f.LowerBound(idx, 0)\n\tif ok\
    \ == true {\n\t\treturn ret1 - ret2 - 1\n\t} else {\n\t\treturn ret1 - ret2\n\t\
    }\n}\nfunc countNumberOfFall(a []int) int {\n\tret := 0\n\tt, _, _ := compressCoordinate(a)\n\
    \tf := NewFenwickTree(len(t))\n\tfor i := len(t) - 1; i >= 0; i-- {\n\t\tf.Add(t[i],\
    \ 1)\n\t\tret += f.Sum(0, t[i])\n\t}\n\treturn ret\n} // \u5EA7\u5727\u3057\u3066\
    \u8EE2\u5012\u6570\u3092\u6C42\u3081\u308B\nfunc rangeBIT(n int) (func(int, int,\
    \ int), func(int, int) int) {\n\tf0 := NewFenwickTree(n + 1)\n\tf1 := NewFenwickTree(n\
    \ + 1)\n\trangeAddFn := func(l, r, val int) {\n\t\tf0.Add(l, -val*l)\n\t\tf0.Add(r,\
    \ val*r)\n\t\tf1.Add(l, val)\n\t\tf1.Add(r, -val)\n\t}\n\tsumFn := func(l, r int)\
    \ int {\n\t\tret := f0.Sum(0, r) + f1.Sum(0, r)*r - f0.Sum(0, l) + f1.Sum(0, l)*l\n\
    \t\treturn ret\n\t}\n\treturn rangeAddFn, sumFn\n} // [l,r)\u306B\u3064\u3044\u3066\
    val\u3092\u533A\u9593\u52A0\u7B97\u3001[l,r)\u306E\u533A\u9593\u548C\u3092\u8FD4\
    \u3059\u30AF\u30ED\u30FC\u30B8\u30E3\u30FC\u3092\u8FD4\u3059\n\n// 2\u6B21\u5143\
    BIT\nfunc NewBit2D(x, y int) *Bit2D {\n\tret := &Bit2D{}\n\tret.data = make([][]int,\
    \ x+1)\n\tfor i := range ret.data {\n\t\tret.data[i] = make([]int, y+1)\n\t}\n\
    \tret.lenX = x + 1\n\tret.lenY = y + 1\n\treturn ret\n}\nfunc (f *Bit2D) Add(x,\
    \ y, v int) {\n\tx += 1\n\ty += 1\n\tif x <= 0 || x >= f.lenX || y <= 0 || y >=\
    \ f.lenY {\n\t\tpanic(\"wrong range\")\n\t\treturn\n\t}\n\tfor i := x; i < f.lenX;\
    \ i += i & -i {\n\t\tfor j := y; j < f.lenY; j += j & -j {\n\t\t\tf.data[i][j]\
    \ += v\n\t\t}\n\t}\n}\nfunc (f *Bit2D) Replace(x, y, v int) {\n\toldValue := f.Sum(x,\
    \ y, x+1, y+1)\n\tdiff := v - oldValue\n\tf.Add(x, y, diff)\n}\nfunc (f *Bit2D)\
    \ Sum(x1, y1, x2, y2 int) int {\n\tsum := func(x, y int) int {\n\t\ts := 0\n\t\
    \tfor i := x; i > 0; i -= i & -i {\n\t\t\tfor j := y; j > 0; j -= j & -j {\n\t\
    \t\t\ts += f.data[i][j]\n\t\t\t}\n\t\t}\n\t\treturn s\n\t}\n\treturn sum(x2, y2)\
    \ - sum(x1, y2) - sum(x2, y1) + sum(x1, y1)\n}\n\n// 2\u6B21\u5143\u7D2F\u7A4D\
    \u548C\nfunc imos2D(h, w int) (func(int, int, int), func(int, int, int, int) int)\
    \ {\n\ttable := intSlice2(h, w, 0)\n\tdoneCompute := false\n\tadd := func(y, x,\
    \ val int) {\n\t\tif x < 0 || y < 0 || x >= w || y >= h {\n\t\t\tpanic(\"Out of\
    \ range\")\n\t\t}\n\t\ttable[y][x] += val\n\t}\n\tcompute := func() {\n\t\tfor\
    \ i := 0; i < h; i++ {\n\t\t\tfor j := 0; j < w; j++ {\n\t\t\t\tif i > 0 {\n\t\
    \t\t\t\ttable[i][j] += table[i-1][j]\n\t\t\t\t}\n\t\t\t\tif j > 0 {\n\t\t\t\t\t\
    table[i][j] += table[i][j-1]\n\t\t\t\t}\n\t\t\t\tif i > 0 && j > 0 {\n\t\t\t\t\
    \ttable[i][j] -= table[i-1][j-1]\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\trangeSum :=\
    \ func(y1, x1, y2, x2 int) int {\n\t\tif doneCompute == false {\n\t\t\tcompute()\n\
    \t\t\tdoneCompute = true\n\t\t}\n\t\tif y1 >= y2 || x1 >= x2 {\n\t\t\treturn 0\n\
    \t\t}\n\t\td1, d2, d3 := 0, 0, 0\n\t\tif y1 != 0 && x1 != 0 {\n\t\t\td1 = table[y1-1][x1-1]\n\
    \t\t}\n\t\tif y1 != 0 {\n\t\t\td2 = table[y1-1][x2-1]\n\t\t}\n\t\tif x1 != 0 {\n\
    \t\t\td3 = table[y2-1][x1-1]\n\t\t}\n\t\treturn table[y2-1][x2-1] + d1 - d2 -\
    \ d3\n\t}\n\treturn add, rangeSum\n} // 2\u6B21\u5143\u7D2F\u7A4D\u548C\u306E\
    add(y,x)\u3068rangeSum(y,x\u306E\u9589\u958B\u533A\u9593)\u306E\u30AF\u30ED\u30FC\
    \u30B8\u30E3\u30FC\u3092\u8FD4\u3059\n\n// \u30BB\u30B0\u30E1\u30F3\u30C8\u6728\
    (int)\nfunc NewSegInt(a []int, e int, op func(int, int) int) *SegInt {\n\tret\
    \ := &SegInt{}\n\tret.e = e\n\tret.op = op\n\tret.len = 1\n\tfor ret.len < len(a)\
    \ {\n\t\tret.len *= 2\n\t}\n\tret.data = make([]int, ret.len*2-1)\n\n\tfor i :=\
    \ 0; i < len(a); i++ {\n\t\tret.data[i] = a[i]\n\t}\n\tfor i := len(a); i < ret.len;\
    \ i++ {\n\t\tret.data[i] = ret.e\n\t}\n\tfor i := 0; i < ret.len*2-2; i += 2 {\n\
    \t\tret.data[i/2+ret.len] = op(ret.data[i], ret.data[i+1])\n\t}\n\treturn ret\n\
    }\nfunc (seg *SegInt) Update(idx int, val int) {\n\tcur := idx\n\tseg.data[cur]\
    \ = val\n\tfor cur < seg.len*2-2 {\n\t\tif cur%2 == 1 {\n\t\t\tseg.data[cur/2+seg.len]\
    \ = seg.op(seg.data[cur-1], seg.data[cur])\n\t\t} else {\n\t\t\tseg.data[cur/2+seg.len]\
    \ = seg.op(seg.data[cur], seg.data[cur+1])\n\t\t}\n\t\tcur = cur/2 + seg.len\n\
    \t}\n}\nfunc (seg *SegInt) Query(l, r int) int {\n\tret := seg.e\n\tfor l < r\
    \ {\n\t\tif l%2 == 1 {\n\t\t\tret = seg.op(ret, seg.data[l])\n\t\t\tl++\n\t\t\
    }\n\t\tif r%2 == 1 {\n\t\t\tret = seg.op(ret, seg.data[r-1])\n\t\t\tr--\n\t\t\
    }\n\t\tif l == r {\n\t\t\treturn ret\n\t\t}\n\t\tl = l/2 + seg.len\n\t\tr = r/2\
    \ + seg.len\n\t}\n\tret = seg.op(ret, seg.data[l])\n\treturn ret\n} // \u534A\u958B\
    \u533A\u9593[l,r)\u306E\u30AF\u30A8\u30EA\nfunc sliceOrderChecker(a []int, asc\
    \ bool) (func(int, int), func(int, int) bool) {\n\tdat := make([]int, len(a))\n\
    \tcopy(dat, a)\n\tvar fn func(int, int) int\n\tif asc == true {\n\t\tfn = func(u,\
    \ v int) int { return condInt(u <= v, 1, 0) }\n\t} else {\n\t\tfn = func(u, v\
    \ int) int { return condInt(u >= v, 1, 0) }\n\t}\n\td := make([]int, len(a))\n\
    \td[len(a)-1] = 1\n\tfor i := 0; i < len(a)-1; i++ {\n\t\td[i] = fn(a[i], a[i+1])\n\
    \t}\n\tseg := NewSegInt(d, 1, func(u, v int) int { return condInt(u == 1 && v\
    \ == 1, 1, 0) })\n\tupdate := func(idx, val int) {\n\t\tdat[idx] = val\n\t\tif\
    \ len(dat) == 1 {\n\t\t\treturn\n\t\t}\n\t\tif idx != len(dat)-1 {\n\t\t\tseg.Update(idx,\
    \ fn(dat[idx], dat[idx+1]))\n\t\t}\n\t\tif idx != 0 {\n\t\t\tseg.Update(idx-1,\
    \ fn(dat[idx-1], dat[idx]))\n\t\t}\n\t}\n\tquery := func(l, r int) bool {\n\t\t\
    if r-l == 0 || seg.Query(l, r-1) == 1 {\n\t\t\treturn true\n\t\t}\n\t\treturn\
    \ false\n\t}\n\treturn update, query\n} // \u30B9\u30E9\u30A4\u30B9\u306E[l,r)\u306B\
    \u3064\u3044\u3066\u6607\u9806(or \u964D\u9806)\u306E\u5834\u5408\u306Btrue\u3068\
    \u306A\u308B\u30AF\u30ED\u30FC\u30B8\u30E3\u30FC\u3092\u8FD4\u3059\nfunc charSet(a\
    \ []int) (func(int, int), func(int, int) ([]int, int, int, []int)) {\n\tconst\
    \ maxCharTypes = 26\n\tdat := make([]int, len(a))\n\tcopy(dat, a)\n\tcount :=\
    \ intSlice(maxCharTypes, 0)\n\tseg := make([]*SegInt, maxCharTypes)\n\tfor i :=\
    \ 0; i < maxCharTypes; i++ {\n\t\tseg[i] = NewSegInt(intSlice(len(a), 0), 0, func(a,\
    \ b int) int { return a + b })\n\t}\n\tfor i := 0; i < len(a); i++ {\n\t\tidx\
    \ := a[i]\n\t\tseg[idx].Update(i, 1)\n\t\tcount[idx]++\n\t}\n\tupdate := func(idx,\
    \ val int) {\n\t\told := dat[idx]\n\t\tnext := val\n\t\tdat[idx] = next\n\t\t\
    seg[old].Update(idx, 0)\n\t\tseg[next].Update(idx, 1)\n\t\tcount[old]--\n\t\t\
    count[next]++\n\t}\n\tquery := func(l, r int) ([]int, int, int, []int) {\n\t\t\
    ret := intSlice(maxCharTypes, 0)\n\t\tminIdx, maxIdx := -1, -1\n\t\tfor i := 0;\
    \ i < maxCharTypes; i++ {\n\t\t\tt := seg[i].Query(l, r)\n\t\t\tif minIdx == -1\
    \ && t != 0 {\n\t\t\t\tminIdx = i\n\t\t\t}\n\t\t\tif t != 0 {\n\t\t\t\tmaxIdx\
    \ = i\n\t\t\t}\n\t\t\tret[i] += t\n\t\t}\n\t\treturn ret, minIdx, maxIdx, count\n\
    \t}\n\treturn update, query\n} // \u533A\u9593[l,r)\u306B\u542B\u307E\u308C\u308B\
    \u6587\u5B57\u30AB\u30A6\u30F3\u30C8\u3001\u8F9E\u66F8\u9806\u306E\u6700\u5C0F\
    \u6587\u5B57\u3001\u6700\u5927\u6587\u5B57\u3001\u5168\u4F53\u306E\u6587\u5B57\
    \u30AB\u30A6\u30F3\u30C8\u3092\u6C42\u3081\u308B\u30AF\u30ED\u30FC\u30B8\u30E3\
    \u30FC\u3092\u8FD4\u3059\n\n// UnionFind\nfunc NewUnionFind(n int) *UnionFind\
    \ {\n\troot := make([]int, n)\n\tsize := make([]int, n)\n\tfor i := 0; i < n;\
    \ i++ {\n\t\troot[i] = i\n\t\tsize[i] = 1\n\t}\n\tuf := &UnionFind{root: root,\
    \ size: size, group: n}\n\treturn uf\n}\nfunc (uf *UnionFind) Union(p int, q int)\
    \ {\n\tqRoot := uf.Root(q)\n\tpRoot := uf.Root(p)\n\tif qRoot == pRoot {\n\t\t\
    return\n\t}\n\tuf.group--\n\tif uf.size[qRoot] < uf.size[pRoot] {\n\t\tuf.root[qRoot]\
    \ = uf.root[pRoot]\n\t\tuf.size[pRoot] += uf.size[qRoot]\n\t} else {\n\t\tuf.root[pRoot]\
    \ = uf.root[qRoot]\n\t\tuf.size[qRoot] += uf.size[pRoot]\n\t}\n}\nfunc (uf *UnionFind)\
    \ Root(p int) int {\n\tif p > len(uf.root)-1 {\n\t\treturn -1\n\t}\n\tfor uf.root[p]\
    \ != p {\n\t\tuf.root[p] = uf.root[uf.root[p]]\n\t\tp = uf.root[p]\n\t}\n\treturn\
    \ p\n}\nfunc (uf *UnionFind) find(p int) int {\n\treturn uf.Root(p)\n}\nfunc (uf\
    \ *UnionFind) Connected(p int, q int) bool {\n\treturn uf.Root(p) == uf.Root(q)\n\
    }\nfunc (uf *UnionFind) Groups() map[int]int {\n\tcm := make(map[int]int)\n\t\
    for i := 0; i < len(uf.root); i++ {\n\t\tt := uf.find(uf.root[i])\n\t\tcm[t]++\n\
    \t}\n\treturn cm\n}\n\n// Int\u578BDeque\nfunc NewDeque() *Deque {\n\tdeq := &Deque{}\n\
    \tdeq.deq = list.New()\n\treturn deq\n}\nfunc (deq *Deque) PushBack(n int) {\n\
    \tdeq.cur = deq.deq.PushBack(n)\n\tdeq.sum += n\n}\nfunc (deq *Deque) PushFront(n\
    \ int) {\n\tdeq.cur = deq.deq.PushFront(n)\n\tdeq.sum += n\n}\nfunc (deq *Deque)\
    \ InsertAfterCurrent(n int) {\n\tdeq.cur = deq.deq.InsertAfter(n, deq.cur)\n\t\
    deq.sum += n\n}\nfunc (deq *Deque) InsertBeforeCurrent(n int) {\n\tdeq.cur = deq.deq.InsertBefore(n,\
    \ deq.cur)\n\tdeq.sum += n\n}\nfunc (deq *Deque) PopBack() int {\n\tret := deq.deq.Back().Value.(int)\n\
    \tdeq.deq.Remove(deq.deq.Back())\n\tdeq.sum -= ret\n\treturn ret\n}\nfunc (deq\
    \ *Deque) PopFront() int {\n\tret := deq.deq.Front().Value.(int)\n\tdeq.deq.Remove(deq.deq.Front())\n\
    \tdeq.sum -= ret\n\treturn ret\n}\nfunc (deq *Deque) DumpBack() []int {\n\tret\
    \ := make([]int, 0, deq.Len())\n\tfor e := deq.deq.Back(); e != nil; e = e.Prev()\
    \ {\n\t\tret = append(ret, e.Value.(int))\n\t}\n\treturn ret\n}\nfunc (deq *Deque)\
    \ DumpFront() []int {\n\tret := make([]int, 0, deq.Len())\n\tfor e := deq.deq.Front();\
    \ e != nil; e = e.Next() {\n\t\tret = append(ret, e.Value.(int))\n\t}\n\treturn\
    \ ret\n}\nfunc (deq *Deque) Back() int {\n\treturn deq.deq.Back().Value.(int)\n\
    }\nfunc (deq *Deque) Front() int {\n\treturn deq.deq.Front().Value.(int)\n}\n\
    func (deq *Deque) Len() int {\n\treturn deq.deq.Len()\n}\n\n// int\u578B\u30D2\
    \u30FC\u30D7\nfunc NewIntPQ(prioritySmallest bool) *IntPQ {\n\tret := &IntPQ{}\n\
    \tret.d = make([]int, 0, 100)\n\tn := ret.Len()\n\tfor i := n/2 - 1; i >= 0; i--\
    \ {\n\t\tret.down(i, n)\n\t}\n\tret.prioritySmallest = prioritySmallest\n\treturn\
    \ ret\n}\nfunc (pq *IntPQ) less(i, j int) bool {\n\tif pq.prioritySmallest ==\
    \ true {\n\t\treturn pq.d[i] < pq.d[j]\n\t} else {\n\t\treturn pq.d[i] > pq.d[j]\n\
    \t}\n}\nfunc (pq *IntPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i]\
    \ }\nfunc (pq *IntPQ) down(i0, n int) bool {\n\ti := i0\n\tfor {\n\t\tj1 := 2*i\
    \ + 1\n\t\tif j1 >= n || j1 < 0 { // j1 < 0 after int overflow\n\t\t\tbreak\n\t\
    \t}\n\t\tj := j1 // left child\n\t\tif j2 := j1 + 1; j2 < n && pq.less(j2, j1)\
    \ {\n\t\t\tj = j2 // = 2*i + 2  // right child\n\t\t}\n\t\tif !pq.less(j, i) {\n\
    \t\t\tbreak\n\t\t}\n\t\tpq.swap(i, j)\n\t\ti = j\n\t}\n\treturn i > i0\n}\nfunc\
    \ (pq *IntPQ) up(j int) {\n\tfor {\n\t\ti := (j - 1) / 2 // parent\n\t\tif i ==\
    \ j || !pq.less(j, i) {\n\t\t\tbreak\n\t\t}\n\t\tpq.swap(i, j)\n\t\tj = i\n\t\
    }\n}\nfunc (pq *IntPQ) Push(x int) {\n\tpq.d = append(pq.d, x)\n\tpq.up(len(pq.d)\
    \ - 1)\n}\nfunc (pq *IntPQ) Pop() int {\n\tn := pq.Len() - 1\n\tx := pq.d[0]\n\
    \tpq.swap(0, n)\n\tpq.down(0, n)\n\tpq.d = pq.d[0 : pq.Len()-1]\n\treturn x\n\
    }\nfunc (pq *IntPQ) Len() int {\n\treturn len(pq.d)\n}\nfunc (pq *IntPQ) Peek()\
    \ int {\n\treturn pq.d[0]\n}\nfunc (pq *IntPQ) PushSlice(a []int) {\n\tfor _,\
    \ v := range a {\n\t\tpq.Push(v)\n\t}\n}\nfunc (pq *IntPQ) Remove(i int) int {\n\
    \tn := pq.Len() - 1\n\tif n != i {\n\t\tpq.swap(i, n)\n\t\tif !pq.down(i, n) {\n\
    \t\t\tpq.up(i)\n\t\t}\n\t}\n\treturn pq.Pop()\n}\nfunc (pq *IntPQ) Fix(i int)\
    \ {\n\tif !pq.down(i, pq.Len()) {\n\t\tpq.up(i)\n\t}\n}\nfunc (pq *IntPQ) PopAndPush(x\
    \ int) int {\n\tret := pq.Peek()\n\tpq.d[0] = x\n\tpq.down(0, pq.Len())\n\treturn\
    \ ret\n}\nfunc (pq *IntPQ) PopUniq() int {\n\tret := pq.Pop()\n\tfor pq.Len()\
    \ != 0 && pq.Peek() == ret {\n\t\tpq.Pop()\n\t}\n\treturn ret\n}\n\n// \u69CB\u9020\
    \u4F53Deque\nfunc NewDataDeque() *DataDeque {\n\tdeq := &DataDeque{}\n\tdeq.deq\
    \ = list.New()\n\treturn deq\n}\nfunc (deq *DataDeque) PushBack(d Data) {\n\t\
    deq.cur = deq.deq.PushBack(d)\n}\nfunc (deq *DataDeque) PushFront(d Data) {\n\t\
    deq.cur = deq.deq.PushFront(d)\n}\nfunc (deq *DataDeque) InsertAfterCurrent(d\
    \ Data) {\n\tdeq.cur = deq.deq.InsertAfter(d, deq.cur)\n}\nfunc (deq *DataDeque)\
    \ InsertBeforeCurrent(d Data) {\n\tdeq.cur = deq.deq.InsertBefore(d, deq.cur)\n\
    }\nfunc (deq *DataDeque) PopBack() Data {\n\tret := deq.deq.Back().Value.(Data)\n\
    \tdeq.deq.Remove(deq.deq.Back())\n\treturn ret\n}\nfunc (deq *DataDeque) PopFront()\
    \ Data {\n\tret := deq.deq.Front().Value.(Data)\n\tdeq.deq.Remove(deq.deq.Front())\n\
    \treturn ret\n}\nfunc (deq *DataDeque) DumpBack() []Data {\n\tret := make([]Data,\
    \ 0, deq.Len())\n\tfor e := deq.deq.Back(); e != nil; e = e.Prev() {\n\t\tret\
    \ = append(ret, e.Value.(Data))\n\t}\n\treturn ret\n}\nfunc (deq *DataDeque) DumpFront()\
    \ []Data {\n\tret := make([]Data, 0, deq.Len())\n\tfor e := deq.deq.Front(); e\
    \ != nil; e = e.Next() {\n\t\tret = append(ret, e.Value.(Data))\n\t}\n\treturn\
    \ ret\n}\nfunc (deq *DataDeque) Back() Data {\n\treturn deq.deq.Back().Value.(Data)\n\
    }\nfunc (deq *DataDeque) Front() Data {\n\treturn deq.deq.Front().Value.(Data)\n\
    }\nfunc (deq *DataDeque) Len() int {\n\treturn deq.deq.Len()\n}\n\n// \u30C7\u30FC\
    \u30BF\u69CB\u9020\u4F53\u7528\u30D2\u30FC\u30D7\nfunc NewItemPQ(prioritySmallest\
    \ bool) *ItemPQ {\n\tret := &ItemPQ{}\n\tret.d = make([]Item, 0, 100)\n\tn :=\
    \ ret.Len()\n\tfor i := n/2 - 1; i >= 0; i-- {\n\t\tret.down(i, n)\n\t}\n\tret.prioritySmallest\
    \ = prioritySmallest\n\treturn ret\n}\nfunc (pq *ItemPQ) less(i, j int) bool {\n\
    \tif pq.prioritySmallest == true {\n\t\treturn pq.d[i].key < pq.d[j].key\n\t}\
    \ else {\n\t\treturn pq.d[i].key > pq.d[j].key\n\t}\n}\nfunc (pq *ItemPQ) swap(i,\
    \ j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }\nfunc (pq *ItemPQ) down(i0, n\
    \ int) bool {\n\ti := i0\n\tfor {\n\t\tj1 := 2*i + 1\n\t\tif j1 >= n || j1 < 0\
    \ { // j1 < 0 after int overflow\n\t\t\tbreak\n\t\t}\n\t\tj := j1 // left child\n\
    \t\tif j2 := j1 + 1; j2 < n && pq.less(j2, j1) {\n\t\t\tj = j2 // = 2*i + 2  //\
    \ right child\n\t\t}\n\t\tif !pq.less(j, i) {\n\t\t\tbreak\n\t\t}\n\t\tpq.swap(i,\
    \ j)\n\t\ti = j\n\t}\n\treturn i > i0\n}\nfunc (pq *ItemPQ) Push(x Item) {\n\t\
    pq.d = append(pq.d, x)\n\tpq.up(len(pq.d) - 1)\n}\nfunc (pq *ItemPQ) up(j int)\
    \ {\n\tfor {\n\t\ti := (j - 1) / 2 // parent\n\t\tif i == j || !pq.less(j, i)\
    \ {\n\t\t\tbreak\n\t\t}\n\t\tpq.swap(i, j)\n\t\tj = i\n\t}\n}\nfunc (pq *ItemPQ)\
    \ Pop() Item {\n\tn := pq.Len() - 1\n\tx := pq.d[0]\n\tpq.swap(0, n)\n\tpq.down(0,\
    \ n)\n\tpq.d = pq.d[0 : pq.Len()-1]\n\treturn x\n}\nfunc (pq *ItemPQ) Len() int\
    \   { return len(pq.d) }\nfunc (pq *ItemPQ) Peek() Item { return pq.d[0] }\nfunc\
    \ (pq *ItemPQ) Remove(i int) Item {\n\tn := pq.Len() - 1\n\tif n != i {\n\t\t\
    pq.swap(i, n)\n\t\tif !pq.down(i, n) {\n\t\t\tpq.up(i)\n\t\t}\n\t}\n\treturn pq.Pop()\n\
    }\nfunc (pq *ItemPQ) Fix(i int) {\n\tif !pq.down(i, pq.Len()) {\n\t\tpq.up(i)\n\
    \t}\n}\n\n// \u6574\u6570\u95A2\u9023\nfunc lcm(a, b int) int {\n\treturn (a /\
    \ gcd(a, b)) * b\n}\nfunc gcd(a, b int) int {\n\tif b == 0 {\n\t\treturn a\n\t\
    }\n\tc := 1\n\tfor c != 0 {\n\t\tc = a % b\n\t\ta, b = b, c\n\t}\n\treturn a\n\
    }\nfunc exgcd(a, b int) (int, int) {\n\tq := 0\n\tx0, x1, y0, y1 := 1, 0, 0, 1\n\
    \tfor b != 0 {\n\t\tq, a, b = a/b, b, a%b\n\t\tx0, x1 = x1, x0-q*x1\n\t\ty0, y1\
    \ = y1, y0-q*y1\n\t}\n\treturn x0, y0\n}\nfunc divisorList(n int) []int {\n\t\
    var l []int\n\tfor i := 1; i*i <= n; i++ {\n\t\tif n%i == 0 {\n\t\t\tl = append(l,\
    \ i)\n\t\t\tif i != n/i {\n\t\t\t\tl = append(l, n/i)\n\t\t\t}\n\t\t}\n\t}\n\t\
    sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })\n\treturn l\n}\nfunc\
    \ divisorCnt(maxNum int) func(int) int {\n\tmemo := intSlice(maxNum, -1)\n\tf\
    \ := func(n int) int {\n\t\tif memo[n] != -1 {\n\t\t\treturn memo[n]\n\t\t}\n\t\
    \tvar cnt int\n\t\tfor i := 1; i*i <= n; i++ {\n\t\t\tif n%i == 0 {\n\t\t\t\t\
    cnt++\n\t\t\t\tif i != n/i {\n\t\t\t\t\tcnt++\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t\t\
    memo[n] = cnt\n\t\treturn cnt\n\t}\n\treturn f\n}\nfunc divisorPairs(n int) []Pair\
    \ {\n\tvar p []Pair\n\tfor i := 1; i*i <= n; i++ {\n\t\tif n%i == 0 {\n\t\t\t\
    p = append(p, Pair{i, n / i})\n\t\t}\n\t}\n\treturn p\n}\nfunc factorization(n\
    \ int) map[int]int {\n\tm := make(map[int]int)\n\tfor i := 2; i*i <= n; i++ {\n\
    \t\tfor n%i == 0 {\n\t\t\tm[i]++\n\t\t\tn = n / i\n\t\t}\n\t}\n\tif n != 0 &&\
    \ n != 1 {\n\t\tm[n]++\n\t}\n\treturn m\n}\nfunc fastFacorization(n int) func(x\
    \ int) map[int]int {\n\trp := make([]int, n+1)\n\tp := primeList(n)\n\trp[1] =\
    \ 1\n\tfor _, v := range p {\n\t\tfor i := v; i <= n; i += v {\n\t\t\trp[i] =\
    \ v\n\t\t}\n\t}\n\tvar f func(int) map[int]int\n\tf = func(k int) map[int]int\
    \ {\n\t\tt := k\n\t\tm := make(map[int]int)\n\t\tfor t != 1 {\n\t\t\tm[rp[t]]++\n\
    \t\t\tt = t / rp[t]\n\t\t}\n\t\treturn m\n\t}\n\treturn f\n}\nfunc isPrime(n int)\
    \ bool { return big.NewInt(int64(n)).ProbablyPrime(0) }\nfunc primeList(n int)\
    \ []int {\n\tif n < 2 {\n\t\treturn nil\n\t} else if n == 2 {\n\t\treturn []int{2}\n\
    \t}\n\tl := make([]int, n+1)\n\tprimes := make([]int, 0, n)\n\tfor i := 4; i <=\
    \ n; i += 2 {\n\t\tl[i] = 1\n\t}\n\tfor i := 3; i <= n; i += 2 {\n\t\tif l[i]\
    \ == 1 {\n\t\t\tcontinue\n\t\t}\n\t\tfor j := i * 2; j <= n; j += i {\n\t\t\t\
    l[j] = 1\n\t\t}\n\t}\n\tprimes = append(primes, 2)\n\tfor i := 3; i <= n; i +=\
    \ 2 {\n\t\tif l[i] == 0 {\n\t\t\tprimes = append(primes, i)\n\t\t}\n\t}\n\treturn\
    \ primes\n}\nfunc powMod(x, k, m int) int {\n\tif k == 0 {\n\t\treturn 1\n\t}\n\
    \tif x > m {\n\t\tx %= m\n\t}\n\tif k%2 == 0 {\n\t\treturn powMod(x*x%m, k/2,\
    \ m)\n\t} else {\n\t\treturn x * powMod(x, k-1, m) % m\n\t}\n}\nfunc pow(a, b\
    \ int) int {\n\tres := 1\n\tfor b > 0 {\n\t\tif b&1 == 1 {\n\t\t\tres = res *\
    \ a\n\t\t}\n\t\ta = a * a\n\t\tb >>= 1\n\t}\n\treturn res\n}\nfunc combination(n,\
    \ r int) int {\n\tif n < r {\n\t\treturn 0\n\t}\n\tr = min(n-r, r)\n\td := make([]int,\
    \ r)\n\n\tfor i := 0; i < r; i++ {\n\t\td[i] = n - i\n\t}\n\tfor i := 2; i <=\
    \ r; i++ {\n\t\tti := i\n\t\tfor j := 0; j < r; j++ {\n\t\t\tg := gcd(d[j], ti)\n\
    \t\t\tif g != 1 {\n\t\t\t\tti /= g\n\t\t\t\td[j] /= g\n\t\t\t}\n\t\t\tif ti ==\
    \ 1 {\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t}\n\tret := 1\n\tfor i := 0; i < r; i++\
    \ {\n\t\tret *= d[i]\n\t}\n\treturn ret\n}\nfunc combinationMod(a, b int) int\
    \ {\n\tt1, t2 := 1, 1\n\tfor i := 2; i <= b; i++ {\n\t\tt2 = (i * t2) % MOD\n\t\
    }\n\tinv := modInv(t2)\n\tfor i := a - b + 1; i <= a; i++ {\n\t\tt1 = (i * t1)\
    \ % MOD\n\t}\n\treturn (t1 * inv) % MOD\n}\nfunc sqrt(x int) int             \
    \     { return int(math.Sqrt(float64(x))) }\nfunc cbrt(x int) int            \
    \      { return int(math.Cbrt(float64(x))) }\nfunc seqSum(x int) int         \
    \       { return (x*x + x) / 2 }\nfunc seqRangeSum(from, to int) int    { return\
    \ seqSum(to) - seqSum(from-1) }\nfunc seqSumMod(x int) int             { return\
    \ modMul(modAdd(modMul(x%MOD, x%MOD), x%MOD), modInv(2)) }\nfunc seqRangeSumMod(from,\
    \ to int) int { return modSub(seqSum(to), seqSum(from-1)) }\nfunc getArithmeticSeqRange(d,\
    \ a, maxY, maxX int) (int, int) {\n\tt := (-a + d - 1) / d\n\treturn max(1, t),\
    \ min(maxX, (maxY-a)/d)\n} // \u7B49\u5DEE\u6570\u5217(\u521D\u9805a,\u516C\u5DEE\
    d)\u306B\u3064\u3044\u3066\u3001y[0,maxY],x[0,maxX]\u306E\u6761\u4EF6\u3067x\u306E\
    \u7BC4\u56F2\u3092\u6C42\u3081\u308B\nfunc MulMatrix(m1, m2 [][]int, mod int)\
    \ [][]int {\n\tsize := len(m1)\n\tret := make([][]int, size)\n\tfor i := 0; i\
    \ < size; i++ {\n\t\tret[i] = make([]int, size)\n\t}\n\tfor i := 0; i < size;\
    \ i++ {\n\t\tfor j := 0; j < size; j++ {\n\t\t\tfor k := 0; k < size; k++ {\n\t\
    \t\t\tret[i][k] += m1[i][j] * m2[j][k]\n\t\t\t\tret[i][k] %= mod\n\t\t\t}\n\t\t\
    }\n\t}\n\treturn ret\n} // \u6B63\u65B9\u884C\u5217\u306E\u639B\u3051\u7B97\n\
    func PowMatrix(m [][]int, n int, mod int) [][]int {\n\t// \u884C\u5217\u7D2F\u4E57\
    \n\t// a_0=s a_1=t\n\t// a_(i+2) = p*a_(i+1) + q*a_i\n\t//           \u2193\n\t\
    // |a_(i+2)| = | p q |^n (a_1)\n\t// |a_(i+1)| = | 1 0 |   (a_0)\n\tsize := len(m)\n\
    \ttm := make([][]int, size)\n\tfor i := 0; i < size; i++ {\n\t\ttm[i] = make([]int,\
    \ size)\n\t\ttm[i][i] = 1\n\t}\n\tfor n >= 2 {\n\t\tif n%2 == 0 {\n\t\t\tm = MulMatrix(m,\
    \ m, mod)\n\t\t\tn = n / 2\n\t\t} else {\n\t\t\ttm = MulMatrix(tm, m, mod)\n\t\
    \t\tm = MulMatrix(m, m, mod)\n\t\t\tn = (n - 1) / 2\n\t\t}\n\t}\n\tm = MulMatrix(tm,\
    \ m, mod)\n\treturn m\n} // \u6B63\u65B9\u884C\u5217\u306E\u51AA\u4E57\n\n// int\u578B\
    \u64CD\u4F5C(1-indexed)\nfunc getDigitLen(x int) int   { return condInt(x <= 0,\
    \ 0, 1+int(math.Log10(float64(x)))) }\nfunc getDigit(x, idx int) int { return\
    \ condInt(idx > 17 || idx < 1, 0, (x/p10[idx-1])%10) }\nfunc setDigit(x, idx,\
    \ v int) int {\n\treturn condInt(idx > 17 || v > 9 || v < 0, 0, x+(v-getDigit(x,\
    \ idx))*p10[idx-1])\n}\nfunc swapDigit(x, idx1, idx2 int) int {\n\tt1 := getDigit(x,\
    \ idx1)\n\tt2 := getDigit(x, idx2)\n\tx = setDigit(x, idx2, t1)\n\tx = setDigit(x,\
    \ idx1, t2)\n\treturn x\n}\n\n// \u30D3\u30C3\u30C8\u64CD\u4F5C(0-indexed)\nfunc\
    \ bitLen(x int) int              { return bits.Len(uint(x)) }                \
    \                        // x\u306E2\u9032\u6570\u3067\u306E\u9577\u3055\u3092\
    \u53D6\u5F97\u3059\u308B\nfunc getNthBit(x, idx int) int      { return condInt(x&(1<<uint(idx))\
    \ == 0, 0, 1) }                     // idx\u756A\u76EE\u306Ebit\u3092\u53D6\u5F97\
    \u3059\u308B\nfunc setNthBit(x, idx, bit int) int { return condInt(bit == 0, x\
    \ & ^(1<<uint(idx)), x|(1<<uint(idx))) } // idx\u756A\u76EE\u306Ebit\u3092\u5909\
    \u66F4\u3059\u308B\nfunc toggleNthBit(x, idx int) int   { return x ^ (1 << uint(idx))\
    \ }                                     // idx\u756A\u76EE\u306Ebit\u3092\u53CD\
    \u8EE2\u3055\u305B\u308B\nfunc itobstr(x, digits int) string  { return fmt.Sprintf(\"\
    %0*b\", digits, x) }                           // \u6570\u5024\u3092\u30D0\u30A4\
    \u30CA\u30EA\u8868\u8A18\u3057\u305Fstring\u306B\u5909\u63DB\u3059\u308B(\u30C7\
    \u30D0\u30C3\u30B0\u7528)\nfunc btoi(b []byte) int {\n\tret, _ := strconv.ParseInt(string(b),\
    \ 2, 64)\n\treturn int(ret)\n}                                // \u30D0\u30A4\u30CA\
    \u30EA\u8868\u8A18[]byte\u3092int\u306B\u5909\u63DB\u3059\u308B\nfunc bitRightmostOne(n\
    \ int) int  { return bits.TrailingZeros(uint(n)) }  // bit\u304C1\u306B\u306A\u3063\
    \u3066\u3044\u308B\u4E00\u756A\u53F3\u5074\u306E\u4F4D\u7F6E\u3092\u8FD4\u3059\
    \nfunc bitRightmostZero(n int) int { return bits.TrailingZeros(^uint(n)) } //\
    \ bit\u304C0\u306B\u306A\u3063\u3066\u3044\u308B\u4E00\u756A\u53F3\u5074\u306E\
    \u4F4D\u7F6E\u3092\u8FD4\u3059\nfunc getBitPosConv(mask int) []int {\n\tt := make([]int,\
    \ bits.OnesCount(uint(mask)))\n\tk := bits.Len(uint(mask))\n\tcnt := 0\n\tfor\
    \ i := 0; i < k; i++ {\n\t\tif getNthBit(mask, i) == 1 {\n\t\t\tt[cnt] = i\n\t\
    \t\tcnt++\n\t\t}\n\t}\n\treturn t\n} // bit\u304C1\u306B\u306A\u3063\u3066\u3044\
    \u308B\u4F4D\u7F6E\u306E\u4E00\u89A7\u3092\u8FD4\u3059\nfunc nbitsNth(bit0Sum,\
    \ bit1Sum, kth int) int {\n\tret := 0\n\tt := 0\n\ta, b := bit0Sum, bit1Sum\n\t\
    s := bit0Sum + bit1Sum\n\tfor i := 0; i < s; i++ {\n\t\tif a <= 0 {\n\t\t\tret\
    \ = setNthBit(ret, s-i-1, 1)\n\t\t} else if b <= 0 {\n\t\t\tret = setNthBit(ret,\
    \ s-i-1, 0)\n\t\t} else {\n\t\t\tn := combination(a+b-1, a-1)\n\t\t\tif n+t <\
    \ kth {\n\t\t\t\tt += n\n\t\t\t\tret = setNthBit(ret, s-i-1, 1)\n\t\t\t\tb--\n\
    \t\t\t} else {\n\t\t\t\tret = setNthBit(ret, s-i-1, 0)\n\t\t\t\ta--\n\t\t\t}\n\
    \t\t}\n\t}\n\treturn ret\n} // 2\u9032\u6570\u3067bit\u306E0\u30681\u3092\u6307\
    \u5B9A\u3057\u305F\u6570\u5206\u542B\u3080\u5024\u3064\u3044\u3066kth\u756A\u76EE\
    \u306E\u5024\u3092\u6C42\u3081\u308B\nfunc bitApply(bitLen, num, bit0, bit1 int)\
    \ int {\n\tret, t := 0, 0\n\tfor i := 0; i < bitLen; i++ {\n\t\tk := getNthBit(num,\
    \ i)\n\t\tif k == 0 {\n\t\t\tt = getNthBit(bit0, i)\n\t\t} else {\n\t\t\tt = getNthBit(bit1,\
    \ i)\n\t\t}\n\t\tret = setNthBit(ret, i, t)\n\t}\n\treturn ret\n} // num\u3092\
    2\u9032\u6570\u306B\u5909\u63DB\u3057\u305F\u5404\u6841\u306B\u3064\u3044\u3066\
    \u30010\u306E\u5834\u5408\u306Fbit0,1\u306E\u5834\u5408\u306Fbit1\u3067\u5BFE\u5FDC\
    \u3059\u308B\u6841\u306E\u5024\u306B\u7F6E\u304D\u63DB\u3048\u308B\nfunc bitIndexSum(a\
    \ []int, x int) int {\n\tret := 0\n\tfor i := 0; i < bitLen(x); i++ {\n\t\tif\
    \ getNthBit(x, i) == 1 {\n\t\t\tret += a[i]\n\t\t}\n\t}\n\treturn ret\n} // a\u306B\
    \u3064\u3044\u3066x\u306Ebit\u304C1\u306Epos\u306E\u548C\u3092\u53D6\u308B\nfunc\
    \ countRangeBit(x, from, to int) int {\n\tt := (1<<(to-from+1) - 1) << from\n\t\
    return bits.OnesCount(uint(t & x))\n} // [l,r]\u306B\u542B\u307E\u308C\u308Bbit1\u306E\
    \u6570\u3092\u30AB\u30A6\u30F3\u30C8\u3059\u308B(0-indexed)\n\n// Pair\u64CD\u4F5C\
    \nfunc subPair(p, t Pair) Pair     { return Pair{p.a - t.a, p.b - t.b} }\nfunc\
    \ addPair(p, t Pair) Pair     { return Pair{p.a + t.a, p.b + t.b} }\nfunc mulPair(p\
    \ Pair, d int) Pair { return Pair{p.a * d, p.b * d} }\nfunc detPair(p, t Pair)\
    \ int      { return p.a*t.b - p.b*t.a }\nfunc dotPair(p, t Pair) int      { return\
    \ p.a*t.a + p.b*t.b }\nfunc distPair(p, t Pair) int     { return dotPair(subPair(p,\
    \ t), subPair(p, t)) }\nfunc distIntMatrix(p []Pair) [][]int {\n\tn := len(p)\n\
    \tret := intSlice2(n, n, INF)\n\tfor i := 0; i < n-1; i++ {\n\t\tfor j := i +\
    \ 1; j < n; j++ {\n\t\t\tt := distPair(p[i], p[j])\n\t\t\tret[i][j] = t\n\t\t\t\
    ret[j][i] = t\n\t\t}\n\t}\n\treturn ret\n}\nfunc distFloatMatrix(p []Pair) [][]float64\
    \ {\n\tn := len(p)\n\tret := make([][]float64, n)\n\tfor i := 0; i < n; i++ {\n\
    \t\tret[i] = make([]float64, n)\n\t\tfor j := 0; j < n; j++ {\n\t\t\tret[i][j]\
    \ = math.MaxFloat32\n\t\t}\n\t}\n\tfor i := 0; i < n-1; i++ {\n\t\tfor j := i\
    \ + 1; j < n; j++ {\n\t\t\tt := math.Sqrt(float64(distPair(p[i], p[j])))\n\t\t\
    \tret[i][j] = t\n\t\t\tret[j][i] = t\n\t\t}\n\t}\n\treturn ret\n}\nfunc intersectSum(p1,\
    \ p2 Pair) int {\n\treturn max(0, min(p1.b, p2.b)-max(p1.a, p2.a))\n} // [a1,b1),[a2,b2)\u306E\
    \u91CD\u8907\u533A\u9593\u306E\u30AB\u30A6\u30F3\u30C8\nfunc normalizePairs(p\
    \ []Pair) []Pair {\n\tma, mb := INF, INF\n\tfor i := 0; i < len(p); i++ {\n\t\t\
    ma = min(ma, p[i].a)\n\t\tmb = min(mb, p[i].b)\n\t}\n\tret := make([]Pair, len(p))\n\
    \tfor i := 0; i < len(p); i++ {\n\t\tret[i].a = p[i].a - ma\n\t\tret[i].b = p[i].b\
    \ - mb\n\t}\n\treturn ret\n}\nfunc addPairs(p []Pair, d Pair) []Pair {\n\tret\
    \ := make([]Pair, len(p))\n\tfor i := 0; i < len(p); i++ {\n\t\tret[i] = addPair(p[i],\
    \ d)\n\t}\n\treturn ret\n}\nfunc sumProductOfPairs(p []Pair) int {\n\tret := 0\n\
    \tfor _, v := range p {\n\t\tret += v.a * v.b\n\t}\n\treturn ret\n}\nfunc mergeIntervals(p\
    \ []Pair) []Pair {\n\tsortPairs(p, true, true, true)\n\tpa, pb := p[0].a, p[0].b\n\
    \tret := make([]Pair, 0)\n\tfor i := 0; i < len(p); i++ {\n\t\tif pb >= p[i].a\
    \ {\n\t\t\tchmax(&pb, p[i].b)\n\t\t} else {\n\t\t\tret = append(ret, Pair{pa,\
    \ pb})\n\t\t\tpa, pb = p[i].a, p[i].b\n\t\t}\n\t\tif i == len(p)-1 {\n\t\t\tret\
    \ = append(ret, Pair{pa, pb})\n\t\t}\n\t}\n\treturn ret\n} // \u8907\u6570\u306E\
    [a,b)\u306E\u533A\u9593\u306B\u3064\u3044\u3066\u91CD\u8907\u3059\u308B\u533A\u9593\
    \u3092\u30DE\u30FC\u30B8\u3059\u308B\n\n// \u30B0\u30EB\u30FC\u30D4\u30F3\u30B0\
    \nfunc splitGroupIndexRange(a []int) []Pair {\n\td := 1 //\u5DEE\u5206\u304Cd\u4EE5\
    \u4E0B\u3060\u3068\u540C\u3058\u30B0\u30EB\u30FC\u30D7\n\tif len(a) == 0 {\n\t\
    \treturn nil\n\t}\n\tif len(a) == 1 {\n\t\treturn []Pair{{0, 0}}\n\t}\n\tbegin\
    \ := 0\n\tret := make([]Pair, 0)\n\tfor i := 1; i < len(a); i++ {\n\t\tif a[i]-a[i-1]\
    \ <= d {\n\t\t\tcontinue\n\t\t}\n\t\tret = append(ret, Pair{begin, i - 1})\n\t\
    \tbegin = i\n\t}\n\tret = append(ret, Pair{begin, len(a) - 1})\n\treturn ret\n\
    } // \u5358\u8ABF\u5897\u52A0\u306E\u30B9\u30E9\u30A4\u30B9\u3067\u96A3\u5408\u3046\
    \u8981\u7D20\u306E\u5DEE\u304Cd\u4EE5\u4E0B\u306E\u3082\u306E\u3092\u30B0\u30EB\
    \u30FC\u30D4\u30F3\u30B0\u3057\u958B\u59CB/\u7D42\u4E86\u306Eindex\u3092\u8FD4\
    \u3059\nfunc divideConnectedGroup(n int, p []Pair) [][]int {\n\tuf := NewUnionFind(n\
    \ + 1)\n\tfor i := 0; i < len(p); i++ {\n\t\tuf.Union(p[i].a, p[i].b)\n\t}\n\t\
    nm := make(map[int][]int)\n\tfor i := 1; i <= n; i++ {\n\t\tk := uf.find(i)\n\t\
    \tnm[k] = append(nm[k], i)\n\t}\n\tret := make([][]int, 0)\n\tfor _, v := range\
    \ nm {\n\t\tret = append(ret, v)\n\t}\n\treturn ret\n} // \u7121\u5411\u30B0\u30E9\
    \u30D5\u3092\u9023\u7D50\u6210\u5206\u6BCE\u306B\u30B0\u30EB\u30FC\u30D4\u30F3\
    \u30B0\u3059\u308B (1-indexed\u3067\u30CE\u30FC\u30C90\u306F\u7121\u8996)\n\n\
    // \u6570\u3048\u3042\u3052\nfunc findSumPattern(a []int, mod int, maxCnt int)\
    \ [][][]int {\n\tm := make([][][]int, mod)\n\tcnt := 0\n\tvar f func(int, int,\
    \ []int)\n\tf = func(idx, val int, t []int) {\n\t\tif idx == len(a) || (cnt >=\
    \ maxCnt && maxCnt != -1) {\n\t\t\treturn\n\t\t}\n\t\tnv := (val + a[idx]) % mod\n\
    \t\tnp := make([]int, len(t))\n\t\tcopy(np, t)\n\t\tnp = append(np, idx+1)\n\t\
    \tm[nv] = append(m[nv], np)\n\t\tcnt++\n\t\tf(idx+1, val, t)\n\t\tf(idx+1, nv,\
    \ np)\n\t\treturn\n\t}\n\tf(0, 0, nil)\n\treturn m\n} // \u30B9\u30E9\u30A4\u30B9\
    \u306E\u90E8\u5206\u5217\u306E\u548C(mod)\u306B\u3064\u3044\u3066\u3001\u7D44\u307F\
    \u5408\u308F\u305B\u3092\u6700\u5927maxCnt\u56DE\u8A18\u9332\u3059\u308B(maxCnt=-1\u306E\
    \u5834\u5408\u306F\u5168\u5217\u6319)\n\n// \u5206\u6570\nfunc compareFrac(a,\
    \ b Frac) int {\n\t// -1(a>b), 1(a<b), 0(a==b)\n\tt1 := a.top * b.bottom\n\tt2\
    \ := a.bottom * b.top\n\tif t1 == t2 {\n\t\treturn 0\n\t} else if t1 < t2 {\n\t\
    \treturn 1\n\t} else {\n\t\treturn -1\n\t}\n}\nfunc AddFrac(a, b Frac) Frac {\n\
    \ttop := a.top*b.bottom + a.bottom*b.top\n\tbottom := a.bottom * b.bottom\n\t\
    g := gcd(top, bottom)\n\treturn Frac{top / g, bottom / g}\n}\nfunc SubFrac(a,\
    \ b Frac) Frac {\n\ttop := a.top*b.bottom - a.bottom*b.top\n\tbottom := a.bottom\
    \ * b.bottom\n\tg := gcd(top, bottom)\n\treturn Frac{top / g, bottom / g}\n}\n\
    func MulFrac(a, b Frac) Frac {\n\ttop := a.top * b.top\n\tbottom := a.bottom *\
    \ b.bottom\n\tg := gcd(top, bottom)\n\treturn Frac{top / g, bottom / g}\n}\n\n\
    // \u30D9\u30AF\u30C8\u30EB\nfunc degreeToRadian(d float64) float64 { return d\
    \ * math.Pi / 180 }\nfunc rotatePoint(p Point, rad float64) Point {\n\tcos :=\
    \ math.Cos(rad)\n\tsin := math.Sin(rad)\n\treturn Point{sin*p.x + cos*p.y, cos*p.x\
    \ - sin*p.y}\n}\nfunc addPoint(a, b Point) Point         { return Point{a.y +\
    \ b.y, a.x + b.x} }\nfunc subPoint(a, b Point) Point         { return Point{a.y\
    \ - b.y, a.x - b.x} }\nfunc mulPoint(a Point, b float64) Point { return Point{a.y\
    \ * b, a.x * b} }\nfunc distance(a, b Point) float64 {\n\tt := subPoint(a, b)\n\
    \treturn math.Hypot(t.x, t.y)\n}\n\n// \u56F3\u5F62\u5224\u5B9A\nfunc isConnectedCircle(p1,\
    \ p2 Pair, r1, r2 int) bool {\n\td := distPair(p1, p2)\n\tif d > pow(r1+r2, 2)\
    \ || d < pow(r1-r2, 2) {\n\t\treturn false\n\t}\n\treturn true\n}\n\n// \u9806\
    \u5217\nfunc sumPermutation(s, n int) [][]int {\n\tvar results [][]int\n\tt :=\
    \ make([]int, n)\n\tvar dfs func(int, int)\n\tdfs = func(sum, index int) {\n\t\
    \tif index == n {\n\t\t\tif sum == s {\n\t\t\t\tresult := make([]int, n)\n\t\t\
    \t\tcopy(result, t)\n\t\t\t\tresults = append(results, result)\n\t\t\t}\n\t\t\t\
    return\n\t\t}\n\t\tfor i := 0; i <= s; i++ {\n\t\t\tt[index] = i\n\t\t\tdfs(sum+i,\
    \ index+1)\n\t\t}\n\t}\n\tdfs(0, 0)\n\treturn results\n} // \u5408\u8A08\u304C\
    s\u3068\u306A\u308Bn\u500B\u306E\u8981\u7D20\u306E\u30B9\u30E9\u30A4\u30B9\u306E\
    \u7D44\u307F\u5408\u308F\u305B\u3092\u8FD4\u3059\nfunc sumPermutations(s, n int)\
    \ [][]int {\n\tperm := make([][]int, 0)\n\tfor i := 0; i <= s; i++ {\n\t\tt :=\
    \ sumPermutation(i, n)\n\t\tperm = append(perm, t...)\n\t}\n\treturn perm\n} //\
    \ \u5408\u8A08\u304Cs\u4EE5\u4E0B\u3068\u306A\u308Bn\u500B\u306E\u8981\u7D20\u306E\
    \u30B9\u30E9\u30A4\u30B9\u306E\u7D44\u307F\u5408\u308F\u305B\u3092\u8FD4\u3059\
    \nfunc nextPermIterator(orderdSlice []int) func() []int {\n\tx := orderdSlice\n\
    \tfirst := true\n\tf := func(x sort.Interface) bool {\n\t\tn := x.Len() - 1\n\t\
    \tif n < 1 {\n\t\t\treturn false\n\t\t}\n\t\tj := n - 1\n\t\tfor ; !x.Less(j,\
    \ j+1); j-- {\n\t\t\tif j == 0 {\n\t\t\t\treturn false\n\t\t\t}\n\t\t}\n\t\tl\
    \ := n\n\t\tfor !x.Less(j, l) {\n\t\t\tl--\n\t\t}\n\t\tx.Swap(j, l)\n\t\tfor k,\
    \ l := j+1, n; k < l; {\n\t\t\tx.Swap(k, l)\n\t\t\tk++\n\t\t\tl--\n\t\t}\n\t\t\
    return true\n\t}\n\tf2 := func() []int {\n\t\tif first == true {\n\t\t\tfirst\
    \ = false\n\t\t\treturn x\n\t\t} else {\n\t\t\tret := f(sort.IntSlice(x))\n\t\t\
    \tif ret == false {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\treturn x\n\t\t}\n\t}\n\
    \treturn f2\n} // nextPermutation()\u306EIterator\nfunc ascPerm(size, low, high\
    \ int) [][]int {\n\tp := intSlice(size, low)\n\tret := make([][]int, 0)\n\tvar\
    \ f func(int, int)\n\tf = func(idx, cur int) {\n\t\tif idx == size {\n\t\t\tret\
    \ = append(ret, cs(p))\n\t\t\treturn\n\t\t}\n\t\tfor i := cur; i <= high; i++\
    \ {\n\t\t\tp[idx] = i\n\t\t\tf(idx+1, i)\n\t\t}\n\t}\n\tf(0, low)\n\treturn ret\n\
    } // low\u301Chigh\u306E\u8981\u7D20\u304C\u6607\u9806(ai <= aj)\u3068\u306A\u308B\
    \u3088\u3046\u306A\u7D44\u307F\u5408\u308F\u305B\u3092\u8FD4\u3059\n\n// bit set\n\
    func newBitSets(size int) *BitSets {\n\tbs := make(BitSets, size)\n\treturn &bs\n\
    }\nfunc (tb *BitSets) SetBit(i int, bitNum int, bitVal uint) {\n\t(*tb)[i].SetBit(&(*tb)[i],\
    \ bitNum, bitVal)\n}\nfunc (tb *BitSets) Mex(u int) int {\n\ttemp := new(big.Int)\n\
    \ttemp.Not(&(*tb)[u])\n\treturn int(temp.TrailingZeroBits())\n} // \u96C6\u5408\
    \u306Emex\u3092\u6C42\u3081\u308B\nfunc (tb *BitSets) Mex2(u int, v int) int {\n\
    \ttemp := new(big.Int)\n\ttemp.Or(&(*tb)[u], &(*tb)[v])\n\ttemp.Not(temp)\n\t\
    return int(temp.TrailingZeroBits())\n} //\u548C\u96C6\u5408\u306Emex\u3092\u6C42\
    \u3081\u308B\n"
  dependsOn:
  - lib/waveletmatrix/waveletmatrix.go
  - lib/waveletmatrix/waveletmatrix_test.go
  - lib/waveletmatrix/range_kth_smallest.test.go
  - lib/waveletmatrix/static_range_frequency.test.go
  isVerificationFile: false
  path: lib/template/main.go
  requiredBy:
  - lib/waveletmatrix/waveletmatrix.go
  - lib/waveletmatrix/waveletmatrix_test.go
  timestamp: '2023-08-06 22:33:19+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - lib/waveletmatrix/range_kth_smallest.test.go
  - lib/waveletmatrix/static_range_frequency.test.go
documentation_of: lib/template/main.go
layout: document
redirect_from:
- /library/lib/template/main.go
- /library/lib/template/main.go.html
title: lib/template/main.go
---
