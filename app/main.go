package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Uncomment this block to pass the first stage
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
		default:
			fmt.Printf("%s: command not found\n", userCommand)
		}

	}
}
