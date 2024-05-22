package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func Generrate_date() {
	year := 2020 + rand.Intn(6)
	month := rand.Intn(12) + 1
	dayInMonth := rand.Intn(31) + 1
	var leep = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leep {
		switch month {
		case 2:
			dayInMonth = 29
		case 4, 6, 9, 11:
			dayInMonth = 30
		}
		day := rand.Intn(dayInMonth) + 1
		fmt.Print("Leep Year \t")
		fmt.Println(era, year, month, day)
	} else {
		switch month {
		case 2:
			dayInMonth = 28
		case 4, 6, 9, 11:
			dayInMonth = 30
		}
		day := rand.Intn(dayInMonth) + 1
		fmt.Print("Not Leep \t")
		fmt.Println(era, year, month, day)
	}
}

func main() {
	count := 10
	for i := 0; i < count; i++ {
		Generrate_date()
	}
}
