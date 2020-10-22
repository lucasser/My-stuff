/*// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
	"strconv"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "GET":
		fallthrough
	case "POST":
		fmt.Printf("GOT POST: %+v\n", r)
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseMultipartForm(10000); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("got %v", r.Form)
		n, err := strconv.ParseUint(r.FormValue("firstn"), 10, 32)
		m, err2 := strconv.ParseUint(r.FormValue("lastn"), 10, 32)
		o := r.FormValue("op")
		fmt.Printf("parsed %v %v %v", n, m, o)
		ans, err3 := computePrep(uint(n), uint(m), o)
		fmt.Fprintf(w, ans)

		if err != nil {
			fmt.Printf("ParseForm() err: %v", err)
			return
		}
		if err2 != nil {
			fmt.Printf("parseForm() err2: %v", err2)
			return
		}
		if err3 != nil {
			fmt.Printf("parseForm() err3: %v", err3)
			return
		}
		// address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
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
func add(n, m uint) uint {
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

func subtract(n, m uint) uint {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	} else {
		m = m - 1
		n = n - 1
	}
	return subtract(n, m)
}

func times(n, m uint) uint {
	if m == 0 {
		return 0
	}
	if m == 1 {
		return n
	} else {
		return add(n, times(n, m-1))
	}
}

func power(n, m uint) uint {
	if m == 1 {
		return n
	} else {
		return times(n, power(n, m-1))
	}
}

func remeinder(n, m uint) uint {
	for n >= m {
		n -= m
	}
	return n
}

func gcd(m, n uint) uint {
	if n > m {
		return gcd(n, m)
	}
	r := m % n
	if r == 0 {
		return n
	}
	return gcd(n, r)
}

func tnumber(n uint) uint {
	if n == 1 {
		return 1
	}
	if n < 0 {
		fmt.Print("invalid entry")
	}
	return add(n, tnumber(n-1))
}

func divide(n, m uint) string {
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
		return cs + " R " + ns
	}
}
func factorial(n uint) uint {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 1
	} else {
		return times(n, factorial(n-1))
	}
}

func lcm(n, m uint) uint {
	if m == 1 {
		return n
	}
	if n == 1 {
		return m
	} else {
		return n * m / gcd(n, m)
	}
}
*/