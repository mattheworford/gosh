package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("gosh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		var command = parts[0]
		var args = parts[1:]

		switch command {
		case "":
			continue
		case "exit":
			os.Exit(0)
		case "cd":
			var path string
			if len(args) > 0 {
				path = args[0]
			} else {
				path, _ = os.UserHomeDir()
			}

			var err = os.Chdir(path)
			if err != nil {
				fmt.Println(strings.Replace(err.Error(), "chdir", "cd", -1))
			}
		default:
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
