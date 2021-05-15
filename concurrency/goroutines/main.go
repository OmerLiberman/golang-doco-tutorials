package main

import (
	"./task"
	"fmt"
	"time"
)


func main() {
	var tasks [50]task.Task

	for i := 0; i < 50; i++ {
		tasks[i] = task.Task{Index: i, IsDone: false}
	}

	fmt.Println("Start running tasks ...")
	for i := 0; i < 50; i++ {
		go tasks[i].DoTask()
	}


	time.Sleep(10000000)
	fmt.Println("Done running tasks.")
}
