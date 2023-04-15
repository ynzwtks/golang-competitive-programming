package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	// グラフの隣接リスト表現
	G := [][]int{
		{1, 2},
		{2, 3},
		{3},
		{},
	}
	// グラフのコスト行列
	C := [][]int{
		{2, 5},
		{3, 1},
		{4},
		{},
	}
	begin := 0

	// dijkstra関数を実行
	dist, prev := dijkstra(G, C, begin)

	// 期待される最小コストの配列
	expectedDist := []int{0, 2, 5, 3}
	if !reflect.DeepEqual(dist, expectedDist) {
		t.Errorf("Expected %v but got %v", expectedDist, dist)
	}

	// 期待される最短経路復元用の配列
	expectedPrev := []int{0, 0, 0, 1}
	if !reflect.DeepEqual(prev, expectedPrev) {
		t.Errorf("Expected %v but got %v", expectedPrev, prev)
	}
}

func TestDijkstraRestorePath(t *testing.T) {
	// 最短経路復元用の配列
	prev := []int{0, 0, 0, 1}
	end := 3

	// restorePath関数を実行
	path := restoreDijkstraPath(end, prev)

	// 期待される最短経路
	expectedPath := []int{0, 1, 3}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected %v but got %v", expectedPath, path)
	}
}
