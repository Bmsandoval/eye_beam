package main


var people map[int]Person

func init() {
	people = map[int]Person{
		1: {ID: 1, Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}},
		2: {ID: 2, Firstname: "Maria", Lastname: "Raboy"},
	}
}