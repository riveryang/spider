package utils

import (
	"strings"
	"math"
	"github.com/pkg/errors"
)

func Similarity(source, target string) (float64, error) {
	if len(strings.Trim(source, " ")) > 0 && len(strings.Trim(target, " ")) > 0 {
		algorithmMap := make(map[int32][]float64)
		for _, ch := range source {
			charIndex := int32(rune(ch))
			fq := algorithmMap[charIndex]
			if fq != nil && len(fq) == 2 {
				fq[0]++
			} else {
				fq = make([]float64, 2)
				fq[0] = 1
				fq[1] = 0
				algorithmMap[charIndex] = fq
			}
		}

		for _, ch := range target {
			charIndex := int32(rune(ch))
			fq := algorithmMap[charIndex]
			if fq != nil && len(fq) == 2 {
				fq[1]++
			} else {
				fq = make([]float64, 2)
				fq[0] = 1
				fq[1] = 0
				algorithmMap[charIndex] = fq
			}
		}

		var sq1 float64
		var sq2 float64
		var denominator float64
		for _, fq := range algorithmMap {
			denominator += fq[0] * fq[1]
			sq1 += fq[0] * fq[0]
			sq2 += fq[1] * fq[1]
		}

		return denominator / math.Sqrt(sq1*sq2), nil
	}

	return 0, errors.New("The Document is null or have not cahrs!!")
}