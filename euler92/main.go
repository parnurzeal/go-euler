package main

import "fmt"

func findNext(num int) int {
	var newNum int
	for num > 0 {
		digit := num % 10
		newNum += digit * digit
		num /= 10
	}
	return newNum
}

func main() {
	// remember right numbers
	rmbNum := make(map[int]int)
	only89 := []int{}
	rmbNum[89] = 89
	rmbNum[1] = 1
	for i := 1; i < 10000000; i++ {
		yes := false
		rmbSeq := []int{}
		for num := i; ; num = findNext(num) {
			rmbSeq = append(rmbSeq, num)
			if val, ok := rmbNum[num]; ok {
				if val == 89 {
					for _, r := range rmbSeq {
						rmbNum[r] = 89
					}
					yes = true
					break
				} else if val == 1 {
					for _, r := range rmbSeq {
						rmbNum[r] = 1
					}
					break
				}
			}
		}
		if yes {
			only89 = append(only89, i)
		}
	}
	fmt.Println(len(only89))
}
