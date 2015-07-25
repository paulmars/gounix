package main

import (
  "code.google.com/p/go.net/websocket"
  "io"
  "net/http"
  "fmt"
)

func echoHandler(ws *websocket.Conn) {

  receivedtext := make([]byte, 100)

   n,err := ws.Read(receivedtext)

  if err != nil {
    fmt.Printf("Received: %d bytes\n",n)
  }

  s := string(receivedtext[:n])
  fmt.Printf("Received: %d bytes: %s\n",n,s)
  io.Copy(ws, ws)
  fmt.Printf("Sent: %s\n",s)
}

func main() {
  http.Handle("/echo", websocket.Handler(echoHandler))
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}