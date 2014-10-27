package main

import (
	"fmt"
	"strconv"
)

func sieve(num int) []int {
	prime := make([]bool, num+1)
	for i := 2; i < len(prime); i++ {
		prime[i] = true
	}
	result := []int{}
	for i := 2; i < len(prime); i++ {
		if prime[i] {
			result = append(result, i)
			for j := i + i; j < len(prime); j += i {
				prime[j] = false
			}
		}
	}

	return result
}

func getNumDivider(num, divider int) int {
	tmpNum := num
	count := 0
	for {
		if tmpNum%divider == 0 {
			tmpNum = tmpNum / divider
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	const a1, a2 = 2, 100
	const b1, b2 = 2, 100
	allPrimes := sieve(a2)
	fmt.Println(allPrimes)
	result := make(map[string]struct{})
	for i := a1; i <= a2; i++ {
		for j := b1; j <= b2; j++ {
			str := ""
			for _, p := range allPrimes {
				str += strconv.Itoa(p) + ":"
				count := getNumDivider(i, p)
				str += strconv.Itoa(count * j)
				str += ","
			}
			result[str] = struct{}{}
		}
	}
	fmt.Println(len(result))
}
