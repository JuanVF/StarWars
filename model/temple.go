package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type Temple struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la interfaz iGameObject
func (t *Temple) OnStart() {
	t.SetPrice(1000)

	t.Size = utils.Point{
		Width:  1,
		Height: 2,
	}

	t.SetType(utils.ARMORY)
}

func (t *Temple) Run() {

}

func (t *Temple) OnHit(player *Player) {

}

func (t *Temple) GetSize() utils.Point {
	return t.Size
}

// Agrega una relacion con otro objeto de la matriz
func (t *Temple) AddRelation(obj GameObject) {
	t.relations = append(t.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (t *Temple) GetRelations() []GameObject {
	return t.relations
}

// Funciones de la interfaz iFactory
func (t *Temple) DoAction() {

}

// Funciones de la interfaz iComponent
func (t *Temple) SetPrice(int64) {

}

func (t *Temple) GetPrice() int64 {
	return t.price
}

func (t *Temple) SetPlayer(owner *Player) {

}

func (t *Temple) GetPlayer() *Player {
	return t.owner
}

func (t *Temple) GetType() int64 {
	return t.componentType
}

func (t *Temple) SetType(ComponentType int64) {
	t.componentType = ComponentType
}
