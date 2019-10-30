package main

import (
	"./addressBook"
	"./cars"
	"./fifo"
)

func main() {
	cars.InitCars()
	fifo.ExecFifo()
	addressBook.ExecAddressBook()
}
