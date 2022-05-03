package main

import (
	"fmt"
	"math"
)

//方法非指针接受器

type Add struct {
	x, y int
}

//a 为接收器 接收器通常使用类型的首字母命名
func (a Add) sum() int {
	return a.x + a.y
}

func (a Add) doubleSum(s Add) int {
	return a.x + a.y + s.x + s.y
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum = sum + path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	s1 := Add{x: 2, y: 3}
	//s1.sum
	fmt.Println(s1.sum())
	s2 := Add{2, 6}
	fmt.Println(s2.doubleSum(s2))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

}
