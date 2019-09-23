package main

import (
	"fmt"
)

func main() {

	//INTEGERS
	var a int = -2
	var b int = 60
	c := 3

	fmt.Println("Multiplication of a and b is: ", (a * b))
	fmt.Println("Random calculation: ", (c + 4 - a*b))

	//FLOAT
	d := 5 + 3.2
	fmt.Println("Value of d is: ", d)

	//COMPLEX
	e := 20 + 27i
	fmt.Println(e)

	//STRING
	f := "Smelly cat, smelly cat"
	g := ", what are they feeding you?"
	fmt.Println(f + g)

	//BOOLEAN
	var h bool
	i := true
	fmt.Println(h)
	fmt.Println("OR Operator: ", h || i)
	fmt.Println("AND Operator: ", h && i)

	//TYPE CONVERSION
	j := 7
	//var k float32 = j
	var k float32 = float32(j)
	fmt.Println("Value of k is ", k)

}