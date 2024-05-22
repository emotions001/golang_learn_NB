package main

import (
	"fmt"
	"math"
)

func main() {
	// 声明浮点类型

	// 以下几种浮点数的定义方式都是相同的,在不特定指定的情况下，GoLang将默认将小数定为float64类型
	pi1 := math.Pi
	var pi2 = math.Pi
	var pi3 float64 = math.Pi
	fmt.Println(pi1, pi2, pi3)
	var pi4 float32 = math.Pi // 这是指定32位浮点数（单精度浮点数），不常用除非特定需要
	fmt.Println(pi4)

	// 零值
	
	// 当声明一个变量但未赋值，默认就是零值
	var price float64
	fmt.Println(price)


	// 打印浮点类型数值，使用 Printf
	third := 10.0/3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

}
