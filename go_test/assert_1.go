package main

import (
	"fmt"
)


//i.(T)  T:int string float ...
func assert(i interface{}){
	v,ok := i.(int)
	fmt.Println(v, ok)
}



//switch case
func findType(i interface{}){
	switch i.(type) {
	case string:
		fmt.Printf("%s :is string", i)
	case int:
		fmt.Printf("%d :is int", i)
	default:
		fmt.Printf("unknow type!")
	}
}

//type VS interface

type Work interface {
	Work()
}

type Shuo struct {
	name string
	age int
}

func (s Shuo) Work() {
	fmt.Println("s: ", s)
}

func getType(s interface{}){
	switch v := s.(type) {
	case Work:
		v.Work()
	default:
		fmt.Printf("unknow type!")
	}
}

func main()  {
	var i interface{} = 56
	assert(i)
	var s interface{} = "hello world"
	assert(s)
	fmt.Printf("=============\n")
	var f interface{} = 10.2
	findType(i)
	findType(s)
	findType(f)
	// output:
	//56 true
	//0 false
	fmt.Printf("===========\n")
	sh := Shuo{
		name: "zs",
		age:  10,
	}
	getType(f)
	getType(sh)
	//output:
	//unknow type!
	//s:  {zs 10}
}