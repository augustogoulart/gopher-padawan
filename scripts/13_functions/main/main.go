package main

import (
	"github.com/agoulartx/gopher-padawan/scripts/13_functions"
	"fmt"
)

func main(){
	// Arguments
	functions.Singleprint("Hello!")
	functions.Multiprint("Welcome,","Jhon")
	fmt.Println("-------------------")

	// Returns
	a := functions.Greet("Augusto","Goulart")
	fmt.Println(a)
	fmt.Println(functions.Namedgreet("Jhon","Doe"))
	fmt.Println(functions.Doublereturn("Carlos","Mario"))
	fmt.Println("-------------------")

	// Variadic
	v := functions.Variadicsum(188.4,487.6,8.9,222.1,543.6)
	fmt.Println(v)
	fmt.Println("-------------------")

}