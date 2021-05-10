package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// 单行写入
/*
func main()  {
	f, err := os.Create("./world.txt")
	if err != nil{
		fmt.Println("Create files is failed")
		return
	}
	l, err := f.WriteString("hello world, Go!Go!Go!Go!Go!Go!")// 单行写入, 原文件存在内容将替换
	if err != nil{
		fmt.Println("Write contents to files is failed", err)
		f.Close()
		return
	}
	fmt.Println(l,":write success!")
	err = f.Close()
	if err != nil{
		fmt.Println("Close file is failed", err)
		return
	}

}
*/


//多行写入 追加
// fmt.Fprintln
/*
func main()  {
	f, err := os.Create("./world.txt")
	if err != nil{
		fmt.Println("Create files is failed.", err)
		return
	}
	d := []string{"hello world.", "This is the first pages.","good good study."}
	for _,v := range d{
		fmt.Fprintln(f,v)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(v,":write success!")
	}
	err = f.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("write all done!")

}
*/

//对文件进行追加
/*
func main()  {
	f, err := os.OpenFile("./world.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println(err)
		return
	}
	newLine := "zhui jia de hang shu."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil{
		fmt.Println(err)
		return
	}
	err = f.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("write success.")
}
*/



// 利用goroutines 并发写入文件
/*
1. 创建一个信道，用于读取和写入生成的随机数。
2. 创建 100 个生产者goroutines。每个 goroutine 将生成一个随机数，并将该随机数写入一个信道。
3. 创建一个消费者 goroutine，它将从信道读取并将生成的随机数写入文件。因此，我们只有一个 goroutine 并发地写入文件，从而避免了竞争条件.
4. 一旦完成，关闭文件。
*/


func produce(data chan int, wg *sync.WaitGroup){
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool){
	f, err := os.Create("./write_nums.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	for d := range data{
		_,err = fmt.Fprintln(f, d)
		if err != nil{
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil{
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main(){
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func(){
		wg.Wait()
		close(data)
	}()
	d := <- done
	if d == true{
		fmt.Println("write success.")
	}else {
		fmt.Println("write failed.")
	}
}
