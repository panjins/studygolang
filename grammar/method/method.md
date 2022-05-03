## 方法
### 1.方法声明

一个方法就是一个包含了接受者的函数，接受者的参数会将该函数附加到这种类型上，相当于为这种类型定义了一种 独占的方法。

```go
package main

import "fmt"

type Add struct {
	x, y int
}
//sum() 为 Add类型的方法
//a 为接收器 接收器通常使用类型的首字母命名
func (a Add) sum() int {
	return a.x + a.y
}

func main() {
	s1 := Add{x: 2, y: 3}
    // 调用sum方法
	//s1.sum 叫选择器
    fmt.Println(s1.sum()) 

}

```

### 2.非指针接收器

我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者`interface`

```go
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
		{1,1},
		{5,1},
		{5,4}
		{1,1},
	}
	fmt.Println(perim.Distance())


```



### 3.指针接收器

当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者 函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需 要用到指针了.

```go
//方法 指针接收器
// 想更新变量的值或者避免拷贝时 使用指针接收器

type Ponit struct {
	x, y float64
}

func (p *Ponit) AllSUM(factor float64) float64 {
	p.x *= factor
	p.y *= factor
	return p.x + p.y
}

func main() {
	r := &Ponit{1, 2}

	fmt.Println(r.AllSUM(2))

	//第二种调用方式
	p := Ponit{1, 2}
	pp := &p
	pp.AllSUM(2)

	//第三种调用方式
	ppp := Ponit{1, 2}
	ppp.AllSUM(2)

}

```

