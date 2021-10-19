package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"os"
)

// Command is an interface which provides the basic functions
// to send command and execute the response
type Command interface {
	Cmd() string
	Execute([]byte, error)
}

// CommandsMap is the map of commands
var CommandsMap = map[int]Command{
	1: cCurrentTime{},
	2: cScreenshot{},
}

// *************** current time *********************
type cCurrentTime struct{}

func (c cCurrentTime) Cmd() string {
	return string("CTME")
}

func (c cCurrentTime) Execute(data []byte, err error) {
	if err != nil {
		log.Fatalln("Error in getting current time, error:", err)
		return
	}
	log.Println("Current remote time:", string(data))
}

// *************** Screenshot *********************
type cScreenshot struct{}

func (c cScreenshot) Cmd() string {
	return string("CSHT")
}

func (c cScreenshot) Execute(data []byte, err error) {
	if err != nil {
		log.Fatalln("Error in remote screenshot, error:", err)
		return
	}
	reader := bytes.NewReader(data)
	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatalln("Error in decoding screenshot, error:", err.Error())
		return
	}

	f, err := os.Create("screenshot.jpg")
	if err != nil {
		log.Fatalln("Error in saving screenshot, error:", err.Error())
		return
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		log.Fatalln("Error in saving screenshot, error:", err.Error())
	}
}
