package model

import "github.com/JuanVF/StarWars/utils"

type World struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
}

// Funciones de la interfaz iGameObject
func (w World) OnStart() {
	w.SetPrice(12000)

	w.Size = utils.Point{
		Width:  2,
		Height: 2,
	}

	w.SetType(utils.WORLD)
}

func (w World) Run() {

}

func (w World) OnHit(player *Player) {
	w.owner.RemoveObject(w)
}

// Funciones de la interfaz iComponent
func (w World) SetPrice(int64) {

}

func (w World) GetPrice() int64 {
	return w.price
}

func (w World) SetPlayer(owner *Player) {

}

func (w World) GetPlayer() *Player {
	return w.owner
}

func (w World) GetType() int64 {
	return w.componentType
}

func (w World) SetType(ComponentType int64) {

}
