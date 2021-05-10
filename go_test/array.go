package main

import (
	"fmt"
)

//当一个切片作为参数传给函数时 函数内部所做的更改在函数外部
// 也可以看得到

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	fmt.Println(nums)
	nums[0] = 100
	fmt.Println(nums)
	nums = append(nums, 101)
	fmt.Println(nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
//func main() {
//	nums := []int{89, 90, 95}
//	find(89, nums...)
//	fmt.Println(nums)
//}

//make(map[type of key]type of value)
// 通过make来创建map
//func main()  {
//	personSalary := make(map[string]int)
//	personSalary["zs"] = 100
//	fmt.Println(personSalary)
//	fmt.Println(personSalary["zs"])
//	for i , v:= range personSalary{
//		fmt.Println(i, v)
//	}
//}

// 声明时初始化
//func main()  {
//	emplyeeSalary := map[string]int{
//		"zs":100,
//		"shuo":200,
//	}
//	emplyeeSalary["dd"] = 90
//	fmt.Println(emplyeeSalary, len(emplyeeSalary))
//	str1 := "Señor"
//	str2 := "hello"
//	nu := len(str1)
//	fmt.Println(nu, utf8.RuneCountInString(str1))
//	str1_and_str2 :=fmt.Sprintf("%s %s", str1, str2) // could use "+"  str1 + str2
//	fmt.Println("str1_and_str2: ",str1_and_str2)
//	for i := range emplyeeSalary{
//		if i == "ds" {
//			fmt.Println(i , emplyeeSalary[i])
//			break
//		}else{
//			fmt.Println(i , emplyeeSalary[i])
//		}
//	}
//}

// struct
type Employee struct{
	firstname string
	lastname string
	age int
	salary int
}

func main()  {
	var a string
	var b string
	var c int
	var d int
	a = "zs"
	b = "shuo"
	c = 10
	d = 100
	emp1 := Employee{
		firstname: a,
		lastname:  b,
		age:       c,
		salary:    d,
	}
	var kk int
	var emp2 Employee
	emp3 := &Employee{
		firstname: "jhon",
		lastname:  "smith",
		age:       20,
		salary:    300,
	}
	fmt.Println(emp1)
	emp1.salary = 200
	fmt.Println(emp1)
	fmt.Println("=============")
	fmt.Println(emp2)
	fmt.Println(kk)
	fmt.Println("===========")
	fmt.Println("address of emp3: ", emp3)
	fmt.Println(*emp3)
	fmt.Println((*emp3).firstname)
	fmt.Println(emp3.firstname)
}