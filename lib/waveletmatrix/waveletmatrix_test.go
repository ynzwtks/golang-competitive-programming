package waveletmatrix

import (
	"testing"
)

func Test_waveletmatrix(t *testing.T) {
	a := []int{4, 5, 5, 5, 2, 1, 5, 6, 1, 3, 5, 0}
	wm := NewWaveletMatrix(a)

	b := []int{1, 10, 100, 1000, 1000, 100, 10, 1}
	wm2 := NewWaveletMatrix(b)

	//Access
	if r := wm.Access(0); r != int(4) {
		t.Error("Expected", 4, "Got", r)
	}
	if r := wm.Access(6); r != int(5) {
		t.Error("Expected", 5, "Got", r)
	}
	if r := wm.Access(11); r != int(0) {
		t.Error("Expected", 0, "Got", r)
	}
	if r := wm.Access(9); r != int(3) {
		t.Error("Expected", 3, "Got", r)
	}
	if r := wm2.Access(0); r != int(1) {
		t.Error("Expected", 1, "Got", r)
	}
	if r := wm2.Access(3); r != int(1000) {
		t.Error("Expected", 1000, "Got", r)
	}
	if r := wm2.Access(7); r != int(1) {
		t.Error("Expected", 1, "Got", r)
	}
	//Rank
	if r := wm.Rank(5, 3); r != int(3) {
		t.Error("Expected", 3, "Got", r)
	}
	if r := wm.Rank(5, 1); r != int(1) {
		t.Error("Expected", 1, "Got", r)
	}
	if r := wm.Rank(1, 8); r != int(2) {
		t.Error("Expected", 2, "Got", r)
	}

	if r := wm2.Rank(1, 7); r != int(2) {
		t.Error("Expected", 2, "Got", r)
	}
	if r := wm2.Rank(100, 4); r != int(1) {
		t.Error("Expected", 1, "Got", r)
	}
	//Select
	if r, _ := wm.Select(5, 4); r != int(6) {
		t.Error("Expected", 6, "Got", r)
	}
	if r, _ := wm.Select(0, 1); r != int(11) {
		t.Error("Expected", 11, "Got", r)
	}
	if r, _ := wm.Select(1, 2); r != int(8) {
		t.Error("Expected", 8, "Got", r)
	}
	if r, _ := wm2.Select(1, 2); r != int(7) {
		t.Error("Expected", 7, "Got", r)
	}
	if r, _ := wm2.Select(1, 1); r != int(0) {
		t.Error("Expected", 0, "Got", r)
	}
	//Quantile
	if r := wm.Quantile(1, 10, 8); r != int(5) {
		t.Error("Expected", 5, "Got", r)
	}
	if r := wm.Quantile(0, 1, 1); r != int(4) {
		t.Error("Expected", 4, "Got", r)
	}
	if r := wm2.Quantile(0, 7, 1); r != int(1) {
		t.Error("Expected", 1, "Got", r)
	}
	if r := wm2.Quantile(1, 3, 1); r != int(10) {
		t.Error("Expected", 10, "Got", r)
	}
}
