package main

import (
	"fmt"
	"time"
	"runtime"
)

type Task struct {
	ID string
}

func (t *Task) Process() bool {
	fmt.Println("process ", t.ID)
	return true
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var TASK_CHANNEL = make(chan Task, 10)



//控制生产和消费的速度，当队列满时，通过select处理超出容量的情况
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
		time.Sleep(1*time.Second)
		i++
		var task = Task{
			ID: fmt.Sprintf("%s_%d", Now(), i),
		}

		select {
		case TASK_CHANNEL <- task:
			fmt.Println("add to queue", task.ID)
		default:
			fmt.Printf("error to add %d\n", i)
		}
	}
}

func consume() {
	for {
		time.Sleep(3*time.Second)
		select {
		case task := <- TASK_CHANNEL:
			fmt.Printf("cunsume task: %s\n", task.ID)
		default:
			fmt.Println("none to consume")
		}
	}
}
