package main

import (
	"fmt"
	"time"
)

func getNext(mychan chan int) {
	next := 0
	for x := 0; x <= 5; x++ {
		mychan <- next
		next++
	}
}

func printValues(mychan chan int) {
	for {
		value := <-mychan
		fmt.Println(value)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	input := ""
	go getNext(ch)
	go printValues(ch)
	fmt.Scanln(&input)
	fmt.Println("End of Program", input)
}
