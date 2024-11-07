# Parking Lot Management System

This project implements a parking lot management system where users can park and leave cars in a parking lot, check the status of the parking lot, and view parking charges based on the time a car is parked. 

The test assignment mentions that the input should include the car color, but the provided `commands.txt` file does not include color input. To align with the requirements, I have added the car color to the `commands.txt` file.


## Features
- Create a parking lot with a given number of slots.
- Park cars in available slots and view the slot number.
- Leave a parking spot and calculate parking charges.
- View the current status of the parking lot with parked cars.
- Commands are processed from a file to simulate operations on the parking lot.

## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Commands Format](#commands-format)
4. [Example Commands](#example-commands)

## Installation

To run the Parking Lot Management system, follow these steps:

### 1. Install Go
Ensure that you have Go installed on your system. You can download and install Go from [here](https://go.dev/dl/).

### 2. Clone the repository
Clone the repository to your local machine:

```bash
git clone https://github.com/ahmadhafizh16/parking-lot-go.git
cd parking-lot-go
```

### 3. Build the application
To compile the application, run the following command:

```bash
go build -o application main.go
```

This will create an executable file named `application`.

### 4. Run the application
To run the system, you need to provide a command file as input. For example:

```bash
./application commands.txt
```

## Usage

The parking lot system can process commands from a file (`commands.txt`). The commands can be used to simulate the creation of the parking lot, parking of cars, leaving parking spots, and viewing the parking lot status.

Each command is processed one by one, and the appropriate action is performed.

## Commands Format

The commands are formatted in a plain text file (`commands.txt`) with each command on a new line. Here is the general format for each command:

1. **`create_parking_lot <capacity>`**: Creates a parking lot with the specified number of slots.
2. **`park <registration_number> <color>`**: Parks a car in the nearest available slot.
3. **`leave <registration_number> <hours_parked>`**: Frees a parking slot and calculates the total parking charge.
4. **`status`**: Prints the current status of the parking lot, including slot numbers, registration numbers, and car colors.

## Example Commands

Here is an example of a `commands.txt` file:

```txt
create_parking_lot 6
park KA-01-HH-1234 Red
park KA-01-HH-9999 Blue
park KA-01-BB-0001 Black
status
leave KA-01-HH-1234 2
leave KA-01-HH-9999 3
status
```

### Explanation of the Commands:

- `create_parking_lot 6`: Creates a parking lot with 6 available slots.
- `park KA-01-HH-1234 Red`: Parks a car with registration number `KA-01-HH-1234` and color `Red` in the nearest available slot.
- `status`: Displays the current status of the parking lot.
- `leave KA-01-HH-1234 2`: Frees the parking slot occupied by car `KA-01-HH-1234`, and calculates the parking charge for 2 hours.
- `leave KA-01-HH-9999 3`: Frees the parking slot occupied by car `KA-01-HH-9999`, and calculates the parking charge for 3 hours.

### Example Output:

```txt
Created a parking lot with 6 slots, charge per hour: $10
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Slot No. | Registration No. | Color
1        | KA-01-HH-1234    | Red
2        | KA-01-HH-9999    | Blue
3        | KA-01-BB-0001    | Black
Registration number KA-01-HH-1234 with Slot number 1 is free. Total charge: $10
Registration number KA-01-HH-9999 with Slot number 2 is free. Total charge: $20
Slot No. | Registration No. | Color
3        | KA-01-BB-0001    | Black
```

