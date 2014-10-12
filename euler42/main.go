package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	in, _ := ioutil.ReadAll(os.Stdin)
	allStr := strings.Split(string(in), "\"")
	// get rid of ,
	words := make([]string, 0, len(allStr))
	for _, content := range allStr {
		if content != "," && content != "" {
			words = append(words, content)
		}
	}
	// calculate value of each word
	// also find max value to calculate triangle number
	values := make([]int, len(words))
	max := 0
	for i, w := range words {
		sum := 0
		for _, c := range w {
			sum += int(c) - 64
		}
		if sum > max {
			max = sum
		}
		values[i] = sum
	}
	fmt.Println(values)
	// calculate traingle number and stop if more than max
	triNum := []float64{}
	i := 1
	calNum := (0.5) * float64(i) * float64(i+1)
	for calNum < float64(max) {
		triNum = append(triNum, calNum)
		i++
		calNum = (0.5) * float64(i) * float64(i+1)
	}
	fmt.Println(triNum)
	// loop check for triangle word
	count := 0
	for _, v := range values {
		for _, num := range triNum {
			if float64(v) == num {
				count++
			}
		}
	}
	fmt.Println(count)
}
