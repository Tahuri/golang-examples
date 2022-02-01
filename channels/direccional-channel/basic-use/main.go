package main

import (
	"fmt"
	"time"
)

func getNext(mychan chan<- int) {
	next := 0
	for x := 0; x <= 5; x++ {
		mychan <- next
		next++
	}
	close(mychan)
}

func getNextString(mychan chan<- string) {
	next := 10
	for x := 0; x <= 10; x++ {
		mychan <- fmt.Sprintf("%v", next)
		next++
	}
	close(mychan)
}

func printValues(mychanInt <-chan int, mychanString <-chan string) {
	for {
		select {
		case myInt, ok := <-mychanInt:
			_ = ok
			fmt.Println("MyInt", myInt)
			time.Sleep(time.Second)
		case myString, ok := <-mychanString:
			_ = ok
			fmt.Println("MyString", myString)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan string)
	input := ""
	go getNext(ch)
	go getNextString(ch2)
	go printValues(ch, ch2)
	fmt.Scanln(&input)
	fmt.Println("End of Program", input)
}
