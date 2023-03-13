package main

import (
	"fmt"
)

// type Color int

// type Car struct {
// 	Make  string
// 	Model string
// 	Color Color
// }

// func main() {
// 	// Example byte slice containing two objects of different types
// 	data := []byte{
// 		// First object is a Car with a Make of "Honda", Model of "Civic", and Red color
// 		0x01,                                // Object type (1 = Car)
// 		0x00, 0x06, 'H', 'o', 'n', 'd', 'a', // Make (length 6)
// 		0x00, 0x05, 'C', 'i', 'v', 'i', 'c', // Model (length 5)
// 		0x00, 0x00, 0x00, 0x00, // Color (0 = Red)

// 		// Second object is a Car with a Make of "Toyota", Model of "Camry", and Green color
// 		0x01,                                     // Object type (1 = Car)
// 		0x00, 0x06, 'T', 'o', 'y', 'o', 't', 'a', // Make (length 6)
// 		0x00, 0x05, 'C', 'a', 'm', 'r', 'y', // Model (length 5)
// 		0x00, 0x00, 0x00, 0x01, // Color (1 = Green)
// 	}

// 	// Parse the byte slice into objects of the correct types
// 	var objects []interface{}
// 	for len(data) > 0 {
// 		// Read the object type byte
// 		objType := data[0]
// 		data = data[1:]

// 		switch objType {
// 		case 0x01:
// 			// Car object
// 			var car Car
// 			car.Make, data = readString(data)
// 			car.Model, data = readString(data)
// 			car.Color = Color(binary.LittleEndian.Uint32(data))
// 			data = data[4:]
// 			objects = append(objects, &car)
// 		default:
// 			// Unknown object type
// 			fmt.Printf("Unknown object type: 0x%x\n", objType)
// 			return
// 		}
// 	}

// 	// Print out the parsed objects
// 	for _, obj := range objects {
// 		switch obj := obj.(type) {
// 		case *Car:
// 			fmt.Printf("Car: %s %s (%d)\n", obj.Make, obj.Model, obj.Color)
// 		}
// 	}
// }

// // Helper function to read a string from a byte slice
// func readString(data []byte) (string, []byte) {
// 	length := int(binary.LittleEndian.Uint16(data))
// 	return string(data[2 : 2+length]), data[2+length:]
// }

type MessageType byte

const (
	CarMessageType MessageType = iota + 1
	TruckMessageType
)

type Color int

type Car struct {
	Make  string
	Model string
	Color Color
}

type Truck struct {
	Make      string
	Model     string
	NumWheels int
}

type Message struct {
	Type MessageType
	Data []byte
}

func main() {
	// Example message byte slice containing two objects of different types
	msgBytes := []byte{
		// First object is a Car with a Make of "Honda", Model of "Civic", and Red color
		0x01,                                // Message type (1 = Car)
		0x00, 0x06, 'H', 'o', 'n', 'd', 'a', // Make (length 6)
		0x00, 0x05, 'C', 'i', 'v', 'i', 'c', // Model (length 5)
		0x00, 0x00, 0x00, 0x00, // Color (0 = Red)

		// Second object is a Truck with a Make of "Ford", Model of "F-150", and 4 wheels
		0x02,                           // Message type (2 = Truck)
		0x00, 0x04, 'F', 'o', 'r', 'd', // Make (length 4)
		0x00, 0x05, 'F', '-', '1', '5', '0', // Model (length 5)
		0x00, 0x00, 0x00, 0x04, // NumWheels (4)
	}

	// Parse the message into objects of the correct types
	var cars []*Car
	var trucks []*Truck
	var messages []Message

	for len(msgBytes) > 0 {
		msgType := MessageType(msgBytes[0])
		msgBytes = msgBytes[1:]

		switch msgType {
		case CarMessageType:
			// Car object
			var car Car
			car.Make, msgBytes = readString(msgBytes)
			car.Model, msgBytes = readString(msgBytes)
			car.Color = Color(msgBytes[0])
			msgBytes = msgBytes[1:]
			cars = append(cars, &car)
		case TruckMessageType:
			// Truck object
			var truck Truck
			truck.Make, msgBytes = readString(msgBytes)
			truck.Model, msgBytes = readString(msgBytes)
			truck.NumWheels = int(msgBytes[0])
			msgBytes = msgBytes[1:]
			trucks = append(trucks, &truck)
		default:
			// Unknown message type
			fmt.Printf("Unknown message type: 0x%x\n", msgType)
			return
		}

		messages = append(messages, Message{
			Type: msgType,
			Data: msgBytes,
		})
	}

	// Print out the parsed objects
	fmt.Println("Cars:")
	for _, car := range cars {
		fmt.Printf("- %s %s (%d)\n", car.Make, car.Model, car.Color)
	}

	fmt.Println("Trucks:")
	for _, truck := range trucks {
		fmt.Printf("- %s %s (%d wheels)\n", truck.Make, truck.Model, truck.NumWheels)
	}

}
