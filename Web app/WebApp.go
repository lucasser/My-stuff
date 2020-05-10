package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	openbrowser("http://localhost:8090/")
	mainServer()
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

func mainFactor() {
	numCPU := uint(runtime.NumCPU())
	for {
		var a uint
		fmt.Printf("Enter an interger to factor: ")
		fmt.Scanf("%d\n", &a)

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
