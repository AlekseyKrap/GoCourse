package cars

import "fmt"

type Cars struct {
	Mark         string
	Year         int
	IsRunning    bool
	isOpenWindow bool
}

type PassengerCar struct {
	Cars
	trunkVolume int
}
type Truck struct {
	Cars
	carcaseVolume int
}

func InitCars() {

	var anyPassengerCar = PassengerCar{Cars: Cars{
		Mark:         "B",
		Year:         2019,
		isOpenWindow: false,
		IsRunning:    true},
		trunkVolume: 10}

	var anyTruck = Truck{Cars: Cars{
		Mark:         "D",
		Year:         2020,
		isOpenWindow: true,
		IsRunning:    false},
		carcaseVolume: 100}

	fmt.Println("anyPassengerCar", anyPassengerCar)
	fmt.Println("anyTruck", anyTruck)

}
