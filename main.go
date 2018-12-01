package main

import (
	"fmt"
	utils "github.com/sameetn/mystringutils"
)

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Println(utils.Upper(s))
	fmt.Println(utils.Lower(s))
}
