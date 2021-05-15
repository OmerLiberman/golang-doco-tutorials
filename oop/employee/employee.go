package employee

import (
	"fmt"
)

type employee struct {
	FirstName string
	LastName string
	DayOffs int
}

func New(firstName, lastName string, daysOff int) employee {
	e := employee{firstName, lastName, daysOff}
	return e
}

func (e employee) GetRemainingDayOffs() int {
	return e.DayOffs
}

func (e *employee) AddDayOffs(addition int) {
	e.DayOffs += addition;
}

func (e *employee) DecreaseDayOff(deduction int) {
	e.DayOffs -= deduction;
}

func (e employee) GetData() {
	fmt.Printf("employee: First name - %s, Last name - %s, Remaining Days off - %d\n", e.FirstName, e.LastName, e.DayOffs)
}



