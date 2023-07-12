package WorkerPoolSimple

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

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	//time.Sleep(1 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup, workerId int, job Job) {
	//for job := range jobs {
	output := Result{job, digits(job.randomno), workerId}
	results <- output
	//}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	check := true
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		select {
		case v, ok := <-jobs:
			if ok {
				go worker(&wg, i, v)
			}
			if !ok {
				check = false
			}
		}
		if !check {
			break
		}
		if i == noOfWorkers-1 {
			i = -1
		}
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result() {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d, WorkerID: %d\n", result.job.id, result.job.randomno, result.sumofdigits, result.workerid)
	}
}
func WorkerPoolSimple() {
	var wg = sync.WaitGroup{}
	startTime := time.Now()
	noOfJobs := 20
	wg.Add(2)
	go allocate(noOfJobs)
	go result()
	noOfWorkers := 5
	createWorkerPool(noOfWorkers)
	wg.Done()
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
