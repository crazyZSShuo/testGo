package main

import "fmt"

//func main()  {
//	a := "hello"
//	fmt.Println("开始打印...")
//	defer fmt.Println("延迟调用defer: defer print...")
//	fmt.Printf("Original String: %s\n", string(a))
//	fmt.Printf("Reversed String: ")
//	for _, v := range []rune(a){
//		defer fmt.Printf("%c", v)  //LIFO
//	}
//	fmt.Println("结果打印...")
//}



// defer 延迟调用 入栈操作 先进后出 Lifo 快要return时 打印
func main() {
	var s = "hello"
	fmt.Printf("开始打印...\n")
	defer fmt.Printf("defer test1\n") // 入栈底部 （最后打印，出栈）
	defer fmt.Printf("defer test2\n") // 入栈中部 （其次打印，出栈）
	defer fmt.Printf("defer test3\n") // 入栈顶部 （优先打印，出栈）
	fmt.Println(s)
	fmt.Printf("结束打印...\n")
}