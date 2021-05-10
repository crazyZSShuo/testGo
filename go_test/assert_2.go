package main

import "fmt"

type Describer interface {
	Describer()
}

type Person4 struct {
	name string
	age int
}

func (p Person4) Describer(){
	fmt.Println(p.name, p.age)
}

type Address struct {
	country string
	state string
}

func (a *Address) Describer(){
	fmt.Println(a.country, a.state)
}

func main()  {
	var d Describer
	p1 := Person4{
		name: "zhangsan",
		age:  10,
	}

	d = p1
	d.Describer()

	a1 := Address{
		country: "China",
		state:   "beijing",
	}
	d = &a1
	d.Describer()


}