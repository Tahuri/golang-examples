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
	rawCtx := req.Context()
	// ctx, cancelFunc := context.WithCancelCause(rawCtx)
	// ctx, cancelFunc := context.WithCancelCause(context.Background())
	// ctx, _ := context.WithCancelCause(context.Background())
	ctx, _ := context.WithCancelCause(rawCtx)
	go writeMessageInGoRutina(ctx, time.Second*1, "WwithTimeoutOk-1")
	go writeMessageInGoRutina(ctx, time.Second*2, "WwithTimeoutOk-2")
	go writeMessageInGoRutina(ctx, time.Second*3, "WwithTimeoutOk-3")
	go writeMessageInGoRutina(ctx, time.Second*4, "WwithTimeoutOk-4")
	go writeMessageInGoRutina(ctx, time.Second*5, "WwithTimeoutOk-5")
	go writeMessageInGoRutina(ctx, time.Second*6, "WwithTimeoutOk-6")
	go writeMessageInGoRutina(ctx, time.Second*7, "WwithTimeoutOk-7")
	go writeMessageInGoRutina(ctx, time.Second*8, "WwithTimeoutOk-8")
	go writeMessageInGoRutina(ctx, time.Second*9, "WwithTimeoutOk-9")
	// cancelFunc(errors.New("test-causa"))
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
		err := context.Cause(ctx)
		log.Printf("In SubGorutina %s - Processed - %s - Cause: %s", message, ctx.Err(), err)
	case <-ctx.Done():
		err := context.Cause(ctx)
		log.Printf("In SubGorutina %s - Cancelled - %s - Cause: %s", message, ctx.Err(), err)
	}
}
