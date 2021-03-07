package chipsequence

import (
	"strconv"
	"strings"

	"github.com/aaronschweig/cdma-calculator/util"
)

func CalculateWithRisingValue(chipSequences [][8]int) [][8]int {
	sequences := make([][8]int, 4)
	copy(sequences, chipSequences)

	A := sequences[0]

	for i := 0; i < 256; i++ {
		repr := strconv.FormatInt(int64(i), 2)
		for {
			repr = "0" + repr
			if len(repr) > 8 {
				repr = repr[:8]
				break
			}
		}
		bitArr := strings.Split(repr, "")
		var seq []int
		for _, bit := range bitArr {
			parsed, _ := strconv.ParseInt(bit, 2, 8)
			if int(parsed) == 0 {
				seq = append(seq, -1)
			} else {
				seq = append(seq, 1)
			}
		}

		var conv [8]int
		copy(conv[:], seq)

		if util.DotProduct(A, conv) == 0 {
			if util.CheckPairwiseOrthogonal(append(sequences, conv)) {
				sequences = append(sequences, conv)
			}
		}
	}

	return sequences
}
