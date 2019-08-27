package main

import "fmt"

type geometry interface {
	shape() string
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (rect) shape() string{
	return "长方形"
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64{
	return 2 * (r.width + r.height)
}

func (circle) shape() string{
	return "圆形"
}

func (c circle) area() float64{
	return 2 * 3.1415926 * c.radius * c.radius
}

func (c circle) perim() float64{
	return 2 * 3.1415926 * c.radius
}

func measure(g geometry)  {
	fmt.Println(g.shape())
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3, height: 4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}