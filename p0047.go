package main

import "fmt"

// prime factors
func f(n int) []int {
	var m []int
	for i := 2; i < n; i++ {
		if n%i == 0 {
			m = append(m, i)
			m1 := f(n / i)
			if len(m1) > 0 {
				m = append(m, m1...)
			} else {
				m = append(m, n/i)
			}
			return m
		}
	}
	return []int{}
}

// count factors
func c(l []int) int {
	m := make(map[int]int)
	count := 0
	for i := 0; i < len(l); i++ {
		if _, ok := m[l[i]]; ok {
			m[l[i]]++
		} else {
			m[l[i]] = l[i]
			count++
		}
	}
	return count
}

func main() {
	b := false
	i := 2
	for !b {
		if c(f(i)) == 4 && c(f(i+1)) == 4 && c(f(i+2)) == 4 && c(f(i+3)) == 4 {
			fmt.Println(i)
			b = true
		}
		i++
	}
}
