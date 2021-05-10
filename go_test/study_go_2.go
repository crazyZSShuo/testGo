package main

import "fmt"

type Address struct {
	city string
	state string
}

type Person struct {
	firstname string
	lastname string
	age int
	address Address
}


//利用提升字段
type Person2 struct {
	firstname string
	lastname string
	age int
	Address
}

func main()  {
	person1 := Person{
		firstname: "zs",
		lastname:  "shuo",
		age:       100,
		address:   Address{
			city: "henan",
			state: "zhoukou",
		},
	}
	fmt.Println(person1.address)
	fmt.Println(person1.address.city) //经过address 调用city
	fmt.Println("======================")
	// 提升字段
	p2 := Person2{
		firstname: "jack",
		lastname:  "mark",
		age:       100,
		Address:   Address{
			city: "beijing",
			state: "beijing",
		},
	}
	fmt.Println(p2.firstname)
	fmt.Println(p2.city)  //直接通过实例调用
}
