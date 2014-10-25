package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(bytes), "\n")
	savedSum := 0
	for _, l := range lines {
		if cal(l) != cal(shortform(l)) {
			fmt.Println(cal(l), l, cal(shortform(l)), shortform(l))
		}
		//fmt.Println(l, cal(l), shortform(l))
		savedSum += len(l) - len(shortform(l))
	}
	fmt.Println(savedSum)
}

func shortform(str string) string {
	sum := cal(str)
	M := sum / 1000
	sum = sum % 1000
	var newform string
	for i := 0; i < M; i++ {
		newform += "M"
	}
	DC := sum / 100
	if DC == 9 {
		newform += "CM"
	} else if DC == 4 {
		newform += "CD"
	} else {
		D := DC / 5
		for i := 0; i < D; i++ {
			newform += "D"
		}
		C := DC % 5
		for i := 0; i < C; i++ {
			newform += "C"
		}
	}
	sum = sum % 100
	LX := sum / 10
	if LX == 9 {
		newform += "XC"
	} else if LX == 4 {
		newform += "XL"
	} else {
		L := LX / 5
		for i := 0; i < L; i++ {
			newform += "L"
		}
		X := LX % 5
		for i := 0; i < X; i++ {
			newform += "X"
		}
	}
	sum = sum % 10
	VI := sum
	if VI == 9 {
		newform += "IX"
	} else if VI == 4 {
		newform += "IV"
	} else {
		V := VI / 5
		for i := 0; i < V; i++ {
			newform += "V"
		}
		I := VI % 5
		for i := 0; i < I; i++ {
			newform += "I"
		}
	}
	return newform
}

func cal(str string) int {
	sum := 0
	var i int
	for ; i < len(str)-1; i++ {
		switch string(str[i]) {
		case "M":
			sum += 1000
		case "D":
			sum += 500
		case "C":
			if string(str[i+1]) == "M" {
				i++
				sum += 900
			} else if string(str[i+1]) == "D" {
				i++
				sum += 400
			} else {
				sum += 100
			}
		case "L":
			sum += 50
		case "X":
			if string(str[i+1]) == "C" {
				i++
				sum += 90
			} else if string(str[i+1]) == "L" {
				i++
				sum += 40
			} else {
				sum += 10
			}
		case "V":
			sum += 5
		case "I":
			if string(str[i+1]) == "X" {
				i++
				sum += 9
			} else if string(str[i+1]) == "V" {
				i++
				sum += 4
			} else {
				sum += 1
			}
		}
	}
	if i == len(str)-1 {
		switch string(str[len(str)-1]) {
		case "M":
			sum += 1000
		case "D":
			sum += 500
		case "C":
			sum += 100
		case "L":
			sum += 50
		case "X":
			sum += 10
		case "V":
			sum += 5
		case "I":
			sum += 1
		}
	}
	return sum
}
