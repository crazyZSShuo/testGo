package main

import "fmt"

//使用channel 管道进行通信
// 使用make来进行定义  初始channel值为nil


//func Hello(done chan bool){
//	fmt.Println("hello done...")
//	done <- true
//}
//
//func main()  {
//	var a chan bool
//	if a == nil{
//		a := make(chan bool)
//		go Hello(a)
//		data := <- a
//		fmt.Println("data: ", data)
//	}
//}


func calcSquares(number int, squareop chan int){
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit *digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int){
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main()  {
	number := 587
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("finnal result: ", squares + cubes)
}