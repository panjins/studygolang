package main

import "fmt"

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
	//第一种调用 方式
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
