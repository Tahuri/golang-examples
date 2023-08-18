package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/timeoutok", withTimeoutOk)
	log.Println("Before Up Server")
	log.Println(http.ListenAndServe(":8080", nil))
}

func withTimeoutOk(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	go writeMessageInGoRutina(ctx, time.Second*1, "WwithTimeoutOk-1")
	go writeMessageInGoRutina(ctx, time.Second*2, "WwithTimeoutOk-2")
	go writeMessageInGoRutina(ctx, time.Second*3, "WwithTimeoutOk-3")
	go writeMessageInGoRutina(ctx, time.Second*4, "WwithTimeoutOk-4")
	go writeMessageInGoRutina(ctx, time.Second*5, "WwithTimeoutOk-5")
	go writeMessageInGoRutina(ctx, time.Second*6, "WwithTimeoutOk-6")
	go writeMessageInGoRutina(ctx, time.Second*7, "WwithTimeoutOk-7")
	go writeMessageInGoRutina(ctx, time.Second*8, "WwithTimeoutOk-8")
	go writeMessageInGoRutina(ctx, time.Second*9, "WwithTimeoutOk-9")
	time.Sleep(time.Second * 30)
	log.Println("Finish - withTimeoutOk")
	_, _ = io.WriteString(w, `{"message": "With timeout OK"}`)
}

func writeMessageInGoRutina(ctx context.Context, sleep time.Duration, message string) {
	time.Sleep(sleep)
	msg := fmt.Sprintf("In Gorutina: %s", message)
	go writeMessageInSubGoRutina(ctx, sleep+time.Second*2, msg)
	log.Println(msg)
}

func writeMessageInSubGoRutina(ctx context.Context, sleep time.Duration, message string) {
	select {
	case <-time.After(sleep):
		log.Printf("In SubGorutina %s - Processed", message)
	case <-ctx.Done():
		log.Printf("In SubGorutina %s - Cancelled", message)
	}
}
