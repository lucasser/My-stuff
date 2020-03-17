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
func main(){
	fmt.Println(divide(1, 5))
}
