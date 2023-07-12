package WorkerBufferChannel

import (
	"Day4-Goroutines-Channel/Model"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Model.Job, 10)
var results = make(chan Model.Result, 10)
var workRoom = make(chan Model.Worker, 5)
var workerSlice = []Model.Worker{
	{1, "Chuan"},
	{2, "Vuong"},
	{3, "Yen"},
	{4, "Son"},
	{5, "Thu"},
	{6, "Ha"},
	{7, "Van"},
	{8, "Huy"},
}
var nextWorker = cap(workRoom)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	return sum
}

// Create Jobs
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Model.Job{Id: i, Randomno: randomno}
		jobs <- job
	}
	close(jobs)
}
func worker(wg *sync.WaitGroup, worker Model.Worker, job Model.Job) {
	output := Model.Result{Job: job, Sum: digits(job.Randomno), Worker: worker}
	results <- output
	wg.Done()
}
func create1stGroupWorker() {
	for i := 0; i < cap(workRoom); i++ {
		worker := workerSlice[i]
		workRoom <- worker
	}
}
func createWorkerPool(nextWorker int) {
	var wg sync.WaitGroup
	check := true
	for i := 0; i < len(workRoom); i++ {
		select {
		case v, ok := <-jobs:
			if !ok {
				check = false
			} else {
				wg.Add(1)
				go worker(&wg, <-workRoom, v)
				workRoom <- workerSlice[nextWorker]
				if nextWorker == len(workerSlice)-1 {
					nextWorker = 0
				} else {
					nextWorker++
				}
			}
		}
		if !check {
			break
		}
		if i == len(workRoom)-1 {
			i = -1
		}
	}
	wg.Wait()
	close(results)
}
func result() {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d, WorkerID: %d, WorkerName: %s\n", result.Job.Id, result.Job.Randomno, result.Sum, result.Worker.Id, result.Worker.Name)
	}
}

func WorkerBuffer() {
	var wg = sync.WaitGroup{}
	startTime := time.Now()
	wg.Add(2)
	create1stGroupWorker()
	noOfJobs := 19
	go allocate(noOfJobs)
	go result()
	createWorkerPool(nextWorker)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
