package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	basedir := "C:\\Users\\Lucas\\my_stuff\\code\\just_for_fun\\askiicmd\\"
	for {
		var a string
		fmt.Printf("enter command:")
		fmt.Scanf("%s\n", &a)
		fmt.Printf("Ok printing %s\n", a)

		f, err := os.Open(basedir + "askiicodes.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		// Splits on newlines by default.
		scanner := bufio.NewScanner(f)

		line := 1
		x := 0
		// https://golang.org/pkg/bufio/#Scanner.Scan
		//analizing input
		if a == "list" {
			fmt.Println("The available commands are:")
			for scanner.Scan() {
				split := strings.Split(scanner.Text(), "|")
				fmt.Printf("%s\n", split[0])
			}
		} else {
			for scanner.Scan() {
				split := strings.Split(scanner.Text(), "|")
				if split[0] == a {
					print(basedir, split[1])
					x++
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

func print(basedir string, link string) {
	ascii, err := ioutil.ReadFile(basedir + "askii\\" + link)
	if err != nil {
		fmt.Printf("sorry file not found,\n err = %s\n", err)
		return
	}
	fmt.Println(string(ascii))
}
