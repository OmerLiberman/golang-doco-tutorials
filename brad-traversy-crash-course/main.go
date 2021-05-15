package main

import (
	"fmt"
	"math"
	"net/http"
)

func funcExample(name string) string {
	return "Hello " + name;
}

func arrays() {
	var fruits [2] string
	fruits[0] = "Apple"
	fmt.Println(fruits)

	fruits2 := [] string {"Apple", "Orange", "Grapes"}
	fmt.Println(fruits2, len(fruits2))
	fmt.Println(fruits2[1: 2])
}

func conditionals() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d %d", x, y)
	} else if x == y {
		fmt.Printf("")
	} else {

	}

	switch x {
	case 1:
		fmt.Println("4")
	default:
		fmt.Println("Def")
	}
}

func loops() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func maps() {
	emails := make(map[string]string)

	emails["Omer"] = "liberman.omer1@gmail.com"
	fmt.Println(emails)
	delete(emails, "Omer")
	fmt.Println(emails)

	emails2 := map[string]string{"Omer": "oliberman@synamedia.com"}
	fmt.Println(emails2)
}

func ranges() {
	ids := []int{10, 11, 12, 13}

	for i, id := range ids {
		fmt.Println(i, id)
	}
	for _, id := range ids {
		fmt.Println(id)
	}

	emails := map[string]string{"Omer": "oliberman@synamedia.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s", k, v)
	}
}

func pointers() {
	a := 5
	b := &a

	fmt.Println(a, *b, b)

	*b = 10
	fmt.Println(a, *b, b)
}

func closures() func(int) int {
	s := 0
	return func (x int ) int {
		s += x
		return s
	}
}

type Person struct {
	first string
	last string
	age int
}

func (p Person) greet() string {
	return "Hello " + p.first;
}

func (p *Person) bDay() {
	p.age++;
}

func structures() {
	p := Person{"Omer", "Liber", 26}
	fmt.Println(p, p.first)
	p.greet()
	p.bDay()
	fmt.Println(p)
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func interfaces() {
	c := Circle{0, 0, 1}
	r := Rectangle{10, 10}

	fmt.Println(c.area())
	fmt.Println(r.area())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello !")
}

func web() {
	http.HandleFunc("/", index)
	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}


func main() {
	// Data types:
	// bool, int, int8, int16, int32, int64
	// uint, int, uint8, uint16, uint32, uint64
	// byte
	// rune
	// float32, float65
	// complex64, complex128

	//var name string = "Omer"
	//fmt.Println(name)
	//name2 := "Danielle"
	//fmt.Println(name2)
	//fmt.Printf("%T\n", name)

	//const isCool = true
	//isCool = false // error !

	//arrays()
	//conditionals()
	//loops()
	//maps()
	//ranges()
	//pointers()

	//sum := closures()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(sum(i))
	//}

	//structures()
	//interfaces()

	web()
}
