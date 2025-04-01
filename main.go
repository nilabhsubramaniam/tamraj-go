package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
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
}
