// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

// Функция SolutionUniqOne выполняет подсчёт количества единственных значений в массиве A
func SolutionUniqOne(A []int) int {
	var unique int = 0
	counts := make(map[int]int)
	for _, number := range A {
		counts[number]++
	}
	for _, count := range counts {
		if count == 1 {
			unique++
		}
	}
	return unique
}
