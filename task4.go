package main

import "fmt"

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // Dapat mengirim hingga buffer penuh
	}
	close(ch) // Menutup channel setelah selesai
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Println(num) // Membaca nilai dari channel
	}
}

func main() {
	ch := make(chan int, 5) // Buffered channel dengan kapasitas 5
	go produce(ch)          // Memulai producer
	consume(ch)             // Memulai consumer
}
