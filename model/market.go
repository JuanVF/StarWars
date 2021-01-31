package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type Market struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la estructura
func (m *Market) Sell(gunType int64) (int64, error) {
	return 0, nil
}

func (m *Market) Buy(gunType, amount int64) error {
	return nil
}

// Funciones de la interfaz iGameObject
func (m *Market) OnStart() {
	m.SetPrice(2000)

	m.Size = utils.Point{
		Width:  1,
		Height: 2,
	}

	m.SetType(utils.MARKET)
}

func (m *Market) Run() {

}

func (m *Market) OnHit(player *Player) {
	m.owner.RemoveObject(m)
}

func (m *Market) GetSize() utils.Point {
	return m.Size
}

// Agrega una relacion con otro objeto de la matriz
func (m *Market) AddRelation(obj GameObject) {
	m.relations = append(m.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (m *Market) GetRelations() []GameObject {
	return m.relations
}

// Remueve una relacion
func (m *Market) RemoveRelation(obj GameObject) {
	index := -1

	for i := 0; i < len(m.relations); i++ {
		if m.relations[i] == obj {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	m.relations = append(m.relations[:index], m.relations[:index+1]...)
}

// Funciones de la interfaz iFactory
func (m *Market) DoAction() {

}

func (m *Market) Stop() {

}

// Funciones de la interfaz iComponent
func (m *Market) SetPrice(int64) {

}

func (m *Market) GetPrice() int64 {
	return m.price
}

func (m *Market) SetPlayer(owner *Player) {

}

func (m *Market) GetPlayer() *Player {
	return m.owner
}

func (m *Market) GetType() int64 {
	return m.componentType
}

func (m *Market) SetType(ComponentType int64) {
	m.componentType = ComponentType
}
