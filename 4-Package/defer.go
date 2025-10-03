package main

import "fmt"

func main() {
	fmt.Println("Primeira linha")
	//A ordem de execução do defer é um lifo (last in first out), ou seja, ultimo a entrar e primeiro a sair.
	defer fmt.Println("Segunda linha")
	defer fmt.Println("Quarta linha")
	defer fmt.Println("Quinta linha")
	fmt.Println("Terceira linha")
}
