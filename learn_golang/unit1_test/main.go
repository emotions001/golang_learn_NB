package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Game Begin")
	var rand_num = rand.Intn(100) + 1
	var gusee_num int
	fmt.Println("Input a num:")
	fmt.Scan(&gusee_num)
	fmt.Println(rand_num)

}
