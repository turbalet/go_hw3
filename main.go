package main

import (
	"fmt"
	"github.com/turbalet/hw3/task2"
	"github.com/turbalet/hw3/task3"
	"github.com/turbalet/hw3/task4"
	"github.com/turbalet/hw3/task5"
	"github.com/turbalet/hw3/task6"
	"github.com/turbalet/hw3/task7"
)

func main() {
	fmt.Println("Itoa: ", task2.Itoa(-322))
	fmt.Println(task3.Atoi("32er"))
	fmt.Println("Reversed string: ", task4.Reverse("абвг dj"))

	// sorting imports of itoa.go
	task5.SortImports("C:\\Users\\acer\\go\\src\\test\\task2\\itoa.go")

	// task 6 - fibonacci
	generator := task6.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")
	}
	fmt.Println()

	// task 7
	s := "some text"
	i := 12
	res, err := task7.RuneByIndex(&s, &i)
	fmt.Println( "RuneByIndex: ",string(res), err)
}
