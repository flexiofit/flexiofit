// server.go

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	mainFile := "cmd/server/main.go"

	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		log.Fatalf("Error: %s does not exist", mainFile)
	}

	cmd := exec.Command("go", "run", mainFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
