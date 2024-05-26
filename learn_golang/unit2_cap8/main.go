package main

import (
	"fmt"
	"math/big"
)

// big包大数
func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000",10)
	fmt.Println("Alpha Centauri is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")
	fmt.Printf("type of lightSpeed is %T, and it's %[1]v\n", lightSpeed)
	fmt.Printf("type of secondsPerDay is %T, and it's %[1]v\n", secondsPerDay)
}

