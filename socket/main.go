package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
  fmt.Println("Hello, World!")

  server := socketio.NewServer(nil);

  server.OnConnect("/", func(s socketio.Conn) error {
    s.SetContext("")
    fmt.Println("connected: ", s.ID())
    return nil
  })

  server.OnEvent("/chat", "chat", func(s socketio.Conn, msg string) string {
    log.Println("Received a message from client: " + msg)
    s.SetContext(msg)
    return "recv" + msg
  })

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  http.Handle("/socket.io/", server)
  log.Println("Serving at localhost:5000...")
  log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
