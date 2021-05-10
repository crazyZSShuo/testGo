package main

import "fmt"

type Worker interface {
	Work()
}

type Person3 struct {
	name string
	age int
}

func (p Person3) Work() {
	fmt.Println(p.name)
}

func printName(w Worker){
	fmt.Println(w)
}

func main() {
	p := Person3{
		name: "zs",
		age:  10,
	}
	var w Worker
	w = p
	printName(w)
	p.Work()
}
