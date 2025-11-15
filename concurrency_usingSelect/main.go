package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// for i := 1; i <= 5; i++ {
	// 	go func(id int) {
	// 		ch1 <- id
	// 	}(i)
	// }
	// for i := 6; i <= 10; i++ {
	// 	go func(id int) {
	// 		ch2 <- id
	// 	}(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	select {
	// 	case data, ok := <-ch1:
	// 		if ok {
	// 			fmt.Printf("Received %d from ch1\n", data)
	// 		} else {
	// 			fmt.Println("Ch1 is closed")
	// 		}
	// 	case data, ok := <-ch2:
	// 		if ok {
	// 			fmt.Printf("Received %d from ch2\n", data)
	// 		} else {
	// 			fmt.Println("Ch2 is closed")
	// 		}
	// 	}

	// }

	//Listening to Multiple Channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start Goroutine 1 to send data to ch1
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	// Start Goroutine 2 to send data to ch2
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
	}()

	// Main Goroutine receives and prints data from ch1 and ch2
	for i := 0; i < 10; i++ {
		select {
		case data := <-ch1:
			fmt.Println("Received from ch1:", data)
		case data := <-ch2:
			fmt.Println("Received from ch2:", data)
		}
	}

	fmt.Println("Done.")
}
