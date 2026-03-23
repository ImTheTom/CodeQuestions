package main

import (
	"fmt"
	"time"
)

const (
	unspecified = iota
	big
	medium
	small
)

type ParkingSystem struct {
	BigSpots    int
	MediumSpots int
	SmallSpots  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		BigSpots:    big,
		MediumSpots: medium,
		SmallSpots:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case big:
		if this.BigSpots > 0 {
			this.BigSpots--
			return true
		}
	case medium:
		if this.MediumSpots > 0 {
			this.MediumSpots--
			return true
		}
	case small:
		if this.SmallSpots > 0 {
			this.SmallSpots--
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	system := Constructor(1, 1, 0)
	fmt.Println(system.AddCar(1))
	fmt.Println(system.AddCar(2))
	fmt.Println(system.AddCar(3))
	fmt.Println(system.AddCar(1))

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
