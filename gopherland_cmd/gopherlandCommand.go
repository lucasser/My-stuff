package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd string
	for x in os.Args[] {
		cmd += os.Args[x]
	}
	fmt.Println(cmd)
	var program string
	x := 0
	fmt.Printf("enter program:\n")
	fmt.Scanf("%s\n", &program)
	fmt.Printf("Oh, I got %s\n", program)

	if program == "list" {
		fmt.Println("The available commands are:\n Factorize, Calculate, Askii")
		x++
	} else {
		cmnd := exec.Command("askiicmd.exe")
		cmnd.Run()
		//cmnd.Start()
		//log.Println("log")
		x++
	}
	if x == 0 {
		fmt.Println("Program not found :(\n Please type 'list' to view all commands.")
	}
}
