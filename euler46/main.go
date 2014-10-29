package main

import "fmt"

func sieve(num int) intSlice {
	nonPrime := make([]bool, num+1)
	allPrime := []int{}
	for i := 2; i <= num; i++ {
		if !nonPrime[i] {
			allPrime = append(allPrime, i)
			for j := 2 * i; j <= num; j += i {
				nonPrime[j] = true
			}
		}
	}
	return allPrime
}

func twiceSquare(num int) intSlice {
	allNum := intSlice{}
	for i := 1; i*i*2 <= num; i++ {
		allNum = append(allNum, i*i*2)
	}
	return allNum
}

type intSlice []int

func (slice intSlice) contain(num int) bool {
	for _, n := range slice {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	const max = 1000000
	prime := sieve(max)
	allSquare := twiceSquare(max)
	for i := 33; i < max; i += 2 {
		if !prime.contain(i) {
			var conj bool
			for _, tSqr := range allSquare {
				if tSqr > i {
					break
				}
				if prime.contain(i - tSqr) {
					conj = true
					break
				}
			}
			if !conj {
				fmt.Println(i)
				return
			}
		}
	}
}
