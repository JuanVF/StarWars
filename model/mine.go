package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type Mine struct {
	owner         *Player
	Size          utils.Point
	price         int64
	amount        int64
	time          int64
	currentTime   int64
	componentType int64
	relations     []GameObject
}

// Funciones de la mina

// Permite configurar la cantidad de hierro que se produce en x segundos
func (m *Mine) Configure(time, amount int64) {
	m.time = time * 1000
	m.amount = amount
}

// Funciones de la interfaz iGameObject
func (m *Mine) OnStart() {
	m.SetPrice(1000)

	m.Size = utils.Point{
		Width:  1,
		Height: 2,
	}

	m.SetType(utils.MINE)

	m.currentTime = utils.GetCurrentTime()

	//m.time = 60000
	//m.amount = 50
	m.time = 2000
	m.amount = 500

	go m.DoAction()
}

// Cada x segundos produce una n cantidad de hierro
func (m *Mine) Run() {
}

func (m *Mine) OnHit(player *Player) {
	m.owner.RemoveObject(m)
}

func (m *Mine) GetSize() utils.Point {
	return m.Size
}

// Agrega una relacion con otro objeto de la matriz
func (m *Mine) AddRelation(obj GameObject) {
	m.relations = append(m.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (m *Mine) GetRelations() []GameObject {
	return m.relations
}

// Remueve una relacion
func (m *Mine) RemoveRelation(obj GameObject) {
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
func (m *Mine) DoAction() {
	for {
		if m.currentTime+m.time-utils.GetCurrentTime() < 0 {
			m.owner.AddSteel(m.amount)

			m.currentTime = utils.GetCurrentTime()
		}
	}
}

func (m *Mine) Stop() {

}

// Funciones de la interfaz iComponent
func (m *Mine) SetPrice(int64) {

}

func (m *Mine) GetPrice() int64 {
	return m.price
}

func (m *Mine) SetPlayer(owner *Player) {
	m.owner = owner
}

func (m *Mine) GetPlayer() *Player {
	return m.owner
}

func (m *Mine) GetType() int64 {
	return m.componentType
}

func (m *Mine) SetType(ComponentType int64) {
	m.componentType = ComponentType
}
