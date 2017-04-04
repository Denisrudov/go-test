package main

import (
	"names/api"
	"fmt"
)

func main() {

	service := api.NewService()

	randomName, err := service.GetName()
	chukJoke, err := service.GetNorris()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(randomName.Name)
	fmt.Println(chukJoke.Value.Joke)

}
