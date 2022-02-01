package main

import (
	"fmt"
	"time"
)

func getNext(mychan chan<- int) {
	next := 0
	for x := 0; x <= 5; x++ {
		mychan <- next
		fmt.Println("Msg to mychanIn sended:", next)
		next++
	}
	close(mychan)
}

func getNextString(mychan chan<- string) {
	next := 10
	for x := 0; x <= 10; x++ {
		mychan <- fmt.Sprintf("%v", next)
		fmt.Println("Msg to mychanString sended:[", next, "]")
		next++
	}
	close(mychan)
}

func printValues(mychanInt <-chan int, mychanString <-chan string) {
	mychanIntOk := false
	mychanStringOk := false
	for completed := false; !completed; {
		select {
		case myInt, ok := <-mychanInt:
			if mychanIntOk && mychanStringOk {
				completed = true
			}
			if !ok {
				mychanIntOk = true
				continue
			}
			fmt.Println("MyInt", myInt)
			time.Sleep(time.Second)
		case myString, ok := <-mychanString:
			if mychanIntOk && mychanStringOk {
				completed = true
			}
			if !ok {
				mychanStringOk = true
				continue
			}
			fmt.Println("MyString", myString)
			time.Sleep(time.Second)
		}
	}
	fmt.Println("out of for")
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
