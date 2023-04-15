package main

// Bipartite は二部グラフの構造体です。
type Bipartite struct {
	n           int
	group       map[Pair]int
	uf          *UnionFind
	isBipartite bool
}

// NewBipartite は、頂点数 n とリンク情報 l を引数に取り、
// 二部グラフの構造体を作成して返す関数です。
func NewBipartite(n int, l [][]int) *Bipartite {
	b := &Bipartite{}
	b.n = n
	b.group = make(map[Pair]int)
	b.uf = NewUnionFind(b.n * 2)
	b.isBipartite = true
	for i := 0; i < len(l); i++ {
		b.uf.Union(l[i][0], l[i][1]+b.n)
		b.uf.Union(l[i][0]+b.n, l[i][1])
	}
	for i := 0; i < b.n; i++ {
		t1 := b.uf.Root(i)
		t2 := b.uf.Root(i + b.n)
		if t1 == t2 { //二部グラフでない
			b.isBipartite = false
		}
		b.group[Pair{t1, t2}]++
	}
	return b
}

// IsBipartite は二部グラフであるかどうかを判定し、結果を返すメソッドです。
func (b *Bipartite) IsBipartite() bool {
	return b.isBipartite
}

// GetBipartiteNum は二部グラフの各グループの要素数のペアを取得し、
// Pair 型のスライスとして返すメソッドです。
func (b *Bipartite) GetBipartiteNum() []Pair {
	p := make([]Pair, 0)
	for i, v := range b.group {
		k1 := b.group[Pair{i.b, i.a}]
		k2 := v
		p = append(p, Pair{k1, k2})
	}
	return p
}
