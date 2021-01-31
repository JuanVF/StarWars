package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type Armory struct {
	owner         *Player
	price         int64
	gunType       int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la interfaz iGameObject
func (a *Armory) OnStart() {
	a.SetPrice(1000)

	a.Size = utils.Point{
		Width:  1,
		Height: 2,
	}

	a.SetType(utils.ARMORY)
}

func (a *Armory) Run() {

}

func (a *Armory) OnHit(player *Player) {
	a.owner.RemoveObject(a)
}

func (a *Armory) GetSize() utils.Point {
	return a.Size
}

// Agrega una relacion con otro objeto de la matriz
func (a *Armory) AddRelation(obj GameObject) {
	a.relations = append(a.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (a *Armory) GetRelations() []GameObject {
	return a.relations
}

// Remueve una relacion
func (a *Armory) RemoveRelation(obj GameObject) {
	index := -1

	for i := 0; i < len(a.relations); i++ {
		if a.relations[i] == obj {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	a.relations = append(a.relations[:index], a.relations[:index+1]...)
}

// Funciones de la interfaz iFactory
func (a *Armory) DoAction() {
	switch a.componentType {
	case utils.BOMB:
	case utils.MISSILE:
	case utils.COMBOSHOT:
	case utils.MULTISHOT:
	default:
	}
}

func (a *Armory) Stop() {

}

// Funciones de la interfaz iComponent
func (a *Armory) SetPrice(typeC int64) {
	a.price = typeC
}

func (a *Armory) GetPrice() int64 {
	return a.price
}

func (a *Armory) SetPlayer(owner *Player) {

}

func (a *Armory) GetPlayer() *Player {
	return a.owner
}

func (a *Armory) GetType() int64 {
	return a.componentType
}

func (a *Armory) SetType(ComponentType int64) {
	a.componentType = ComponentType
}
