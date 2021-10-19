package main

// CommandExecutor is the main class to execute remote commands
type CommandExecutor struct {
	client *Client
}

func (c *CommandExecutor) executeCommand(cmdNo int) {
	cmd, _ := CommandsMap[cmdNo]
	res, err := c.client.SendCommand(cmd.Cmd())
	cmd.Execute(res, err)
}

// GetRemoteTime Get the remote time and display it
func (c *CommandExecutor) GetRemoteTime() {
	c.executeCommand(1)
}

// GetRemoteScreenshot get the remote screenshot
func (c *CommandExecutor) GetRemoteScreenshot() {
	c.executeCommand(2)
}
