// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

import (
	"log"
	"math"
)

// Функция SolutionSquareGenerator выполняет генерацию квадратов n натуральных чисел, начиная со start
func SolutionSquareGenerator(start int, n int) []int {
	var square []int
	max := int(math.Sqrt(math.MaxInt64))
	if start <= 0 {
		log.Fatal("Заданное число не является натуральным")
	}
	if n < 1 {
		log.Fatal("Количество выводимых чисел не должно быть меньше одного")
	}
	if start > (max - n + 1) {
		log.Fatal("Введённый диапазон чисел выходит за границы int")
	}
	for i := start; i < start+n; i++ {
		square = append(square, i*i)
	}
	return square
}
