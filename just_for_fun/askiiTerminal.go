package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		var a string
		fmt.Printf("enter command:")
		fmt.Scanf("%s\n", &a)
		fmt.Printf("Ok printing %s\n", a)

		f, _ := os.Open("askiicodes.txt")
		defer f.Close()

		// Splits on newlines by default.
		scanner := bufio.NewScanner(f)

		line := 1
		x := 0
		// https://golang.org/pkg/bufio/#Scanner.Scan
		//analizing input
		if a == "list" {
			ba, _ := readFile("askiicodes.txt")
			fmt.Printf("The available commands are\n%s\n", ba)
		} else {

			for scanner.Scan() {
				if strings.Contains(scanner.Text(), a) {
					x++
					fmt.Printf("%d\n", line)
				}

				line++
			}
			if x == 0 {
				fmt.Println("command not found\nif you want to know all the commands, type 'list'")
			}
			if err := scanner.Err(); err != nil {
				// Handle the error
			}
		}
	}
}
