package evenNumber

import (
	"fmt"
	"strconv"
)

func IsEvenNumber(str string, divider int64) (res bool, error error) {
	i, error := strconv.ParseInt(str, 10, 64)
	if error != nil {
		return
	}
	if (i % divider) == 0 {
		res = true
	}
	return
}
func EvenNumber() {
	var str string
	fmt.Println("Введите  число:")
	fmt.Scanln(&str)
	isEven, error := IsEvenNumber(str, 2)
	if error != nil {
		fmt.Println("что-то пошло не так")
		return
	}
	if isEven {
		fmt.Println("Чиcло четное")
	} else {
		fmt.Println("Чиcло не четное")
	}

}

func EvenNumber3() {
	var str string
	fmt.Println("Введите  число:")
	fmt.Scanln(&str)
	isEven, error := IsEvenNumber(str, 3)
	if error != nil {
		fmt.Println("что-то пошло не так")
		return
	}
	if isEven {
		fmt.Println("Чиcло делится на 3")
	} else {
		fmt.Println("Чиcло не делится на 3")
	}

}
