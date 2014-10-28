package main

import "fmt"

func checkPandig(allNum []int) bool {
	checker := make([]bool, 10)
	for _, num := range allNum {
		tmpNum := num
		for tmpNum != 0 {
			remainder := tmpNum % 10
			if checker[remainder] {
				return false
			} else {
				checker[remainder] = true
			}
			tmpNum = tmpNum / 10
		}
	}
	if checker[0] {
		return false
	}
	for i := 1; i <= 9; i++ {
		if !checker[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPandig([]int{39, 186, 7254}))
	all := []int{}
	for i := 1; i < 10; i++ {
		for j := 1000; j < 10000; j++ {
			sum := j * i
			if sum > 10000 {
				continue
			} else if checkPandig([]int{i, j, sum}) {
				isIn := false
				for _, n := range all {
					if n == sum {
						isIn = true
						break
					}
				}
				if !isIn {
					all = append(all, sum)
				}
			}
		}
	}
	for i := 11; i < 100; i++ {
		for j := 100; j < 1000; j++ {
			sum := j * i
			if sum > 10000 {
				continue
			} else if checkPandig([]int{i, j, sum}) {
				isIn := false
				fmt.Println(i, j, sum)
				for _, n := range all {
					if n == sum {
						isIn = true
						break
					}
				}
				if !isIn {
					all = append(all, sum)
				}
			}
		}
	}
	fmt.Println(all)
	total := 0
	for _, n := range all {
		total += n
	}
	fmt.Println(total)
}
