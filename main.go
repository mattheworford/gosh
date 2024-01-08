package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("gosh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		commands := strings.Split(input, "|")

		var cmds []*exec.Cmd
		var output io.ReadCloser

		for _, command := range commands {
			command = strings.TrimSpace(command)
			parts := strings.Fields(command)
			var command = parts[0]
			var args = parts[1:]

			switch command {
			case "":
				continue
			case "exit":
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
