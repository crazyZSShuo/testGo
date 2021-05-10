package main

import (
	"fmt"
)

type student struct {
	firstname string
	lastname string
	grade string
	country string
}


func filter(s []student)[]student{
	var stu []student
	for _, v := range s {
		if v.country == "India"{
			v.firstname = "koko"
			stu = append(stu, v)
		}
	}
	return stu
}



func iMap(s []int) []int{
	var nums_new []int
	for _, v := range s{
		nums_new = append(nums_new, v*5)
	}
	return nums_new
}



func main()  {
	s1 := student{
		firstname: "zs",
		lastname:  "shuo",
		grade:     "B",
		country:   "China",
	}
	s2 := student{
		firstname: "lisi",
		lastname:  "da",
		grade:     "A",
		country:   "India",
	}
	s := []student{s1, s2}
	fmt.Printf("%T\n", s)
	fmt.Println(s)
	stu := filter(s)
	fmt.Println(s)
	fmt.Println(stu)
	fmt.Println("=====================")
	nums := []int{1,2,3,4}
	fmt.Println(iMap(nums))

}
