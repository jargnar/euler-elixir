package main

import (
	"math"
)

func nthprime() {
	n := 10001
	lim := int(math.Log(float64(2*n)) * float64(2*n))

	isPrime := make(map[int]bool)
	for i := 2; i <= lim; i++ {
		isPrime[i] = true
	}

	counter := 0
loop:
	for i := 2; i <= lim; i++ {
		if isPrime[i] == true {
			counter = counter + 1
			if counter == n {
				println(i)
				break loop
			}
			for j := i; j <= lim; j = j + i {
				isPrime[j] = false
			}
		}
	}
}