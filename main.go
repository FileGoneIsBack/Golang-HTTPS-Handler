package main

import (
	"api/core"
	"api/core/master"
	"os/exec"
	"os"
)

func main() {
	// Start the configs
	core.Initialize() // controls any objectS YOU need to add to config.json for further complex

	// Clear the screen
	clearScreen()

	master.NewV2()
	s.Stop()
}

//clear screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
