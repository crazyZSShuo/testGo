package main

import (
	"fmt"
	"strconv"
)


func main()  {
	var i = 12
	//数字转字符串
	I := strconv.Itoa(i)
	fmt.Printf("%T\n", I)
	fmt.Println(I)

	//字符串转整型
	var s = "12"
	S, err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%T\n", S)
		fmt.Println(S)
	}
}
