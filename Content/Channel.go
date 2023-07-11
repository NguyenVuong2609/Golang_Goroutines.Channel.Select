package Content

import (
	"fmt"
	"time"
)

func Channel() {
	myChan := make(chan int)

	go func() {
		myChan <- 1
	}()

	fmt.Println("data received:", <-myChan)
	fmt.Println("--------------------------------------------")
	go func() {
		for i := 1; i <= 5; i++ {
			myChan <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(<-myChan)
	}
	fmt.Println("-------------- Close channel -----------------")
	close(myChan)
	v, ok := <-myChan
	fmt.Println("Received ", v, ok)
	fmt.Println("-------------- For range ------------------")
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
	fmt.Println("----------------- Buffer Channel -------------------")
	bufferedChan := make(chan int, 3)

	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	bufferedChan <- 1
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	bufferedChan <- 2
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	bufferedChan <- 3
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))
	fmt.Println("----------------- Get data from Buffer Channel ----------------")
	bufferedChan2 := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		bufferedChan2 <- i
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-bufferedChan2, "Buffer chan 2")
	}
}

func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
