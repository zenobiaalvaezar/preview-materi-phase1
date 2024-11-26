package main

import (
	"fmt"
)

func sendNumbers(even, odd chan int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func main() {
	even := make(chan int)
	odd := make(chan int)

	go sendNumbers(even, odd)

	// Process all values from both channels
	for even != nil || odd != nil {
		select {
		case num, ok := <-even:
			if ok {
				fmt.Printf("Received an even number: %d\n", num)
			} else {
				even = nil // Mark the channel as closed
			}
		case num, ok := <-odd:
			if ok {
				fmt.Printf("Received an odd number: %d\n", num)
			} else {
				odd = nil // Mark the channel as closed
			}
		}
	}
}
