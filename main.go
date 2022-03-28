package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	city2 "rebrainpkg/city"
	newcolor "rebrainpkg/color"
	. "rebrainpkg/wordz"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")

	fmt.Println(Hello)    //Вызов переменной из пакета wordz
	fmt.Println(Random()) //Вызов функции из пакета wordz

	fmt.Println(xstrings.Shuffle(city2.City()))
	fmt.Println(city2.Digit())
}
