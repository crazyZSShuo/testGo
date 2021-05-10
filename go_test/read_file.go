package main

import (
	"bufio"
	"fmt"
	"os"
)


//全部读取文件
//func main(){
//	// 读取正常， err = nil, 返回字节切片 eg：[72 101 108 108 ...]
//	data, err := ioutil.ReadFile("./hello.txt")
//	if err != nil{
//		fmt.Println("File reading error", err)
//		return
//	}
//	fmt.Println("Contents of files:", string(data))
//}

// 逐行读取文件
func main()  {
	data, err := os.Open("./hello.txt")
	if err != nil{
		fmt.Println("Read files is failed")
		return
	}
	f := bufio.NewScanner(data)
	for f.Scan(){
		fmt.Println("contents from files:", f.Text())
	}
}
