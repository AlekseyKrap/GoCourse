package anyInterface

import "fmt"

//Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

type AnyInterface interface {
	Sum() float64
}

type Struct1 struct {
	A float64
	D float64
}
type Struct2 struct {
	A float64
	B float64
}

func (s Struct1) Sum() float64 {
	return s.A + s.D
}
func (s Struct2) Sum() float64 {
	return s.A + s.B
}

func totalSum(a ...AnyInterface) (sum float64) {
	for _, o := range a {
		sum += o.Sum()
	}
	return
}

func Init() {
	var s1 = Struct1{D: 1, A: 2}
	var s2 = Struct2{B: 3, A: 2}

	fmt.Println("doubleTotalSum", totalSum(s1, s2, s1, s2))

}
