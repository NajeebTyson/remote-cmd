package core

import (
	"bufio"
	"log"
	"net"
)

const port = "3030"

// Server struct to create server
type Server struct {
	isRunning bool
}

// NewServer creates a new Server and return it
func NewServer() *Server {
	return &Server{}
}

// Stop stop server and close connection
func (s *Server) Stop() {
	s.isRunning = false
}

// Run creates a server and listen for new connections
func (s *Server) Run() error {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalln("Unable to create server, error:", err.Error())
		return err
	}
	log.Println("Server is now listening on", port)

	defer l.Close()

	s.isRunning = true
	for s.isRunning {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln("Error connecting client, error:", err.Error())
			s.isRunning = false
			return err
		}

		go s.handleConnection(c)
	}
	return nil
}

func (s *Server) handleConnection(conn net.Conn) {
	log.Println("Client connected", conn.RemoteAddr().String())
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Println("Client connection closed", conn.RemoteAddr().String())
			conn.Close()
			return
		}
		command := string(buffer[:len(buffer)-1])
		log.Println("command:", command)
		res, _ := commandsMap[command]
		tm, _ := res.Execute()
		conn.Write(append(tm, byte('\n')))
	}
}
