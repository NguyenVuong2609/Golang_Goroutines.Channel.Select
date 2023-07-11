package Content

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Goroutines() {
	go sayHello("Vuong")

	sayHello("Viet Nam")
	time.Sleep(time.Second)
	fmt.Println("------------------ Capture variables ------------------")
	for i := 1; i <= 20; i++ {
		go func(value int) {
			fmt.Println(value)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("------------------ Gosched ------------------------")
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("I am Goroutine 1")
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("I am Goroutine 2")
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("--------------------- WaitGroup ------------------------")
	var wg sync.WaitGroup
	wg.Add(3) /////// adds an entry to the waitgroup counter
	go hello(&wg)
	go bye(&wg)
	go thank(&wg)
	wg.Wait() ////// blocks execution until the goroutine finishes
	fmt.Println("main function")
	fmt.Println("-------------------- Mutex --------------------------")
	Mutex()
}

func sayHello(name string) {
	for i := 0; i <= 2; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}
func hello(wgrp *sync.WaitGroup) {
	fmt.Println("Hello")
	wgrp.Done() /////// notifies the waitgroup that it finished
}
func bye(wgrp *sync.WaitGroup) {
	fmt.Println("Bye!!!")
	wgrp.Done()
}
func thank(wgrp *sync.WaitGroup) {
	fmt.Println("Thanks!!!")
	wgrp.Done()
}
