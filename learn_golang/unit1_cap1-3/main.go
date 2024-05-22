package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// cap1-2
	var num = rand.Intn(10) + 1 // rand伪随机函数输入10产生0-9范围随机数，所以+1
	fmt.Println(num)
	num = rand.Intn(10)
	fmt.Println(num)
	// cap3
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "Walk outside"
	var exit = strings.Contains(command, "outside") // strings包的Contains函数检测字符是否存在于字符串中
	fmt.Println("You leave the cave:", exit)

	// if 分支
	var command2 = "go east"
	if command2 == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command2 == "go inside" {
		fmt.Println("this is not going to excute")
	} else {
		fmt.Println("Didn't quite get thet.")
	}

	// 逻辑运算符号“ || 或 、 && 与  ”

	var year = 2100
	var leep = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leep {
		fmt.Println("2100 is leep year")
	} else {
		fmt.Println("2100 isn't leep year")
	}

	// 跟大多数编程语言一样，Go也采用了短路逻辑：如果位于||运算符之前的第
	// 一个条件为真，那么位于11运算符之后的条件就可以被忽略，没有必要再对其进行求值。
	// 具体到代码清单3-4中的例子，当给定年份可以被400整除时，程序就不必再进行后续的判断了。
	// &&运算符的行为与||运算符正好相反：只有在两个条件都为真的情况下，运算结果才为真。
	// 对于代码清单3-4中的例子，如果给定年份无法被4整除，那么程序就不会对后续条件进行求值。

	// switch 分支判断
	var command3 = "go inside"
	switch command3 { // 将命令与多个给定的分支进行比较
	case "go east":
		fmt.Println("You have further up the mountain.")
	case "enter cave", "go inside": // 使用逗号分隔多个可选项
		fmt.Println("You find youself in a cave.")
	case "read sign":
		fmt.Println("this not gone be excute.")
	default:
		fmt.Println("Dontn't quit get that.")
	}

	// GoLang的循环
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second) // 引入Time包的延时函数
		count--
	}
	fmt.Println("Liftoff!")

	var degress = 0
	for {
		fmt.Println(degress)
		degress++
		if degress >= 360 {
			degress = 0
			if rand.Intn(2) == 0 { // 在随机数不等于0之前，degress会一直从0-360不停重复循环（旋转）
				break
			}
		}
	}

}
