package fibonacci

import "fmt"

func Fibonacci() {
	var i int = 0
	var j int = 1
	var sum int = 0
	var k int = 0
	for i <= 100 {
		sum = j + k
		fmt.Println(sum)
		j = k
		k = sum
		i++
	}

}
