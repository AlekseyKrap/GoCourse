package cost

import "fmt"

func ConvertСurrency() {
	const costDollar float32 = 63.99
	fmt.Println("Введите сумму в рублях:")
	var costRub float32
	fmt.Scanln(&costRub)

	fmt.Println("Итого: ", costRub/costDollar)
}
