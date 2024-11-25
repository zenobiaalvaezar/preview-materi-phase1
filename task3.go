package main

import "fmt"

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // Menunggu consumer membaca sebelum lanjut
	}
	close(ch) // Menutup channel setelah selesai
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Println(num) // Membaca nilai dari channel
	}
}

func main() {
	ch := make(chan int) // Unbuffered channel
	go produce(ch)       // Memulai producer
	consume(ch)          // Memulai consumer
}
