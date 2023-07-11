package Content

import "fmt"

func FactorialCheck() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(i, "!=")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	factorial(c, quit)
}

func factorial(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = x*y, y+1
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
