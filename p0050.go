package main

import "fmt"
import "math"

const Max = 1000000

var primes []int = find_primes()

func is_prime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func find_primes() []int {
	var p []int
	for i := 2; i < Max; i++ {
		if is_prime(i) {
			p = append(p, i)
		}
		if i%10000 == 0 {
			fmt.Println(i)
			fmt.Println(len(p))
		}
	}
	return p
}

func main() {
	fmt.Println(len(primes))
	for l := 21; l < len(primes); l++ {
		for offset := 0; offset < len(primes)-l; offset++ {
			var s int = 0
			for i := offset; i < offset+l; i++ {
				s += primes[i]
			}
			if s >= Max {
				break
			}
			if is_prime(s) {
				fmt.Println(primes[offset], l)
				fmt.Println(s)
			}
		}

	}
}
