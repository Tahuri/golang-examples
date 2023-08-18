package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CtxKey string

const (
	ctxMainKey = CtxKey("MainKey")
	ctxSubKey  = CtxKey("SubMainKey")
)

func main() {
	http.HandleFunc("/timeoutok", withTimeoutOk)
	log.Println("Before Up Server")
	log.Println(http.ListenAndServe(":8080", nil))
}

func withTimeoutOk(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-1"), time.Second*1)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-2"), time.Second*2)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-3"), time.Second*3)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-4"), time.Second*4)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-5"), time.Second*5)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-6"), time.Second*6)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-7"), time.Second*7)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-8"), time.Second*8)
	go writeMessageInGoRutina(context.WithValue(ctx, ctxMainKey, "WwithTimeoutOk-9"), time.Second*9)
	time.Sleep(time.Second * 30)
	log.Println("Finish - withTimeoutOk")
	_, _ = io.WriteString(w, `{"message": "With timeout OK"}`)
}

func writeMessageInGoRutina(ctx context.Context, sleep time.Duration) {
	time.Sleep(sleep)
	message := (ctx.Value(ctxMainKey)).(string)
	msg := fmt.Sprintf("In Gorutina: %s", message)
	go writeMessageInSubGoRutina(context.WithValue(ctx, ctxSubKey, msg), sleep+time.Second*2)
	log.Println(msg)
}

func writeMessageInSubGoRutina(ctx context.Context, sleep time.Duration) {
	message := (ctx.Value(ctxSubKey)).(string)
	select {
	case <-time.After(sleep):
		log.Printf("In SubGorutina %s - Processed - %s", message, ctx.Err())
	case <-ctx.Done():
		log.Printf("In SubGorutina %s - Cancelled - %s", message, ctx.Err())
	}
}
