package main

import (
	"fmt"
	"runtime"
	"time"
	"math"
	"random"
)

type TaskResult struct {
	Result string
}

type Task struct {
	ID         string
	ResultChan chan *TaskResult
}

func (t *Task) Process() bool {
	fmt.Println("process ", t.ID)
	return true
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var TASK_CHANNEL = make(chan Task, 10)

//获取结果的超时控制, 利用select+time.After
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go produce()

	go consume()

//	var closed = make(chan bool)
//	<-closed

	time.Sleep(math.MaxInt32*time.Second)

}

func produce() {
	var i int = 0

	for {
		i++

		time.Sleep(1 * time.Second)
		go func(i int){
			var task = Task{
				ID:         fmt.Sprintf("%s_%d", Now(), i),
				ResultChan: make(chan *TaskResult),
			}

			TASK_CHANNEL <- task

			//获取结果的超时控制, 利用select+time.After
			select {
			case result := <-task.ResultChan:
				fmt.Printf("正常返回，add to queue %s, and get result: %v\n", task.ID, result)
			case <-time.After(2 * time.Second):
				fmt.Printf("任务%s获取结果超时\n", task.ID)
			}

		}(i)

	}
}

func consume() {
	for {


		task := <-TASK_CHANNEL
		fmt.Printf("cunsume task: %s\n", task.ID)

		time.Sleep( 1 * time.Second)
		task.ResultChan <- &TaskResult{
			Result: fmt.Sprintf("result %s is ok", task.ID),
		}

	}
}
