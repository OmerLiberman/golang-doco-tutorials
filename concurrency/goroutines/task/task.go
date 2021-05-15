package task

import (
	"fmt"
	"time"
)

type Task struct {
	IsDone bool
	Index int
}

func (t *Task) DoTask() int {
	fmt.Printf("Start with task: %d\n", t.Index)
	t.IsDone = true
	time.Sleep(20000)
	fmt.Printf("Done with task: %d\n", t.Index)
	return t.Index
}
