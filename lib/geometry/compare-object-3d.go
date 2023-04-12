package main

import (
	"fmt"
	"math"
	"sort"
)

type Point3D struct {
	x, y, z int
}
type Object3D struct {
	points []Point3D
}

func NewObject3D() *Object3D {
	return &Object3D{}
}
func (o3 *Object3D) Add(p Point3D) {
	o3.points = append(o3.points, p)
}
func (o3 *Object3D) IsSame(p Object3D) bool {
	b1 := o3.points
	b2 := p.points
	normalize := func(b []Point3D) []Point3D {
		mx, my, mz := math.MaxInt64, math.MaxInt64, math.MaxInt64
		for i := 0; i < len(b); i++ {
			mx = min(mx, b[i].x)
			my = min(my, b[i].y)
			mz = min(mz, b[i].z)
		}
		normalized := make([]Point3D, len(b))
		for i, coord := range b {
			normalized[i] = Point3D{coord.x - mx, coord.y - my, coord.z - mz}
		}
		return normalized
	}
	if len(b1) != len(b2) {
		return false
	}
	if len(b1) == 1 || len(b2) == 2 {
		return true
	}
	b1 = normalize(b1)
	b2 = normalize(b2)
	max := Point3D{0, 0, 0}
	for _, coord := range b2 {
		if coord.x > max.x {
			max.x = coord.x
		}
		if coord.y > max.y {
			max.y = coord.y
		}
		if coord.z > max.z {
			max.z = coord.z
		}
	}
	sort.Slice(b1, func(i, j int) bool {
		return fmt.Sprintf("%v", b1[i]) < fmt.Sprintf("%v", b1[j])
	})
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			sort.Slice(b2, func(i, j int) bool {
				return fmt.Sprintf("%v", b2[i]) < fmt.Sprintf("%v", b2[j])
			})

			if fmt.Sprintf("%v", b1) == fmt.Sprintf("%v", b2) {
				return true
			}

			for k := range b2 {
				t := b2[k].x
				b2[k].x = max.y - b2[k].y
				b2[k].y = t
			}
			max.x, max.y = max.y, max.x
		}
		if i&1 != 0 {
			for j := range b2 {
				t := b2[j].y
				b2[j].y = max.z - b2[j].z
				b2[j].z = t
			}
			max.y, max.z = max.z, max.y
		} else {
			for j := range b2 {
				t := b2[j].z
				b2[j].z = max.x - b2[j].x
				b2[j].x = t
			}

			max.x, max.z = max.z, max.x
		}
	}
	return false
}
