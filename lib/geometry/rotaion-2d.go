package main

import (
	"fmt"
	"math"
	"sort"
)

type Point2D struct {
	x, y int
}
type Object2D struct {
	points     []Point2D
	np         []Point2D
	needUpdate bool
}

func NewObject2D() *Object2D {
	return &Object2D{}
}
func (o2 *Object2D) Add(p Point2D) {
	o2.needUpdate = true
	o2.points = append(o2.points, p)
}
func (o2 *Object2D) normalize() {
	mx, my := math.MaxInt64, math.MaxInt64
	for i := 0; i < len(o2.points); i++ {
		mx = min(mx, o2.points[i].x)
		my = min(my, o2.points[i].y)
	}
	normalized := make([]Point2D, len(o2.points))
	for i, coord := range o2.points {
		normalized[i] = Point2D{coord.x - mx, coord.y - my}
	}
	sort.Slice(normalized, func(i, j int) bool {
		if normalized[i].y == normalized[j].y {
			return normalized[i].x < normalized[j].x
		} else {
			return normalized[i].y < normalized[j].y
		}
	})
	o2.np = normalized
	o2.needUpdate = false
}
func (o2 *Object2D) Rotate90() {
	p := o2.points
	conv := func(x, y int) Point2D {
		nx, ny := -y, x
		return Point2D{nx, ny}
	}
	for i := 0; i < len(p); i++ {
		t := conv(p[i].x, p[i].y)
		p[i].x = t.x
		p[i].y = t.y
	}
	o2.normalize()
	o2.needUpdate = false
}
func (o2 *Object2D) IsSame(o Object2D) bool {
	if len(o2.points) != len(o.points) {
		return false
	}
	if o2.needUpdate == true {
		o2.normalize()
	}
	if o.needUpdate == true {
		o.normalize()
	}
	for i := 0; i < len(o2.np); i++ {
		if o2.np[i].x == o.np[i].x && o2.np[i].y == o.np[i].y {
			continue
		}
		return false
	}
	return true
}
func (o2 *Object2D) Show() {
	if o2.needUpdate == true {
		o2.normalize()
		o2.needUpdate = false
	}
	fmt.Println(o2.points, o2.np)
}
