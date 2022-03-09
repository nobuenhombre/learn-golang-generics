package main

import (
	"fmt"
)

// GenericNumber
// описывает интерфейс для разных типов
type GenericNumber interface {
	int64 | float64 | int8
}

// sumNumbers
// при помощи GENERICS описывается универсальная функция
// которая может оперировать разными типами
func sumNumbers[N GenericNumber](s []N) N {
	var total N

	for _, v := range s {
		total += v
	}

	return total
}

func main() {
	ints := []int64{32, 64, 96, 128}
	floats := []float64{32.31, 64.51, 97.21, 123.81}
	bytes := []int8{8, 12, 16, 24, 32}

	sumI := sumNumbers[int64](ints)
	sumF := sumNumbers[float64](floats)
	sumB := sumNumbers[int8](bytes)

	fmt.Printf("%v (%T) \n", sumI, sumI)
	fmt.Printf("%.2f (%T) \n", sumF, sumF)
	fmt.Printf("%v (%T) \n", sumB, sumB)
}
