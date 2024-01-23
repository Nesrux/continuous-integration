package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := Somar(a, b)
	fmt.Printf("A Soma de A + B é %d", c)
}
func Somar(a, b int) int {
	return a + b
}

/*
func Subtracao(a, b int) int {
	return a - b
}

func Multi(a, b int) int {
	return a * b
}**/
