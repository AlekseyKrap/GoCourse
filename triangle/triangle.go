package triangle

import (
	"fmt"
	"math"
)

func Triangle() {
	var catet1 float64
	var catet2 float64
	var s float64
	var c float64
	var p float64

	fmt.Println("\nВведите длинну катета1:")
	fmt.Scanln(&catet1)
	fmt.Println("\nВведите длинну катета2:")
	fmt.Scanln(&catet2)
	s = (catet1 * catet2) / 2
	fmt.Println("\nПлощадь треугольника: ", s)
	c = math.Sqrt(math.Pow(catet1, 2) + math.Pow(catet2, 2))
	fmt.Println("\nГипотенуза треугольника: ", c)
	p = catet1 + catet2 + c
	fmt.Println("\nПериметр треугольника: ", p)

}
