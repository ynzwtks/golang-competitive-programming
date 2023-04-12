package main

// Vector
type point struct {
	x, y int
}

func readPoint() point {
	a, b := readint2()
	return point{a, b}
}
func mulPoint(a point, l int) point {
	return point{a.x * l, a.y * l}
}
func addPoint(a, b point) point {
	return point{a.x + b.x, a.y + b.y}
}
func diffPoint(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}
func cross(p1, p2 point) int {
	return p1.x*p2.y - p2.x*p1.y
}
func dot(p1, p2 point) int {
	return p1.x*p2.x + p1.y*p2.y
}
