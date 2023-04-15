package main

import (
	"reflect"
	"testing"
)

func TestCompressBFS(t *testing.T) {
	// グラフの隣接リスト表現
	G := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {},
	}
	begin := 0

	// CompressBFS関数を実行
	visited, prev := compressBFS(G, begin)

	// 期待される到達可否マップ
	expectedVisited := map[int]int{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
	}
	if !reflect.DeepEqual(visited, expectedVisited) {
		t.Errorf("Expected %v but got %v", expectedVisited, visited)
	}

	// 期待される経路情報マップ
	expectedPrev := map[int]int{
		0: -1,
		1: 0,
		2: 0,
		3: 1,
	}
	if !reflect.DeepEqual(prev, expectedPrev) {
		t.Errorf("Expected %v but got %v", expectedPrev, prev)
	}
}

func TestRestoreCompressedBFSPath(t *testing.T) {
	// 経路情報マップ
	prev := map[int]int{
		0: -1,
		1: 0,
		2: 0,
		3: 1,
	}
	begin, end := 0, 3

	// restoreCompressedBFSPath関数を実行
	path := restoreCompressedBFSPath(prev, begin, end)

	// 期待される復元されたパス
	expectedPath := []int{0, 1, 3}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected %v but got %v", expectedPath, path)
	}
}
