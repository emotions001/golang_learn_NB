// 关于变量和作用域
// 声明变量的方式：
// var count = 10	(通用变量声明方式)
// count := 10		(简短变量声明方式)
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // era变量在整个包都是可用的

func function1() { // era变量和year变量都处于作用域之内
	year := 2024
	switch month := rand.Intn(12) + 1; month { // 变量era、year和month都处于作用域之内
	case 2:
		day := rand.Intn(29) + 1 // 变量era、year、month和day都处于作用域之内
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // 这两个day变量是全新声明的变量，跟上面声明的同名变量并不相同
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // 这两个day变量是全新声明的变量，跟上面声明的同名变量并不相同
		fmt.Println(era, year, month, day)
	} // month变量和day变量不再处于作用域之内
}

// 以下是针对function1的优化重构版本，收窄了作用域，减少了计算机负担
func function2() {
	year := 2024
	month := rand.Intn(12) + 1
	dayInMonth := rand.Intn(31) + 1
	switch month {
	case 2:
		dayInMonth = 29
	case 4, 6, 9, 11:
		dayInMonth = 30
	}
	day := rand.Intn(dayInMonth) + 1
	fmt.Println(era, year, month, day)
}

func main() {
	function1()
	function2()
}

// 因为对era变量的声明位于main函数之外的包作用域中，所以它对于main包中的所有函数都是可见的。
// 注意因为包作用域在声明变量时不允许使用简短声明，所以我们无法在这个作用域中使用
