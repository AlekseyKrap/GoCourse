package primeNumbers

import "fmt"

func isprime(n int) bool {
	if n == 1 {
		return false
	}
	if n%n != 0 {
		return false
	}

	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func PrimeNumbers() {

	var x []int

	for i := 2; len(x) <= 100; i++ {
		if isprime(i) {
			x = append(x, i)
		}
	}

	fmt.Println(x)

}
