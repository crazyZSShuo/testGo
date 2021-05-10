package main

import (
	"errors"
	"fmt"
	"math"
)

// 利用New 自定义错误信息

func circleArea(radius float64) (float64, error){
	if radius < 0{
		return 0, errors.New("Area calculation failed, radius is less than zero")
		//return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius) //利用Errorf  格式化错误信息
	}
	return math.Pi * radius * radius, nil
}

func main()  {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil{
		fmt.Println("error:",err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
