package main

import "fmt"

func fullname(firstname *string , lastname *string){
	defer fmt.Println("defer called in fullname")
	if firstname == nil{
		panic("Error: firstname cannot be nil")
	}
	if lastname == nil{
		panic("Error: lastname cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstname, *lastname)
	fmt.Println("returned normally from fullname")
}

func main()  {
	defer fmt.Println("defer called in main")
	firstname := "zs"
	fullname(&firstname, nil)
	fmt.Println("returned normally from main")

}



