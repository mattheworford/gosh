package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func cd(args []string) {
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
}

func pwd(writer *io.PipeWriter) {
	defer func(writer *io.PipeWriter) {
		err := writer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(writer)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintln(writer, dir)
	if err != nil {
		fmt.Println(err)
	}
}
