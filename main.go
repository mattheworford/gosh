package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func main() {
	signal.Ignore(os.Interrupt)

	history := InitHistory()

	for {
		fmt.Print("gosh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		commands := strings.Split(input, "|")

		var cmds []*exec.Cmd
		var output io.ReadCloser

		for _, commandLine := range commands {
			commandLine = strings.TrimSpace(commandLine)
			parts := strings.Fields(commandLine)
			command := parts[0]
			var args = parts[1:]

			switch command {
			case "":
				continue
			case "exit":
				_ = history.file.Close()
				os.Exit(0)
			case "cd":
				cd(args)
			case "pwd":
				pipeReader, pipeWriter := io.Pipe()
				go pwd(pipeWriter)
				if len(commands) == 1 {
					_, err := io.Copy(os.Stdout, pipeReader)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					output = pipeReader
				}
			case "history":
				for i, hc := range history.commands {
					fmt.Printf("%d: %s\n", i, hc)
				}
			default:
				cmd := exec.Command(command, args...)
				cmd.Stderr = os.Stderr

				cmds = append(cmds, cmd)

				if output != nil {
					cmd.Stdin = output
				}
				var err error
				output, err = cmd.StdoutPipe()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			history.Append(commandLine)
		}
		if len(cmds) > 0 {
			cmds[len(cmds)-1].Stdout = os.Stdout
		}

		for _, cmd := range cmds {
			err := cmd.Start()
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		for _, cmd := range cmds {
			err := cmd.Wait()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
