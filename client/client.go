package main

import (
	"bufio"
	"log"
	"net"
)

const (
	host     = "localhost"
	port     = "3030"
	connType = "tcp"
)

// Client struct
type Client struct {
	conn net.Conn
}

// NewClient creates and return new client
func NewClient() (*Client, error) {
	client := &Client{}

	var err error
	client.conn, err = net.Dial(connType, host+":"+port)
	if err != nil {
		log.Fatalln("Error connecting:", err.Error())
		return nil, err
	}

	return client, nil
}

// SendCommand send command to remote to execute
func (c *Client) SendCommand(command string) ([]byte, error) {
	log.Println("Text to send:", command)
	c.conn.Write([]byte(command + "\n"))

	return bufio.NewReader(c.conn).ReadBytes(byte('\n'))
}
