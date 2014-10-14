package main

import (
	"fmt"
	"math/big"
)

func cal(n1, n2, n3 *big.Int) (numerator *big.Int, divisor *big.Int) {
	n1.Mul(n1, n3)
	numerator = n1.Add(n1, n2)
	divisor = n3
	return
}

func calSeq(seq []*big.Int) (*big.Int, *big.Int) {
	numerator, divisor := cal(seq[len(seq)-2], big.NewInt(1), seq[len(seq)-1])
	for i := len(seq) - 3; i >= 0; i-- {
		numerator, divisor = divisor, numerator
		numerator, divisor = cal(seq[i], numerator, divisor)
	}

	return numerator, divisor
}

func createESeq(num int) []*big.Int {
	seq := []*big.Int{big.NewInt(2)}
	k := int64(1)
	for len(seq) < num {
		seq = append(seq, big.NewInt(1), big.NewInt(k*2), big.NewInt(1))
		k++
	}
	seq = seq[:num]
	return seq
}
func main() {
	eSeq := createESeq(100)
	numerator, divisor := calSeq(eSeq)
	fmt.Println(numerator, divisor)
	text, _ := numerator.MarshalText()
	sum := 0
	for _, c := range text {
		sum += int(c) - 48
	}
	fmt.Println(sum)
}
