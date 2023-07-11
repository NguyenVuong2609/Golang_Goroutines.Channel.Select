package Content

import (
	"fmt"
	"sync"
)

func Mutex() {
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit(20, &wg)
	go withdraw(80, &wg)
	go deposit(40, &wg)
	wg.Wait()
	fmt.Printf("Balance is: %d\n", balance)
}

var (
	mutex   sync.Mutex
	balance int
)

func deposit(val int, wg *sync.WaitGroup) {
	mutex.Lock() // lock
	balance += val
	mutex.Unlock() // unlock
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mutex.Lock() // lock
	balance -= val
	mutex.Unlock() // unlock
	wg.Done()
}

func init() {
	balance = 100
}
