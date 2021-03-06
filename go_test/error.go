package main

import (
	"fmt"
	"net"
)

func main(){
	addr, err := net.LookupHost("golangboot123.com")
	fmt.Println(addr, err)
	if err, ok := err.(*net.DNSError);ok{
		fmt.Println(ok)
		if err.Timeout(){
			fmt.Println("operation timed out")
		}else if err.Temporary() {
			fmt.Println("temporary error")
		}else{
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)

}
