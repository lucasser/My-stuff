package main

import (
	"fmt"
	"math"
	"sync"
)

func add(n, m uint) uint {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	m = m - 1
	n = n + 1
	return add(n, m)
}

func subtract(n, m uint) uint {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	m = m - 1
	n = n - 1
	return subtract(n, m)
}

func times(n, m uint) uint {
	if m == 0 {
		return 0
	}
	if m == 1 {
		return n
	}
	return add(n, times(n, m-1))
}

func power(n, m uint) uint {
	if m == 1 {
		return n
	}
	return times(n, power(n, m-1))
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
	}
	var c = 0
	for n >= m {
		n = n - m
		c = c + 1
	}
	var cs = fmt.Sprintf("%d", c)
	var ns = fmt.Sprintf("%d", n)
	return cs + " R " + ns
}
func factorial(n uint) uint {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	return times(n, factorial(n-1))
}

func lcm(n, m uint) uint {
	if m == 1 {
		return n
	}
	if n == 1 {
		return m
	}
	return n * m / gcd(n, m)
}
func factor(n uint, numCPU uint) []uint {
	if n == 1 {
		return []uint{1}
	}
	if n == 2 {
		return []uint{1, 2}
	}
	res := []uint{}
	var (
		start, stop, min, max, perCPU uint
	)

	wg := sync.WaitGroup{}
	mux := sync.RWMutex{}
	h := float64(n)
	start = 1
	s := math.Sqrt(h)
	stop = uint(s) + 1
	perCPU = (stop) / numCPU
	if perCPU == 0 {
		perCPU = 1
	}
	fmt.Println(perCPU)
	min = start
	max = start + perCPU
	for {
		wg.Add(1)
		if max > stop {
			max = stop
		}
		go func(min, max uint) {
			fmt.Printf("Starting %v to %v\n", min, max)
			defer wg.Done()
			facts := factorRange(n, min, max)
			fmt.Printf("Got %d factors for %v to %v\n", len(facts), min, max)

			mux.Lock()
			defer mux.Unlock()
			res = append(res, facts...)
		}(min, max)
		min += perCPU
		max += perCPU
		if min > stop || perCPU == 0 {
			break
		}
		if max > stop {
			max = stop
		}
	}
	wg.Wait()
	res = append(res)
	return res
}

func factorRange(n uint, min, max uint) []uint {
	var res []uint
	m := uint(min)
	var k uint
	for m < max {
		if n%m == 0 {
			k = n / m
			if k == m {
				res = append(res, m)
				break
			}
			res = append(res, m, k)
		}
		m++
	}
	return res
}
