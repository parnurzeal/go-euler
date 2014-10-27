package main

import (
	"fmt"
	"math"
)

func isPrime(num int64) bool {
	switch {
	case num == 1:
		return false
	case num < 4:
		return true
	case num%2 == 0:
		return false
	case num < 9:
		return true
	case num%3 == 0:
		return false
	}
	sqrt := int64(math.Sqrt(float64(num)))
	for i := int64(5); i < sqrt; i += 6 {
		if num%i == 0 {
			return false
		}
		if num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var i int64
	fmt.Scanf("%d", &i)
	fmt.Println(isPrime(i))
}
