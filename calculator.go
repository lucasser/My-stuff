package main

import (
	"fmt"
)

func divide(n,m int) string {
        if m == 0 {
        	return "invalid entry"
        }
	if m == 1 {
		var ns = fmt.Sprintf("%d", n)
        	return ns
        } else {

	var c = 0
        for n >= m {
        	n = n - m
        	c = c + 1
        }
	var cs = fmt.Sprintf("%d", c)
	var ns = fmt.Sprintf("%d", n)
	return cs + " R" + ns
        }
}
func factorial(n int) int {
        if n == 1 {
        	return 1
        }
	if n == 0 {
        	return 1
        } else {
        	return n * factorial(n - 1)
	}
}
func main(){
	for {
		var o string
		var a, b int

	  fmt.Printf("Enter an interger: ")
		fmt.Scanf("%d\n", &a)
		fmt.Printf("Oh, I got %d\n", a)

		fmt.Printf("Enter another interger: ")
		fmt.Scanf("%d\n", &b)
		fmt.Printf("Now I got %d\n", b)


		fmt.Println("enter an operation: ")
		fmt.Scanf("%s\n", &o)

		if o == "f" {
			fmt.Println(factorial(a))
		} else {
			if o == "d" {
				fmt.Println(divide(a, b))
			} else {
				fmt.Printf("why did you give me %s\n", o)
				fmt.Println("enter f or d for operation")
			}
		}
	}
}
