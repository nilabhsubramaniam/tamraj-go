package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Person struct {
	first string
	last  string
	age   int
}

type SecretAgent struct {
	Person
	licenseToKill bool
}


// Interface and Polymorphism
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func main() {
	sum(1, 2, 3, 4, 5)
	rectangle := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 7}

	printShapeInfo(rectangle)
	printShapeInfo(circle)
	createFile("output.txt")
	src := rand.NewSource(time.Now().UnixNano()) // Create a new source
	rng := rand.New(src)                         // Create a new random generator
	ch := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan string)

	x := 10 // Define x
	y := 5  // Define y

	if z := 2 * rng.Intn(x); z >= x {
		fmt.Printf("The value of x %v and the type of x %T \n", x, x)
	} else {
		fmt.Printf("The value of z %v and the type of y %T \n", z, y) // Corrected: use z, not y
	}

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the workweek!")
	case "Friday":
		fmt.Println("Weekend is near!")
	default:
		fmt.Println("It's a regular day.")
	}

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Channel 1 "
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from Channel 2 "
	}()

	// select statement to understand

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case ms2 := <-ch2:
		fmt.Println(ms2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! No response within 2 seconds.")
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received, moving on...")

	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first for loop \n ", i)
	}

	// nested loop
	for i := 0; i < 10; i++ {
		fmt.Println("----")
		for j := 0; j < 10; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// create a programm that has a loop that prints every number from 0 to 99

	for i := 0; i <= 99; i++ {
		fmt.Println((i))
	}

	// switch case

	for i := 0; i < 42; i++ {
		X := rand.Intn(5)

		switch X {
		case 0:
			fmt.Println("X is 0")
		case 1:
			fmt.Println("X is 1")
		case 2:
			fmt.Println("X is 2")
		case 3:
			fmt.Println("X is 3")
		case 4:
			fmt.Println("X is 4")
		}

	}

	// use modulus and print odd numbers using continiue
	for i := 0; i <= 100; i++ {

		// printing the i value when it is odd
		if i%2 != 0 {

			fmt.Println("odd number is ", i)
		}
	}

	// For range loop

	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %v \t value  is %v \n", i, v)
	}

	// map - from range
	m := map[string]int{
		"James":  42,
		"amount": 32,
	}

	for k, v := range m {
		fmt.Printf("key  %v \t value is %v\n", k, v)
	}

	fmt.Println(".......................")
	for i := 0; i <= 100; i++ {

		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t x is %v\n", i, x)
		}
	}

	// Array literal
	a := [3]int{42, 43, 45}
	fmt.Println(a)

	b := [...]string{"hello", "gofers"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 1

	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)

	// Use a  array literal  to store  these elements  in an array and let the compiler determine  the length of array  and then also print out the array and the length of the array
	as := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(as)
	fmt.Println(len(as))
	fmt.Printf("%T\n", as)

	// Slice literals- for range loop
	xa := []string{"Almond", "apple", "orange"}
	fmt.Println(xa)
	fmt.Printf("%T\n", xa)

	for i, v := range xa {
		fmt.Printf(" %v -  %v\n", i, v)
	}

	// Slice make
	si := []string{"apple", "bananana", "carrot", "dish", "eager"}
	fmt.Println(si)

	xsi := make([]int, 0, 10)
	fmt.Println(xsi)
	fmt.Println(len(xsi))
	fmt.Println(cap(xsi))
	xsi = append(xsi, 1, 2, 3, 7, 9, 11, 12, 13, 14, 15)
	fmt.Println(xsi)

	// slice internals and underlying array
	n := []float64{3, 1, 4, 2, 7}
	fmt.Println(medianOne(n))
	fmt.Println("After medianOne", n)

	siua := []float64{3, 1, 4, 2, 7}
	fmt.Println(medianTwo(siua))

	// map introduction
	am := map[string]int{
		"Aakasha":    32,
		"Amabar":     35,
		"Neelamabar": 43,
	}

	fmt.Println("the age of Neelamabar was", am["Neelamabar"])
	fmt.Println("..................................")
	fmt.Printf("%#v\n", am)
	an := make(map[string]int)
	an["sneha"] = 32
	an["snehal"] = 39
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	fmt.Println(len(an))

	// add element - for range loop over map
	for k, v := range an {
		fmt.Printf("key %v \t value is %v\n", k, v)
	}

	// delete element from map
	delete(an, "snehal")
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	// check if key exists in map
	v, ok := an["sneha"]
	if ok {
		fmt.Printf("key %v exists with value %v\n", "sneha", v)
	} else {
		fmt.Printf("key %v does not exist\n", "sneha")
	}
	fmt.Println("..................................")
	//  struct
	p1 := Person{
		first: "John",
		last:  "Doe",
		age:   30,
	}

	p2 := Person{
		first: "Jane",
		last:  "Smith",
		age:   25,
	}

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Printf("Person 1: %+v\n", p1)
	fmt.Println(p1.first, p2.first)

	fmt.Println("..................................")

	sa1 := SecretAgent{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   40,
		},
		licenseToKill: true,
	}

	fmt.Println(sa1)

}

// uses same underlying array
// everything is pass by value in go
// the value is being passed into this function
// and then assigned to siua
func medianOne(x []float64) float64 {
	sort.Float64s(x) // Sort the slice
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(siua []float64) float64 {
	n := make([]float64, len(siua))
	copy(n, siua)    // Copy the slice
	sort.Float64s(n) // Sort the copy
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
	}
	return (n[i-1] + n[i]) / 2
}

// Variadic function
func sum(nums ...int) int {
	fmt.Println("Numbers:", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
	return total
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	
	}
	defer file.Close() // Ensure the file is closed when done
	s:= []byte("Hello, World!\n")
	_, err = file.Write(s)
	if err != nil {
		log.Fatalf("error writing to file: %v", err)
	}
	return file, nil
}
