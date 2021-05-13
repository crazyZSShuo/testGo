package main

import (
	"fmt"
	"sync"
)

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



// defer 延迟调用 入栈操作 先进后出 Lifo 函数返回之前调用defer 打印
//func main() {
//	var s = "hello"
//	fmt.Printf("开始打印...\n")
//	defer fmt.Printf("defer test1\n") // 入栈底部 （最后打印，出栈）
//	defer fmt.Printf("defer test2\n") // 入栈中部 （其次打印，出栈）
//	defer fmt.Printf("defer test3\n") // 入栈顶部 （优先打印，出栈）
//	fmt.Println(s)
//	fmt.Printf("结束打印...\n")
//}


type rect struct {
	length int
	width int
}

func (r rect) area(wg *sync.WaitGroup){
	defer wg.Done()

	if r.length < 0{
		fmt.Printf("rect %v's length should be greater than zero \n", r)
		return
	}
	if r.width < 0{
		fmt.Printf("rect %v's width should be greater than zero \n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main(){
	var wg sync.WaitGroup
	r1 := rect{
		length: -67,
		width:  10,
	}
	r2 := rect{
		length: 67,
		width:  -10,
	}
	r3 := rect{
		length: 67,
		width:  10,
	}
	rects := []rect{r1, r2, r3}
	for _,v := range rects{
		wg.Add(1)
		go v.area(&wg)
	}
	defer fmt.Println("All go goroutines is done!")
	defer wg.Wait()
}









