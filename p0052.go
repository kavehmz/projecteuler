package main

import (
	"fmt"
	"math"
)

func digitsSignature(n int) int {
	r := 0
	for n != 0 {
		d := n - (n/10)*10
		if (r & int(math.Exp2(float64(d)))) == 0 {
			r += int(math.Exp2(float64(d)))
		}
		n = (n - d) / 10
	}
	return r
}

func main() {

	fmt.Println(digitsSignature(112))
	x := 1
	for {
		d1 := digitsSignature(1 * x)
		d2 := digitsSignature(2 * x)
		d3 := digitsSignature(3 * x)
		d4 := digitsSignature(4 * x)
		d5 := digitsSignature(5 * x)
		d6 := digitsSignature(6 * x)
		if d1 == d2 && d2 == d3 && d3 == d4 && d4 == d5 && d5 == d6 {
			fmt.Println("Number is:", x)
			break
		}
		x++
	}

}
