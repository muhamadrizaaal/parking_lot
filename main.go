package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	registrationNumber string
	color              string
}

type ParkingLot struct {
	slots []*Car
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		slots: make([]*Car, capacity),
	}
}

func (p *ParkingLot) Park(car *Car) (int, error) {
	for i, slot := range p.slots {
		if slot == nil {
			p.slots[i] = car
			return i + 1, nil
		}
	}
	return 0, fmt.Errorf("Sorry, parking lot is full")
}

func (p *ParkingLot) Leave(slotNumber int, hours int) (*Car, int, error) {
	if slotNumber < 1 || slotNumber > len(p.slots) {
		return nil, 0, fmt.Errorf("Slot number %d is out of range", slotNumber)
	}
	slotIndex := slotNumber - 1
	car := p.slots[slotIndex]
	if car == nil {
		return nil, 0, fmt.Errorf("Slot number %d is already empty", slotNumber)
	}
	p.slots[slotIndex] = nil
	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}
	return car, charge, nil
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No.\tRegistration No.")
	for i, car := range p.slots {
		if car != nil {
			fmt.Printf("%d\t%s\n", i+1, car.registrationNumber)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input file")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var parkingLot *ParkingLot

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		command := tokens[0]

		switch command {
		case "create_parking_lot":
			if len(tokens) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			capacity, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Invalid capacity")
				continue
			}
			parkingLot = NewParkingLot(capacity)
			fmt.Printf("Created parking lot with %d slots\n", capacity)

		case "park":
			if len(tokens) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			car := &Car{registrationNumber: tokens[1], color: tokens[2]}
			slotNumber, err := parkingLot.Park(car)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Allocated slot number: %d\n", slotNumber)
			}

		case "leave":
			if len(tokens) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			slotNumber, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Invalid slot number")
				continue
			}
			hours, err := strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println("Invalid hours")
				continue
			}
			car, charge, err := parkingLot.Leave(slotNumber, hours)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Registration number %s with Slot Number %d is free with Charge %d\n",
					car.registrationNumber, slotNumber, charge)
			}

		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			parkingLot.Status()

		default:
			fmt.Println("Unknown command:", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
