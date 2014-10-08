package main

import (
	"fmt"
	"math"
)

func sieve(num int) (primes []int) {
	nonPrimes := make([]bool, num+1)
	nonPrimes[0] = true
	for i := 2; i <= num; i++ {
		if !nonPrimes[i] {
			for j := i + i; j <= num; j += i {
				nonPrimes[j] = true
			}
		}
	}
	for i := 1; i <= num; i++ {
		if !nonPrimes[i] {
			primes = append(primes, i)
		}
	}
	return
}

func isPrime(num int64) bool {
	switch {
	case num == 1:
		return false
	case num == 2:
		return true
	case num%2 == 0:
		return false
	case num < 9:
		return true
	}
	sqrt := int64(math.Sqrt(float64(num)))
	for i := int64(5); i < sqrt; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	primes := sieve(100)
	fmt.Println(primes)
	for _, p := range primes {
		fmt.Print(isPrime(int64(p)), " ")
	}
}
