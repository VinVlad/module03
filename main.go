package main

import (
	"fmt"
	"github.com/fatih/color"
	. "gopackages/wordz" //Добавляем пакет wordz через точку
	newcolor "gopackages/color"
)

func main()  {
	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(Hello) //Вызов переменной из пакета wordz
	fmt.Println(Random())//Вызов функции из пакета wordz

}
