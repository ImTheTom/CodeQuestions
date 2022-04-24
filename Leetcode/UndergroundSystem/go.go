package main

import "fmt"

type UndergroundSystem struct {
	CheckIns  []*UndergroundSystemCheckIn
	CheckOuts []*UndergroundSystemCheckOut
}

type UndergroundSystemCheckIn struct {
	ID          int
	StationName string
	Time        int
}

type UndergroundSystemCheckOut struct {
	ID              int
	CheckInStation  string
	CheckOutStation string
	Time            int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		CheckIns:  []*UndergroundSystemCheckIn{},
		CheckOuts: []*UndergroundSystemCheckOut{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.CheckIns = append(this.CheckIns, &UndergroundSystemCheckIn{
		ID:          id,
		StationName: stationName,
		Time:        t,
	})
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	var checkInLoc string
	var checkInTime int
	for i, checkIn := range this.CheckIns {
		if checkIn != nil && checkIn.ID == id {
			checkInLoc = checkIn.StationName
			checkInTime = checkIn.Time
			this.CheckIns[i] = nil
			break
		}
	}
	this.CheckOuts = append(this.CheckOuts, &UndergroundSystemCheckOut{
		ID:              id,
		CheckInStation:  checkInLoc,
		CheckOutStation: stationName,
		Time:            t - checkInTime,
	})
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	totalTime := 0.0
	numberOfTimes := 0
	for _, checkOut := range this.CheckOuts {
		if checkOut.CheckOutStation == endStation && checkOut.CheckInStation == startStation {
			totalTime += float64(checkOut.Time)
			numberOfTimes++
		}
	}
	return totalTime / float64(numberOfTimes)
}

func main() {
	fmt.Println("Running main")
	id1 := 27
	station1, station2 := "Waterloo", "Paradise"
	t1, t2 := 8, 20
	obj := Constructor()
	obj.CheckIn(id1, station1, t1)
	obj.CheckOut(id1, station2, t2)
	param_3 := obj.GetAverageTime(station1, station2)
	fmt.Println(param_3)
}
