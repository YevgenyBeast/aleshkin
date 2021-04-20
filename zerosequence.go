// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

import "math"

// Функция SolutionBinaryGap находит самую длинную последовательность нулей в двоичном представлении числа N
func SolutionBinaryGap(N int) int {
	var num uint = 0
	max := 0
	seqlen := 0
	if N == 0 {
		max = 1
	}
	if N > 0 {
		num = uint(N)
	}
	if N < 0 {
		num = math.MaxUint64 - uint(math.Abs(float64(N))) + 1
	}
	for num > 0 {
		if num%2 == 0 {
			seqlen++
			num = num / 2
		}
		if num%2 == 1 {
			if seqlen > max {
				max = seqlen
			}
			seqlen = 0
			num = num / 2
		}
	}
	return max
}
