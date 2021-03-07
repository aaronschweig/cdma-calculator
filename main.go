package main

import (
	"fmt"

	"github.com/aaronschweig/cdma-calculator/chipsequence"
	"github.com/aaronschweig/cdma-calculator/util"
)

type Sequence = [8]int

func main() {
	A := Sequence{-1, -1, -1, 1, 1, -1, 1, 1}
	B := Sequence{-1, -1, 1, -1, 1, 1, 1, -1}
	C := Sequence{-1, 1, -1, 1, 1, 1, -1, -1}
	D := Sequence{-1, 1, -1, -1, -1, -1, 1, -1}

	chipSequences := []Sequence{A, B, C, D}

	fmt.Println("Permutation:\n---------------------------")
	resultWithPermutation := chipsequence.CalculateWithPermutations(chipSequences)

	util.ConvertToDecimal(resultWithPermutation)

	fmt.Println("\nRising Value:\n---------------------------")

	resultWithRisingValue := chipsequence.CalculateWithRisingValue(chipSequences)

	util.ConvertToDecimal(resultWithRisingValue)

	fmt.Println("\nFalling Value:\n---------------------------")

	resultWithFallingValue := chipsequence.CalculateWithFallingValue(chipSequences)

	util.ConvertToDecimal(resultWithFallingValue)

}
