package main

import (
	"Day4-Goroutines-Channel/WorkerBufferChannel"
	"fmt"
)

func main() {
	fmt.Println("-------------------- Goroutines -----------------------")
	//Content.Goroutines()
	fmt.Println("-------------------- Channel -----------------------")
	//Content.Channel()
	fmt.Println("-------------------- Buffer Channel Worker Queue --------------------")
	//Content.BufferChannel()
	fmt.Println("-------------------- More Channel ---------------------")
	//Content.MoreChannel()
	fmt.Println("-------------------- Select Quit Channel ---------------------")
	//Content.SelectQuitChannel()
	//Content.FactorialCheck()
	fmt.Println("-------------------- Worker Pool --------------------------")
	//Content.WorkerPool()
	//WorkerPoolSimple.WorkerPoolSimple()
	//WorkerPoolSimple.WorkerSplit()
	WorkerBufferChannel.WorkerBuffer()
	fmt.Println("-------------------- For Select ---------------------------")
	//Content.ForSelect()
	fmt.Println("-------------------- Mutex Race Condition ---------------------------")
	//Content.MutexRaceCondition()
	fmt.Println("---------------------- Wait Group ---------------------------")
	//Content.WaitGroup()
}
