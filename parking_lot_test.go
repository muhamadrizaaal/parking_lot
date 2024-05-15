package main

import (
	"testing"
)

func TestParkingLot(t *testing.T) {
	parkingLot := NewParkingLot(6)
	if len(parkingLot.slots) != 6 {
		t.Errorf("Expected 6 slots, got %d", len(parkingLot.slots))
	}

	car := &Car{registrationNumber: "KA-01-HH-1234", color: "White"}
	slot, err := parkingLot.Park(car)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if slot != 1 {
		t.Errorf("Expected slot 1, got %d", slot)
	}

	car, charge, err := parkingLot.Leave(1, 4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if charge != 30 {
		t.Errorf("Expected charge 30, got %d", charge)
	}
}
