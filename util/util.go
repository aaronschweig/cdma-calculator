package util

import (
	"fmt"
	"strconv"
	"strings"
)

func StringRepresentation(v [8]int) string {
	var bits []string
	for _, element := range v {
		if element == -1 {
			bits = append(bits, "0")
		} else {
			bits = append(bits, "1")
		}
	}
	return strings.Join(bits, "")
}

func ConvertToDecimal(chipSequences [][8]int) {
	for _, v := range chipSequences {
		representation := StringRepresentation(v)
		decimal, err := strconv.ParseInt(representation, 2, 64)
		if err != nil {
			panic(err)
		}
		fmt.Println(decimal)
	}
}

// 1/m vernachlässigbar, da 0/m = 0
func DotProduct(s1, s2 [8]int) (sum int) {
	for i, v := range s1 {
		sum += s2[i] * v
	}
	return sum
}

func CheckPairwiseOrthogonal(chipSequences [][8]int) bool {
	for i, s := range chipSequences {

		if i+1 > len(chipSequences) {
			continue
		}

		for j := i + 1; j < len(chipSequences); j++ {
			if DotProduct(s, chipSequences[j]) != 0 {
				return false
			}
		}

	}
	return true
}
