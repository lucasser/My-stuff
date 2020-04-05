package main

import "fmt"

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
