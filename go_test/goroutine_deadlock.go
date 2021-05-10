package main

import (
"fmt"
)


// 死锁  deadlock

//func main() {
//	ch := make(chan string, 2)
//	ch <- "naveen"
//	ch <- "paul"
//	ch <- "steve"
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}


func whileTrue(ch chan string){
	for {
		data := <- ch
		fmt.Println(data)
	}
}


func main() {
	ch := make(chan string, 2)
	data := [3]string{"naveen", "paul", "steve"}
	for _, v := range data{
		ch <- v
		go whileTrue(ch)
	}

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
