package main

import (
	"fmt"
	"strconv"
)

func checkConcaPan(num int) bool {
	str := strconv.Itoa(num)
	for j := 1; j <= 4; j++ {
		num, _ := strconv.Atoi(str[0:j])
		pandigital := true
		for k, l := 2, j; ; k++ {
			nextNum := num * k
			nextNumStr := strconv.Itoa(nextNum)
			nextLastIdx := l + len(nextNumStr)
			if nextLastIdx > 9 {
				pandigital = false
				break
			}
			realStr := str[l:nextLastIdx]
			//fmt.Println(nextNumStr, realStr)
			if nextNumStr != realStr {
				pandigital = false
				break
			}
			l = nextLastIdx
			if l == 9 {
				break
			}
		}
		if pandigital {
			return true
		}
	}
	return false
}

var allPandi []int

func generatePandi(str []string, result []string) {
	if len(str) == 0 {
		concatStr := ""
		for _, num := range result {
			concatStr += num
		}
		concatInt, _ := strconv.Atoi(concatStr)
		allPandi = append(allPandi, concatInt)
		return
	}
	for i := 0; i < len(str); i++ {
		newResult := make([]string, len(result)+1)
		copy(newResult, result)
		newStr := make([]string, len(str))
		copy(newStr, str)
		newStr = append(newStr[0:i], newStr[i+1:]...)
		generatePandi(newStr, append(newResult, str[i]))
	}
}

func main() {
	//const num = 192384586
	generatePandi([]string{"9", "8", "7", "6", "5", "4", "3", "2", "1"}, []string{})
	for _, num := range allPandi {
		if checkConcaPan(num) {
			fmt.Println(num)
			break
		}
	}
}
