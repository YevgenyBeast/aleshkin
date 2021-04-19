// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

// Функция SolutionSquareGenerator выполняет генерацию квадратов n натуральных чисел, начиная со start
func SolutionSquareGenerator(start int, n int) []int {
	var square []int
	for i := start; i < start+n; i++ {
		square = append(square, i*i)
	}
	return square
}
