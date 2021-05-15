package main

import (
	"../oop/employee"
)

func main() {
	e := employee.New("Omer", "Liberman", 26)

	e.GetData()
	e.AddDayOffs(2)
	e.GetData()
	e.DecreaseDayOff(-5)
	e.GetData()
}

