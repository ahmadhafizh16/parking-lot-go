package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandProcessor struct {
	ParkingLot *ParkingLot
}

func (cp *CommandProcessor) ProcessCommands(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		if len(command) == 0 {
			continue
		}
		cp.handleCommand(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func (cp *CommandProcessor) handleCommand(command []string) {
	switch command[0] {
	case "create_parking_lot":
		cp.handleCreateParkingLot(command)
	case "park":
		cp.handlePark(command)
	case "leave":
		cp.handleLeave(command)
	case "status":
		cp.handleStatus()
	default:
		fmt.Println("Invalid command:", command[0])
	}
}

func (cp *CommandProcessor) handleCreateParkingLot(command []string) {
	if len(command) != 2 {
		fmt.Println("Invalid command format for create_parking_lot")
		return
	}

	capacity, err := strconv.Atoi(command[1])
	if err != nil {
		fmt.Println("Invalid capacity value:", command[1])
		return
	}

	cp.ParkingLot = NewParkingLot(capacity)
	fmt.Printf("Created a parking lot with %d slots, charge per hour: $%d\n", capacity, cp.ParkingLot.BasicCharges)
}

func (cp *CommandProcessor) handlePark(command []string) {
	if len(command) != 3 || cp.ParkingLot == nil {
		fmt.Println("Invalid command or parking lot not initialized")
		return
	}

	car := Car{RegistrationNumber: command[1], Color: command[2]}
	slot, err := cp.ParkingLot.Park(car)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Allocated slot number: %d\n", slot)
}

func (cp *CommandProcessor) handleLeave(command []string) {
	if len(command) != 3 || cp.ParkingLot == nil {
		fmt.Println("Invalid command or parking lot not initialized")
		return
	}

	carNumber := command[1]
	hoursParked, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Invalid hours value:", command[2])
		return
	}

	totalCharges, slotIndex, err := cp.ParkingLot.Leave(carNumber, hoursParked)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Registration number %s with Slot number %d is free. Total charge: $%d\n", carNumber, slotIndex, totalCharges)
}

func (cp *CommandProcessor) handleStatus() {
	if cp.ParkingLot != nil {
		cp.ParkingLot.Status()
	} else {
		fmt.Println("Parking lot not initialized")
	}
}
