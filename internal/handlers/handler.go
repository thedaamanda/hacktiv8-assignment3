package handlers

import (
	"assignment3/internal/services"
	"fmt"
)

func GetDataStatus() {
	data, err := services.ReadJSON(services.GetFiles())

	if err != nil {
		panic(err)
	}

	fmt.Println("=========================================")
	fmt.Println("Updating monitoring...")
	fmt.Printf("Current Water Value	: %d m\n", data.Status.Water)
	fmt.Printf("Current Wind Value	: %d m/s\n", data.Status.Wind)
	fmt.Println("=========================================")

	services.Randomize(&data)

	fmt.Printf("Updated Water Value	: %d m\n", data.Status.Water)
	fmt.Printf("Updated Wind Value	: %d m/s\n", data.Status.Wind)
	fmt.Println("Update data success!")

	err = services.WriteJSON(data, services.GetFiles())

	if err != nil {
		panic(err)
	}
}
