package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
}

func commandType(command string) string {
	if _, exist := builtins[command]; exist {
		return fmt.Sprintf("%s is a shell builtin", command)
	}
	return fmt.Sprintf("%s: not found", command)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v\n", err)
		}
		cmd := strings.Split(strings.TrimSpace(input), " ")
		userCommand := cmd[0]
		args := cmd[1:]
		switch userCommand {
		case "exit":
			exitCode, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Error converting exit code: %v\n", err)
			}
			os.Exit(exitCode)
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			if len(args) > 0 {
				command := commandType(args[0])
				fmt.Println(command)
			}
		default:
			fmt.Printf("%s: command not found\n", userCommand)
		}
	}
}
