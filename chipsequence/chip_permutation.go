package chipsequence

import (
	"strings"

	"github.com/aaronschweig/cdma-calculator/util"
)

type Sequence = [8]int

// Heap-Algorithm
func permutations(arr Sequence) (res [][]int) {
	var helper func([8]int, int)

	helper = func(arr [8]int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp[:], arr[:])
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func generateSequencePermutation() (perms []Sequence) {
	// Starten mit 255_2 => 1 1 1 1 1 1 1 1
	s := Sequence{1, 1, 1, 1, 1, 1, 1, 1}

	// Erste Permutation muss nicht "permutiert" werden, da immer 255
	perms = append(perms, s)

	// Berechnen der verbleibenden Permutationen zwischen 1 - 254
	for i := 0; i < len(s)-1; i++ {
		s[i] = -1

		perms = append(perms, s)

		check := make(map[string]int)

		// Nur einzigartige Permutationen werden berücksichtigt
		for _, v := range permutations(s) {
			var permInt8 [8]int
			copy(permInt8[:], v)
			check[util.StringRepresentation(permInt8)] = 1
		}

		for unique := range check {
			sequence := make([]int, 8)
			uniqueParts := strings.Split(unique, "")
			for i, part := range uniqueParts {
				if part == "0" {
					sequence[i] = -1
				} else {
					sequence[i] = 1
				}
			}
			var ps Sequence
			copy(ps[:], sequence)
			perms = append(perms, ps)
		}

	}

	perms = append(perms, Sequence{-1, -1, -1, -1, -1, -1, -1, -1})

	return perms
}

func CalculateWithPermutations(chipSequences [][8]int) []Sequence {
	sequences := make([][8]int, 4)
	copy(sequences, chipSequences)

	A := sequences[0]

	perms := generateSequencePermutation()

	for _, v := range perms {
		if util.DotProduct(A, v) == 0 {
			if util.CheckPairwiseOrthogonal(append(sequences, v)) {
				sequences = append(sequences, v)
			}
		}
	}

	return sequences
}
