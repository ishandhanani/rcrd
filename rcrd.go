package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/creack/pty"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "n", "", "Name for the recording file")
	flag.StringVar(&fileName, "name", "", "Name for the recording file")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	rcrdDir := filepath.Join(homeDir, ".rcrd")
	err = os.MkdirAll(rcrdDir, 0755)
	if err != nil {
		fmt.Println("Error creating .rcrd directory:", err)
		return
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/zsh" // Default to zsh if $SHELL is not set
	}

	cmd := exec.Command(shell, "-l") // "-l" starts a login shell
	cmd.Env = append(os.Environ(), "RCRD_ACTIVE=1")

	// Set the user's home as the working directory
	cmd.Dir = homeDir

	// Set up pseudo-terminal
	ptmx, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("Error starting pty:", err)
		return
	}
	defer ptmx.Close()

	// Create log file for this session
	var logFileName string
	if fileName != "" {
		logFileName = fileName + ".txt"
	} else {
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		logFileName = timestamp + ".txt"
	}
	logFile, err := os.Create(filepath.Join(rcrdDir, logFileName))
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	// Handle shell output
	go func() {
		io.Copy(io.MultiWriter(os.Stdout, logFile), ptmx)
	}()

	// Handle user input
	go func() {
		io.Copy(ptmx, os.Stdin)
	}()

	cmd.Wait()
}
