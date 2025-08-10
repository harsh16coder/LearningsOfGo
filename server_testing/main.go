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
	hasfirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hassecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error encouttered while reading: %s", err)
	}
	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body=\n%s\n",
		ctx.Value(keyServerAddr),
		hasfirst, first,
		hassecond, second, body)
	io.WriteString(w, "Hello User")
}

func getUserName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got /user request \n", ctx.Value(keyServerAddr))
	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
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
	<-ctx.Done()
}
