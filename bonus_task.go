package main

import (
	"fmt"
)

func sendNumbers(even, odd, errChan chan int) {
	for i := 1; i <= 22; i++ {
		if i > 20 {
			errChan <- i
		} else if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	close(errChan)
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	errChan := make(chan int)

	go sendNumbers(even, odd, errChan)

	// Process all values from all channels
	for even != nil || odd != nil || errChan != nil {
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
		case err, ok := <-errChan:
			if ok {
				fmt.Printf("Error: number %d is greater than 20\n", err)
			} else {
				errChan = nil // Mark the channel as closed
			}
		}
	}
}
