package main

import (
	"api/core"
	"api/core/master"
	"os/exec"
	"os"
	"github.com/janeczku/go-spinner"
)

func main() {
	// Start the configs
	s := spinner.StartNew("Initializing")
	core.Initialize()
	s.Stop()

	// Clear the screen
	clearScreen()

	// Start the webserver
	s = spinner.StartNew("Webserver Started")
	master.NewV2()
	s.Stop()
}

//clear screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
