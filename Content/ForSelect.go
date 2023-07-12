package Content

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done"
}

func ForSelect() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(time.Second)
		select {
		case v := <-ch:
			fmt.Println("Received: ", v)
		}
	}
}
