package main

import (
	"fmt"
	"runtime"
	"time"
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

//接收发送给channel之后返回的结果: 在任务上加一channel用于接收结果
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go produce()

	go consume()

	var closed = make(chan bool)
	<-closed

}

func produce() {
	var i int = 0

	for {
		time.Sleep(1 * time.Second)
		i++
		var task = Task{
			ID:         fmt.Sprintf("%s_%d", Now(), i),
			ResultChan: make(chan *TaskResult),
		}

		TASK_CHANNEL <- task
		result := <- task.ResultChan

		fmt.Printf("add to queue %s, and get result: %v\n",task.ID, result)
	}
}

func consume() {
	for {
		time.Sleep(3 * time.Second)
		select {
		case task := <-TASK_CHANNEL:
			fmt.Printf("cunsume task: %s\n", task.ID)
			task.ResultChan <- &TaskResult{
				Result: fmt.Sprintf("result %s is ok", task.ID),
			}
		default:
			fmt.Println("none to consume")
		}
	}
}
