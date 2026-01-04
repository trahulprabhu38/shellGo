package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handlePrint() {
	fmt.Println("testing new feature")
}

func handleEcho(args string) {
	fmt.Println(args)
}

func handleType(cmd string) {
	builtins := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
	}

	if builtins[cmd] {
		fmt.Println(cmd + " is a shell builtin")
	} else {
		entries, err := os.ReadDir(".")
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			if entry.Name == cmd {

			}
		}

		fmt.Println(cmd + ": not found")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			continue
		}

		cmd := strings.TrimSpace(line)
		if cmd == "" {
			continue
		}

		if cmd == "exit" {
			return
		}

		if strings.HasPrefix(cmd, "echo ") {
			handleEcho(strings.TrimSpace(cmd[5:]))
			continue
		}

		if strings.HasPrefix(cmd, "type ") {
			handleType(strings.TrimSpace(cmd[5:]))
			continue
		}

		fmt.Println(cmd + ": command not found")
	}
}
