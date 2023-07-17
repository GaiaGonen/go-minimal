package main

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func EchoServer(ws *websocket.Conn) {
	_, err := io.Copy(ws, ws)
	if err != nil {
		return
	}
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", websocket.Handler(EchoServer))
	if err != nil {
		panic("ListenAndServer: " + err.Error())
	}

}
