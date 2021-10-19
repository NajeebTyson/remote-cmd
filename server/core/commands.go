package core

import (
	"bytes"
	"errors"
	"image/jpeg"
	"log"
	"time"

	"github.com/vova616/screenshot"
)

// CommandResult is an interface to provide function
// Execute which must be implemented by all the command executor
type CommandResult interface {
	Execute() ([]byte, error)
}

var commandsMap = map[string]CommandResult{
	"CTME": cCurrentTime{},
	"CSHT": cScreenshot{},
}

// *************** current time *********************
type cCurrentTime struct{}

func (c cCurrentTime) Execute() ([]byte, error) {
	return []byte(time.Now().Format(time.RFC850)), nil
}

// *************** Screenshot *********************
type cScreenshot struct{}

func (c cScreenshot) Execute() ([]byte, error) {
	cError := errors.New("Error in taking screenshot")
	img, err := screenshot.CaptureScreen()
	if err != nil {
		log.Fatalln("Error in taking screenshot, error:", err.Error())
		return nil, cError
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Fatalln("Error in encoding image, error:", err.Error())
		return nil, cError
	}
	return buf.Bytes(), nil
}
