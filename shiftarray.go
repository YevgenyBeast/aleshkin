// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

import "log"

// Функция SolutionRotate выполняет сдвиг массива A вправо k раз
func SolutionRotate(A []int, k int) []int {
	n := len(A)
	k = k % n
	if k < 0 {
		log.Fatal("Введено отрицательное количество сдвигов")
	}
	if k == 0 {
		return A
	}
	shiftA := make([]int, n)
	for i := 0; i < n; i++ {
		if (i - k) >= 0 {
			shiftA[i] = A[i-k]
		} else {
			shiftA[i] = A[i-k+n]
		}
	}
	return shiftA
}
