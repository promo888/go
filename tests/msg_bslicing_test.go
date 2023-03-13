package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

const (
	Red = iota
	Green
	Blue
)

type Color int

var Colors = [...]Color{Red, Green, Blue}

type Car struct {
	Make  string
	Model string
	Color Color
}

// Helper function to read a string from a byte slice
func readString(data []byte) (string, []byte) {
	length := int(binary.LittleEndian.Uint16(data))
	return string(data[2 : 2+length]), data[2+length:]
}

func TestReadObjects(t *testing.T) {
	//myCar := Car("Honda", "Civic", Colors[Red])

	// Example byte slice containing two objects of different types
	data := []byte{
		// First object is a Car with a Make of "Honda", Model of "Civic", and Red color
		0x01,                                // Object type (1 = Car)
		0x00, 0x06, 'H', 'o', 'n', 'd', 'a', // Make (length 6)
		0x00, 0x05, 'C', 'i', 'v', 'i', 'c', // Model (length 5)
		0x00, 0x00, 0x00, 0x00, // Color (0 = Red)

		// Second object is a Car with a Make of "Toyota", Model of "Camry", and Green color
		0x01,                                     // Object type (1 = Car)
		0x00, 0x06, 'T', 'o', 'y', 'o', 't', 'a', // Make (length 6)
		0x00, 0x05, 'C', 'a', 'm', 'r', 'y', // Model (length 5)
		0x00, 0x00, 0x00, 0x01, // Color (1 = Green)
	}

	// Expected result
	expected := []string{
		"Car: Honda Civic (0)",
		"Car: Toyota Camry (1)",
	}

	// Parse the byte slice
	var objects []interface{}
	for len(data) > 0 {
		objType := data[0]
		data = data[1:]

		switch objType {
		case 0x01:
			var car Car
			car.Make, data = readString(data)
			car.Model, data = readString(data)
			car.Color = Color(data[0])
			data = data[1:]
			objects = append(objects, &car)
		default:
			t.Errorf("Unknown object type: 0x%x", objType)
		}
	}

	// Compare the parsed objects to the expected result
	var result []string
	for _, obj := range objects {
		switch obj := obj.(type) {
		case *Car:
			result = append(result, fmt.Sprintf("Car: %s %s (%d)", obj.Make, obj.Model, obj.Color))
		default:
			t.Errorf("Unknown object type: %T", obj)
		}
	}

	if !bytes.EqualFold([]byte(expected[0]), []byte(result[0])) || !bytes.EqualFold([]byte(expected[1]), []byte(result[1])) {
		t.Errorf("Unexpected result, expected %v but got %v", expected, result)
	}
}
