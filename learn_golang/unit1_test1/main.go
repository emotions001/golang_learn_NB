package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Game Begin!")
	var rand_num = rand.Intn(100) + 1
	for {
		fmt.Println("currect num is:", rand_num)
		guss_num := rand.Intn(100) + 1
		fmt.Println("Guess num is:", guss_num)
		if rand_num == guss_num {
			fmt.Println("Win")
			break
		}
	}

	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("currect num is:", rand_num)
		if rand_num > 50 {
			fmt.Println("over 50!")
			count -= 1
		} else if rand_num == 50 {
			fmt.Println("is 50!")
			count -= 1
			fmt.Println("Win!")
		} else {
			fmt.Println("under 50!")
			count -= 1
		}

	}
}
