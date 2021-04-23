// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

// Функция SolutionUniq выполняет подсчёт количества уникальных значений в массиве A
func SolutionUniq(A []int) int {
	counts := make(map[int]int)
	for _, number := range A {
		counts[number]++
	}
	return len(counts)
}
