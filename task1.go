package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}
func printLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)
}
