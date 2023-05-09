package main

import (
	"log"
	"net/http"
)

var cmd Cmd
var server http.Server

func StartServer(bind string, remote string) {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	handler := &handler{reverseProxy: remote}
	server.Addr = bind
	server.Handler = handler
	//go func() {
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//}()
}

// func StopServer() {
// 	if err := server.Shutdown(nil); err != nil {
// 		log.Println(err)
// 	}
// }

func main() {
	cmd = parseCmd()
	StartServer(cmd.bind, cmd.remote)
}
