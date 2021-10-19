package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := NewClient()
	if err != nil {
		return
	}

	cmdExecutor := &CommandExecutor{client}

	executeCommands(cmdExecutor)
}

func executeCommands(cmdExecutor *CommandExecutor) {
	for {
		input := showMenuAndGetInput()
		switch input {
		case 0:
			break
		case 1:
			cmdExecutor.GetRemoteTime()
			break
		case 2:
			cmdExecutor.GetRemoteScreenshot()
			break
		default:
			log.Fatalln("Invalid input. Try again!")
		}
	}
}

func showMenuAndGetInput() (input int) {
	menuMsg := "Main menu\n1: Get remote time\n2: Take remote screenshot\n0: Quit"
	fmt.Println(menuMsg)
	fmt.Scan(&input)
	return
}
