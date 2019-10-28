package fibonacci

import "fmt"

func Fibonacci() {
	var j int = 1
	var k int = 0
	for i := 0; i <= 100; i++ {
		j += k
		fmt.Println(j)
		j, k = k, j
	}

}
