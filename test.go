package main

import "fmt"

var a *int

func main() {
	*a = 12
	fmt.Println(*a)
}
