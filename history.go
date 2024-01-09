package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type History struct {
	commands []string
	file     *os.File
}

func InitHistory() *History {
	home, _ := os.UserHomeDir()
	filename := filepath.Join(home, ".gosh_history")
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)

	fmt.Println(filename)
	scanner := bufio.NewScanner(file)
	var commands []string

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	return &History{commands: commands, file: file}
}

func (history *History) Append(command string) {
	toIgnore := map[string]bool{"exit": true, "history": true}
	_, ignoreCommand := toIgnore[command]
	if ignoreCommand {
		return
	}
	history.commands = append(history.commands, command)
	_, _ = history.file.WriteString(command + "\n")
}
