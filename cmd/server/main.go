package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const addr = "localhost:9943"

var globalWsServer *websocket.Server

func main() {
	log.Println("Creating websocket server")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v\n", addr, err)
	}
	log.Printf("Listening on %s\n", listener.Addr().String())
	globalWsServer = &websocket.Server{
		Handler: handleWebsocket,
		Handshake: func(c *websocket.Config, r *http.Request) error {
			return nil
		},
	}
	httpServer := &http.Server{
		Handler: http.HandlerFunc(handleHTTP),
	}
	err = httpServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve HTTP: %v\n", err)
	}
}

func handleHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP: %s at %s\n", r.Method, r.URL.Path)
	globalWsServer.ServeHTTP(rw, r)
}

func handleWebsocket(c *websocket.Conn) {
	log.Printf("Handling Websocket\n")
	for {
		delayNumber, err := rand.Int(rand.Reader, big.NewInt(1000))
		delay := time.Duration(delayNumber.Int64()) * time.Millisecond
		if err != nil {
			log.Printf("Error generating new delay, sleeping for 5 seconds: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		time.Sleep(delay)
		fmt.Fprintf(c, "Waited for %s", delay)
	}
}
