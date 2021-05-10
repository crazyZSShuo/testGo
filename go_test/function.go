package main

import "fmt"

//
////方法：只是一个带有特殊接收器类型的函数
//
////语法: func (t type) methodName () {
////}
//
//type Employee struct {
//	firstname string
//	salary int
//	currency string
//}
//
////在 Employee 类型上创建一个方法 dispalySalary
////方式1：
////func (e Employee) dispalySalary(){
////	fmt.Println(e.firstname, e.salary, e.currency)
////}
//
////方式二:
//func dispalySalary(e Employee){
//	fmt.Println(e.firstname, e.salary, e.currency)
//}
//
//func main()  {
//	emp1 := Employee{
//		firstname: "zs",
//		salary:    100,
//		currency:  "dd",
//	}
//	//emp1.dispalySalary() // 方式1 调用
//	dispalySalary(emp1) // 方式二调用
//}
//

type Employee2 struct {
	name string
	age int
}

//值接收器
func (e Employee2) changeName(newName string){
	e.name = newName
	fmt.Println(e)
}

//指针接收器
func (e *Employee2) changeAge(newAge int){
	e.age = newAge
}

func main(){
	emp := Employee2{
		name: "zs",
		age: 10,
	}
	fmt.Println("before of name: ", emp.name)
	emp.changeName("ds")
	fmt.Println("after of name: ", emp.name)

	fmt.Println("before of age: ", emp.age)
	//(&emp).changeAge(20)
	emp.changeAge(20)
	fmt.Println("after of age: ", emp.age)

}