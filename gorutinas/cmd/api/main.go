package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/timeoutok", withTimeoutOk)
	http.HandleFunc("/withwaitgroup", withWaitGroup)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func withTimeoutOk(w http.ResponseWriter, req *http.Request) {
	go writeMessageInGoRutina(time.Second*1, "WwithTimeoutOk-1")
	go writeMessageInGoRutina(time.Second*2, "WwithTimeoutOk-2")
	go writeMessageInGoRutina(time.Second*3, "WwithTimeoutOk-3")
	go writeMessageInGoRutina(time.Second*4, "WwithTimeoutOk-4")
	go writeMessageInGoRutina(time.Second*5, "WwithTimeoutOk-5")
	go writeMessageInGoRutina(time.Second*6, "WwithTimeoutOk-6")
	go writeMessageInGoRutina(time.Second*7, "WwithTimeoutOk-7")
	go writeMessageInGoRutina(time.Second*8, "WwithTimeoutOk-8")
	go writeMessageInGoRutina(time.Second*9, "WwithTimeoutOk-9")
	time.Sleep(time.Second * 5)
	fmt.Println("Finish - withTimeoutOk")
	_, _ = io.WriteString(w, `{"message": "With timeout OK"}`)
}

func writeMessageInGoRutina(sleep time.Duration, message string) {
	writeMessageInGoRutinaWithWaitWroup(sleep, message, nil)
}

func writeMessageInGoRutinaWithWaitWroup(sleep time.Duration, message string, wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}
	time.Sleep(sleep)
	msg := fmt.Sprintf("In Gorutina: %s", message)
	go writeMessageInSubGoRutina(sleep+time.Second*2, msg, wg)
	fmt.Println(msg)
}

func writeMessageInSubGoRutina(sleep time.Duration, message string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	time.Sleep(sleep)
	fmt.Printf("In SubGorutina: %s \n", message)
}

func withWaitGroup(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	wg.Add(8)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*1, "withWaitGroup-1", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*2, "withWaitGroup-2", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*3, "withWaitGroup-3", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*4, "withWaitGroup-4", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*5, "withWaitGroup-5", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*6, "withWaitGroup-6", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*7, "withWaitGroup-7", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*8, "withWaitGroup-8", &wg)
	go writeMessageInGoRutinaWithWaitWroup(time.Second*9, "withWaitGroup-9", nil)
	wg.Wait()
	fmt.Println("Finish - withWaitGroup")
	_, _ = io.WriteString(w, `{"message": "With withWaitGroup"}`)
}
