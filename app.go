package main

// Script Name:	shell-go
// Author:		Pablo Andrade
// Created: 	19/04/2024
// Version: 	0.1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")

		// Read the keyboard input.
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("path required")
		}

		// Change the directory and return the error
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()

}
