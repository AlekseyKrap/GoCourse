package spiner

//Уберите из первого примера (Фибоначчи и спиннер) функцию,
// вычисляющую числа Фибоначчи.
// Как теперь заставить спиннер вращаться в течение некоторого времени?
// 10 секунд?

import (
	"fmt"
	"time"
)

func RunSpiner() {
	go spinner(10 * time.Millisecond)
	//const n = 42
	//fibN := fibonacci(n)
	//fmt.Printf("\rFibonacci(%d) = %d\n", 10, 1)
	//fmt.Scanln()

	//timer2 := time.NewTimer(time.Second * 10)
	//go func() {
	//	<-timer2.C
	//	fmt.Println("Timer 2 expired")
	//}()
	//
	//stop2 := timer2.Stop()
	//if stop2 {
	//	fmt.Println("Timer 2 stopped")
	//}

	fmt.Scanln()
}

func spinner(delay time.Duration) {

	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)

			time.Sleep(delay)
		}

		//break
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
