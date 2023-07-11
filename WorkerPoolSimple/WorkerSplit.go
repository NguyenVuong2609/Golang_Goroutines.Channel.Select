package WorkerPoolSimple

import (
	"fmt"
	"time"
)

func WorkerSplit() {
	w1 := workerSplit{name: "A"}
	w2 := workerSplit{name: "B"}
	w3 := workerSplit{name: "C"}
	w4 := workerSplit{name: "D"}
	w5 := workerSplit{name: "E"}
	//List Worker
	workers := []workerSplit{w1, w2, w3, w4, w5}

	//List All Work
	workList := make(chan int, 20)
	for i := 1; i <= 20; i++ {
		workList <- i
	}
	//close(workList)

	//List Task for Worker
	sizeWork := 4
	var taskList []chan int
	for i := 0; i < len(workers); i++ {
		taskList = append(taskList, make(chan int, sizeWork))
	}

	//Share Work for a Worker
	for j := 0; j < len(taskList); j++ {
		//fmt.Println("Task list =====>", len(taskList[j]))
		if len(taskList[j]) < cap(taskList[j]) {
			taskList[j] <- <-workList
			j--
			continue
		}
		close(taskList[j])
	}

	//All Worker working
	for i := 0; i < len(workers); i++ {
		//for j := 0; j < cap(taskList[i]); j++ {
		//	go workers[i].Work(taskList[i])
		//	//fmt.Println("Task list =====>", j, "With length =>>>", len(taskList[i]))
		//	time.Sleep(time.Second)
		//}
		for t := range taskList[i] {
			go workers[i].Work1(t)
		}
	}
	time.Sleep(time.Second * 10)
}

type workerSplit struct {
	name string
}

type work interface {
	Work(channel chan int)
}

func (w workerSplit) Work(channel chan int) {
	fmt.Printf("worker %s is working ===> %d\n", w.name, <-channel)
}
func (w workerSplit) Work1(channel int) {
	fmt.Printf("worker %s is working ===> %d\n", w.name, channel)
}
