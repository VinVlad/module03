package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	"module03/Rand"
	"module03/wordz"
)

func main() {
	city := []string{"Владивосток", "Москва", "Санкт-Петербург", "Южно-Сахалинск"}
	wordz.Words = city
	wordz.Prefix = "Random city: "
	randCity := xstrings.Shuffle(Rand.City())
	fmt.Println(randCity)

	digit := []string{"Один", "Писятдва", "228", "1488"}
	wordz.Words = digit
	wordz.Prefix = "Random digit: "
	randdigit := Rand.Digit()
	fmt.Println(randdigit)
}
