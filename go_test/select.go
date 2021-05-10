package main

import (
	"fmt"
	"time"
)

/*

select case  default

select ：从多个发送/接收channel操作中选择。
select 语句将阻塞，直到其中一个发送/接收操作准备就绪
如果有多个操作同时准备好，则随机选择一个操作。


default case :在其他操作没准备好的情况下，开始执行。防止select语句阻塞

*/


func server1(ch chan string){
	time.Sleep(20*time.Second)
	ch <- "from server1"
}

func server2(ch chan string){
	time.Sleep(10*time.Second)
	ch <- "from server2"
}


func main()  {
	startTime := time.Now()
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for {
		time.Sleep(5*time.Second)
		select {
		case s1 := <- output1:
			fmt.Println(s1)
			return
		case s2 := <- output2:
			fmt.Println(s2)
			return
		default:
			fmt.Println("no value from channel...")

		}
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		fmt.Println(diff.Seconds())
	}


}

