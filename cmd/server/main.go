package main

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

const addr = "localhost:6666"

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
}
