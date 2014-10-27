package main

import (
	"fmt"
	"os"
)

func sieve(num int) []int {
	isPrimes := make([]bool, num+1)
	// set default to true
	for i := 2; i < len(isPrimes); i++ {
		isPrimes[i] = true
	}
	// find prime
	primeList := []int{}
	for i := 2; i <= num; i++ {
		if isPrimes[i] {
			// add to prime list
			primeList = append(primeList, i)
			// set to false
			for j := i + i; j <= num; j += i {
				isPrimes[j] = false
			}
		}
	}
	return primeList
}

func getPrimeFactors(num int, primeList []int) []int {
	pFactors := []int{}
	for _, p := range primeList {
		if p > num/2 {
			break
		}
		if num%p == 0 {
			pFactors = append(pFactors, p)
		}
	}
	return pFactors
}

func main() {
	primeList := sieve(10000)
	// find consecutive
	const primeCount = 4
	const consecCount = 4
	i := 2
	consecNums := []int{}
	for {
		if len(getPrimeFactors(i, primeList)) == primeCount {
			consecNums = append(consecNums, i)
			if len(consecNums) == consecCount {
				fmt.Println(consecNums)
				os.Exit(0)
			}
		} else {
			consecNums = []int{}
		}
		i++
	}
}
