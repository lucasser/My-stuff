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
}

func computePrep(a, b uint, o string) (string, error) {
	if o == "!" {
		return fmt.Sprint(factorial(a)), nil
	}
	if o == "/" {
		return divide(a, b), nil
	}
	if o == "^" {
		return fmt.Sprint(power(a, b)), nil
	}
	if o == "%" {
		return fmt.Sprint(remeinder(a, b)), nil
	}
	if o == "gcd" {
		return fmt.Sprint(gcd(a, b)), nil
	}
	if o == "tn" {
		return fmt.Sprint(tnumber(a)), nil
	}
	if o == "+" {
		return fmt.Sprint(add(a, b)), nil
	}
	if o == "*" {
		return fmt.Sprint(times(a, b)), nil
	}
	return "", fmt.Errorf("unknown operation")
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
			res, err := computePrep(a, b, o)
			if err == nil {
				fmt.Println(res)
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
