package main

import (
	"fmt"
	"math"
	// "time"
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
	third := 10.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

	// 声明整数类型

	// 因为Go在进行类型推断的时候总是会选择int类型作为整数值的类型，所以下面这3行代码的意义是完全相同的：
	year1 := 2024
	var year2 = 2024
	var year3 int = 2024
	fmt.Println(year1, year2, year3)
	// 为了避免在Printf函数中重复使用同一个变量两次，我们可以将[1]添加到第二个格式化变量%v中，以此来复用第一个格式化变量的值year1，从而避免代码重复：
	fmt.Printf("year is %T, and the number of year is %[1]v \n", year1) // 使用格式化变量标识符 %T 可以打印出变量的类型

	// future := time.Unix(12622780800, 0)
	// fmt.Println(future)
	// t := time.Now()
	// fmt.Println(t)

	// 使用科学计数法表示一个很大的数字，并且指定数据类型为int64
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")
	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
}
