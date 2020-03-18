package main

import (
	"fmt"
)
func add(n,m uint) uint {
	if n == 0 {
		return m
	} else {
		if m == 0 {
			 return n
		 } else {
			 m = m - 1
			 n = n + 1
		 }
	 }
	return add(n, m)
}
func times(n,m uint) uint {
	if m == 0 {
		return 0
	}
	if m == 1 {
		return n
	} else {
		return add(n, times(n, m - 1))
	}
}
func divide(n,m uint) string {
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
func factorial(n uint) uint {
        if n == 1 {
        	return 1
        }
	if n == 0 {
        	return 1
        } else {
        	return times(n, factorial(n - 1))
	}
}
func main(){
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
				fmt.Println(factorial(a))
				break
			}
			if o == "/" {
				fmt.Println(divide(a, b))
				break
			}
			if o == "+" {
				fmt.Println(add(a, b))
				break
			}
			if o == "*" {
				fmt.Println(times(a, b))
				break
			} else {
					fmt.Printf("why did you give me %s\n", o)
					fmt.Println("enter +, /, ! for operation")
			}
	}
}
}
