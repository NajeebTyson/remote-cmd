package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/NajeebTyson/remote-cmd/server/core"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := core.NewServer()
	go handleSignal(server, sigs)
	server.Run()

	log.Println("Exiting server")
}

func handleSignal(server *core.Server, sigs <-chan os.Signal) {
	select {
	case <-sigs:
		log.Fatalln("Server interrupted")
		server.Stop()
	}
}
