package main

import (
	"fmt"
	"strconv"
)

var primeList = make(map[int]bool)

func sieve(num int) {
	for i := 2; i <= num; i++ {
		if _, ok := primeList[i]; !ok {
			primeList[i] = true
			for j := i + i; j <= num; j += i {
				primeList[j] = false
			}
		}
	}
}

func isCircular(num int) bool {
	str := []byte(strconv.Itoa(num))
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str))
		if !primeList[num] {
			return false
		}
		tmpC := str[0]
		for j := 0; j < len(str)-1; j++ {
			str[j] = str[j+1]
		}
		str[len(str)-1] = tmpC
	}
	return true
}

func main() {
	sieve(1000000)
	//fmt.Println(primeList)
	count := 0
	for i := 2; i < 1000000; i++ {
		if isCircular(i) {
			count++
		}
	}
	fmt.Println(count)
}
