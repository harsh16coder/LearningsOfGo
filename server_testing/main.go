package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got / request \n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello User")
}

func getUserName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got /user request \n", ctx.Value(keyServerAddr))
	io.WriteString(w, "My name is Harsh")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHello)
	mux.HandleFunc("/user", getUserName)
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, keyServerAddr, l.Addr().String())
		},
	}
	serverTwo := &http.Server{
		Addr:    ":8081",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, keyServerAddr, l.Addr().String())
		},
	}
	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed \n")
		} else if err != nil {
			fmt.Printf("error starting the server: %s\n", err)
			os.Exit(1)
		}
		cancelCtx()
	}()
	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed \n")
		} else if err != nil {
			fmt.Printf("error starting the server: %s\n", err)
			os.Exit(1)
		}
		cancelCtx()
	}()
	<-ctx.Done()
}
