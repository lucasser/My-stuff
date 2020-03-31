package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// TODO: read cmd argument
	cmd := os.Args[1]
	// call main_cmd() or
	if cmd == "cmd" {
		main_cmd()
	} else {
		main_server()
	}
	// call main_server()
}

func main_cmd() {
	for {
		var o string
		var a, b uint

		fmt.Printf("Enter an interger: ")
		fmt.Scanf("%d\n", &a)
		fmt.Printf("Oh, I got %d\n", a)

		fmt.Printf("Enter another interger: ")
		fmt.Scanf("%d\n", &b)
		fmt.Printf("Now I got %d\n", b)

		for {
			fmt.Println("enter an operation: ")
			fmt.Scanf("%s\n", &o)

			if o == "!" {
				timeIt(func() {
					fmt.Println(factorial(a))
				})
				break
			}
			if o == "/" {
				fmt.Println(divide(a, b))
				break
			}
			if o == "^" {
				fmt.Println(power(a, b))
				break
			}
			if o == "%" {
				fmt.Println(remeinder(a, b))
				break
			}
			if o == "gcd" {
				fmt.Println(gcd(a, b))
				break
			}
			if o == "tn" {
				fmt.Println(tnumber(a))
				break
			}
			if o == "+" {
				fmt.Println(add(a, b))
				break
			}
			if o == "*" {
				fmt.Println(times(a, b))
				break
			}
			if o == "lucas" {
				fmt.Print("panic: runtime error: index out of range [1] with length 1")
			} else {
				fmt.Printf("why did you give me %s\n", o)
				fmt.Println("enter lucas")
			}
		}
	}
}

func timeIt(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	dur := end.Sub(start)
	fmt.Printf("Time: %v\n", dur)
}
