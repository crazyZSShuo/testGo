package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

工作池 worker pool
缓冲channel的一个重要用途就是--worker pool
worker pool ： 是一组线程，等待分配给他们的任务。一旦完成分配的任务，他们就会为下一个任务提供服务。

worker pool 核心功能

1.创建一个goroutines池，该池监听输入缓冲channel
2.向输入缓冲channel添加任务
3.任务完成后， 将结果写入输出缓冲channel
4.从输出缓冲channel读取和输出结果

*/

type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)


func digits(randomno int) int {
	sum := 0
	no := randomno
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	//time.Sleep(2*time.Second)
	return sum
}

func allocate(noOfJobs int){
	for i := 0; i < noOfJobs; i++{
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool){
	for result := range results{
		fmt.Printf("JOb id %d, input random no %d, sum of digits %d\n",
			result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func worker(wg *sync.WaitGroup){
	for job := range jobs{
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()

}

func createWorkerPool(noOfWorkers int){
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}


func main(){
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 50
	createWorkerPool(noOfWorkers)
	<- done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken: ", diff.Seconds(), " seconds")
}