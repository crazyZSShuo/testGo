package main

import "fmt"

// 函数是一等公民
// 函数可以赋值给变量，也可以当作参数传递给另外一个函数（高阶函数）
// 闭包  = 匿名函数  函数内部可以访问函数之外定义的变量
func appendStr() func(s string) string {
	t := "hello"
	c := func(s string) string {
		t = t + " " + s
		return t
	}
	return c
}

//func main(){
//	a1 := appendStr()
//	a2 := appendStr()
//	fmt.Println(a1("world"))
//	fmt.Println(a2("zs"))
//	fmt.Println(a1("zhangsan"))
//	fmt.Println(a2("lisi"))
//}


//例子1  不适用匿名函数
type Student struct {
	firstname string
	lastname string
	grade string
	country string
}

func has_filter(students []Student) []Student {
	var s []Student
	for _, v := range students{
		if v.grade == "A" {
			s = append(s, v)
		}
	}
	return s
}

//func main(){
//	s1 := Student{
//		firstname: "zs",
//		lastname:  "shuo",
//		grade:     "A",
//		country:   "China",
//	}
//	s2 := Student{
//		firstname: "li",
//		lastname:  "si",
//		grade:     "B",
//		country:   "China",
//	}
//	students := []Student{s1, s2}
//	f := has_filter(students)
//	fmt.Println(f)
//}


//例子2  使用匿名函数
func filters(s []Student, f func(s Student) bool ) []Student {
	var r []Student
	for _, v := range s{
		if f(v) == true{
			r = append(r, v)
		}
	}
	return r
}

func main()  {
	s1 := Student{
		firstname: "zs",
		lastname:  "shuo",
		grade:     "A",
		country:   "China",
	}
	s2 := Student{
		firstname: "li",
		lastname:  "si",
		grade:     "B",
		country:   "China",
	}
	s := []Student{s1, s2}
	f := filters(s, func(s Student) bool {
		if s.grade == "B"{
			return true
		}
		return false
	})
	fmt.Println(f)
}