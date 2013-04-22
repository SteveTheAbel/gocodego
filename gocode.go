package main

// Welcome to Go! Code!
// To review our goals are:
// 1. Learn golang!
// 2. Do cool things!
// 3. Impress children!
//
// Ready!
// Set!
// Go!

import "fmt"
import "time"

func daniel_dressler() {
	fmt.Printf("Hello from Calgary!\n")
	time.Sleep(2990 * time.Millisecond) // Please decrement as needed!
}

func SteveTheAbel_FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Printf("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz")
		} else {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n");
	}
}

func main() {
	SteveTheAbel_FizzBuzz()
	daniel_dressler()
	fmt.Printf("Hope to see you again!\n")
}
