package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// TODO: read cmd argument
	cmd := os.Args[1]
	// call main_cmd() or
	if cmd == "cmd" {
		mainCmd()
	}
	if cmd == "fact" {
		mainFactor()
	}
	if cmd == "lucas" {
		println("panic: runtime error: index out of range [1] with length 1\n\ngoroutine 1 [running]:\nmain.main()\n\t\aC:/Users/Marianna/Dropbox/Lucas/code/calculator.go:11 +0x1ba\n\n")
	} else {
		fmt.Print("wrong command\n")
		fmt.Println("enter 'code.exe lucas'")
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
	if o == "lcm" {
		return fmt.Sprint(lcm(a, b)), nil
	}
	if o == "*" {
		return fmt.Sprint(times(a, b)), nil
	}
	return "", fmt.Errorf("unknown operation")
}

func mainCmd() {
	for {
		var o string
		var a, b uint

		fmt.Printf("Enter an integer: ")
		fmt.Scanf("%d\n", &a)
		fmt.Printf("Oh, I got %d\n", a)

		fmt.Printf("Enter another integer: ")
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
				fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n your computer crashed\n")
				break
			} else {
				fmt.Printf("why did you give me %s\n", o)
				fmt.Println("enter lucas")
			}
		}
	}
}
func mainFactor() {
	numCPU := uint(runtime.NumCPU())
	fmt.Printf("This computer has %d CPUs\n", numCPU)
	//numCPU = 2
	for {
		var a uint
		fmt.Printf("Enter an integer to factor: ")
		fmt.Scanf("%d\n", &a)
		fmt.Printf("Enter number of CPUs to use: ")
		fmt.Scanf("%d\n", &numCPU)

		var factors []uint
		dur := timeIt(func() {
			factors = factor(a, numCPU)
		})
		for _, f := range factors {
			fmt.Println(f)
		}
		fmt.Printf("Computed in %v\n", dur)

		numFactors := len(factors)
		fmt.Printf("t(%v) = %v\n", a, numFactors)

		//productOfFactors := uint(math.Pow(float64(a), float64(numFactors/2)))
		//fmt.Printf("p(%v) = %v\n", a, productOfFactors)
	}
}

func timeIt(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	dur := end.Sub(start)
	return dur
}
