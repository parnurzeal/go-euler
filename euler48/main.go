package main

import "fmt"

func main() {
	var sum int64
	for i := int64(1); i <= 1000; i++ {
		var tmp int64 = 1
		for j := int64(0); j < i; j++ {
			tmp *= i
			tmp = tmp % 10000000000
		}
		sum += tmp
	}
	fmt.Println(sum % 10000000000)
}
