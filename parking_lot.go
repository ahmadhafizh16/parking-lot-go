package main

import "fmt"

type ParkingLot struct {
	Capacity     int
	Slots        []Car
	BasicCharges int
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:     capacity,
		Slots:        make([]Car, capacity),
		BasicCharges: 10,
	}
}

// Park assigns the nearest available slot to the car.
func (pl *ParkingLot) Park(car Car) (int, error) {
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].RegistrationNumber == "" {
			pl.Slots[i] = car
			return i + 1, nil // return the allocated slot number (1-indexed)
		}
	}
	return -1, fmt.Errorf("Sorry, Parking lot is full")
}

// Leave frees up the parking slot and calculates the parking charge.
func (pl *ParkingLot) Leave(carNumber string, hoursParked int) (int, int, error) {
	slotIndex := -1
	for i, car := range pl.Slots {
		if car.RegistrationNumber == carNumber {
			slotIndex = i
			break
		}
	}

	if slotIndex == -1 {
		return -1, -1, fmt.Errorf("Registration number %s not found", carNumber)
	}

	pl.Slots[slotIndex] = Car{} // Free the slot

	// Calculate charges: $10 for the first 2 hours, $10 for each additional hour.
	totalCharges := pl.BasicCharges
	if hoursParked > 2 {
		totalCharges += (hoursParked - 2) * pl.BasicCharges
	}

	return totalCharges, (slotIndex + 1), nil
}

// Status prints the current status of the parking lot.
func (pl *ParkingLot) Status() {
	fmt.Println("Slot No. | Registration No. | Color")
	for i, car := range pl.Slots {
		if car.RegistrationNumber != "" {
			fmt.Printf("%-8d | %-16s | %s\n", i+1, car.RegistrationNumber, car.Color)
		}
	}
}
