package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Variable declaration
	var a int
	fmt.Println(a) //Prints 0

	// Initialization
	a = 10
	fmt.Println(a) //Prints 10

	a = 20

	// Declare and initialize on the same line
	var b int = 20
	fmt.Println(b)

	//Type inference
	var c = 30
	fmt.Println("Type of c: ", reflect.TypeOf(c))

	//Shorthand variable declaration
	d := 40
	fmt.Println(d)

	//Multiple variable declaration
	var e, f int
	fmt.Println(e, f)

	//Multiple variable declaration with initialization
	var g, h, i int = 2, 3, 4
	fmt.Println(g, h, i)

	//Declaring and initializing multiple variables of different types
	var (
		j int = 0
		k     = "bro"
		l     = ", you a pro"
	)
	fmt.Println(j, k, l)

	//Shorthand declaration for multiple variables
	m, n := 9, 10       //Same type
	p, q := "bro", true //Different types
	fmt.Println(m, n, p, q)

}