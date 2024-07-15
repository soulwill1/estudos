package main

import "fmt"

func main() {
	resultado := soma(1, 1)
	fmt.Println(resultado)
}

func soma(a int, b int) int{
	return a + b
}