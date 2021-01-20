package main

import (
	"fmt"

	"github.com/JuanVF/StarWars/model"
)

func main() {
	//controller.StartServer()

	player := model.Player{}
	mina := model.Mine{}

	mina.SetPlayer(&player)
	mina.OnStart()

	var prev int64 = 0

	for {
		mina.Run()

		if prev != player.Steel {
			fmt.Println("El jugador ahora tiene: ", player.Steel)
			prev = player.Steel
		}
	}
}
