package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type World struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la interfaz iGameObject
func (w *World) OnStart() {
	w.SetPrice(12000)

	w.Size = utils.Point{
		Width:  2,
		Height: 2,
	}

	w.SetType(utils.WORLD)
}

func (w *World) Run() {

}

func (w *World) OnHit(player *Player) {
	w.owner.RemoveObject(w)
}

func (w *World) GetSize() utils.Point {
	return w.Size
}

// Agrega una relacion con otro objeto de la matriz
func (w *World) AddRelation(obj GameObject) {
	w.relations = append(w.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (w *World) GetRelations() []GameObject {
	return w.relations
}

// Remueve una relacion
func (w *World) RemoveRelation(obj GameObject) {
	index := -1

	for i := 0; i < len(w.relations); i++ {
		if w.relations[i] == obj {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	w.relations = append(w.relations[:index], w.relations[:index+1]...)
}

// Funciones de la interfaz iComponent
func (w *World) SetPrice(int64) {

}

func (w *World) GetPrice() int64 {
	return w.price
}

func (w *World) SetPlayer(owner *Player) {

}

func (w *World) GetPlayer() *Player {
	return w.owner
}

func (w *World) GetType() int64 {
	return w.componentType
}

func (w *World) SetType(ComponentType int64) {
	w.componentType = ComponentType
}
