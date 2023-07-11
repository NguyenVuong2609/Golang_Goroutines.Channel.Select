package Content

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
	workerid    int
}

var results = make(chan Result, 10)
var jobs1 = make(chan Job, 5)
var jobs2 = make(chan Job, 5)
var jobs3 = make(chan Job, 5)
var jobs4 = make(chan Job, 5)
var jobs5 = make(chan Job, 5)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	//time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup, num int) {
	var jobsPool = make(chan Job, 5)
	switch {
	case num == 0:
		jobsPool = jobs1
	case num == 1:
		jobsPool = jobs2
	case num == 2:
		jobsPool = jobs3
	case num == 3:
		jobsPool = jobs4
	case num == 4:
		jobsPool = jobs5
	}
	for job := range jobsPool {
		output := Result{job, digits(job.randomno), num}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs, noOfWorkers int) {
	for i, j := 0, 0; i < noOfJobs && j < noOfWorkers; i, j = i+1, j+1 {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		switch {
		case j == 0:
			jobs1 <- job
		case j == 1:
			jobs2 <- job
		case j == 2:
			jobs3 <- job
		case j == 3:
			jobs4 <- job
		case j == 4:
			jobs5 <- job
			j = -1
		}
	}
	close(jobs1)
	close(jobs2)
	close(jobs3)
	close(jobs4)
	close(jobs5)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d, WorkerID: %d\n", result.job.id, result.job.randomno, result.sumofdigits, result.workerid)
	}
	done <- true
}
func WorkerPool() {
	startTime := time.Now()
	noOfJobs := 22
	noOfWorkers := 5
	go allocate(noOfJobs, noOfWorkers)
	done := make(chan bool)
	go result(done)
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
