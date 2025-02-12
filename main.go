package main

import "fmt"

var x int = 40

const y int = 42

func main() {
	// printing value for hexadecimal and binary
	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	// fmt.Printf("%v  \t  %b \t %#X \n", a, a, a)
	// fmt.Printf("%v  \t  %b \t %#X \n", b, b, b)
	// fmt.Printf("%v  \t  %b \t %#X \n", c, c, c)
	// fmt.Printf("%v  \t  %b \t %#X \n", d, d, d)
	// fmt.Printf("%v  \t  %b \t %#X \n", f, f, f)
	// fmt.Printf("%v  \t  %b \t %#X \n", e, e, e)
	// fmt.Printf("%v  \t  %b \t %#X \n", d, d, d)
	fmt.Printf("The value of x %v and the type of x %T \n", x, x)
	fmt.Printf("The value of y %v and the type of y %T \n", y, y)
	z := 42
	fmt.Println(z)
	fmt.Println("Hello, World!")
}
