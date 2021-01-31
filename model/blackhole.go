package model

import (
	"strconv"

	"github.com/JuanVF/StarWars/utils"
)

type BlackHole struct {
	owner         *Player
	price         int64
	gunType       int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la interfaz iGameObject
func (a *BlackHole) OnStart() {
	a.SetPrice(1000)

	a.Size = utils.Point{
		Width:  1,
		Height: 2,
	}

	a.SetType(utils.ARMORY)
}

func (a *BlackHole) Run() {

}

func (a *BlackHole) OnHit(player *Player) {
	ocol := -1
	orow := -1

	for col := 0; col < len(a.owner.Matrix); col++ {
		for row := 0; row < len(a.owner.Matrix[col]); row++ {
			if a.owner.Matrix[col][row] == a {
				ocol = col
				orow = row
			}
		}
	}

	if ocol != -1 && orow != -1 {
		player.FactoryChan <- "Server: has golpeado un agujero negro"
		if player.Matrix[ocol][orow] != nil {
			player.FactoryChan <- "Devolviendo ataque en {" + strconv.Itoa(ocol) + "}, {" + strconv.Itoa(orow) + "}"

			player.Matrix[ocol][orow].OnHit(a.owner)
		}
	}
}

func (a *BlackHole) GetSize() utils.Point {
	return a.Size
}

// Agrega una relacion con otro objeto de la matriz
func (a *BlackHole) AddRelation(obj GameObject) {
	a.relations = append(a.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (a *BlackHole) GetRelations() []GameObject {
	return a.relations
}

// Remueve una relacion
func (a *BlackHole) RemoveRelation(obj GameObject) {
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
func (a *BlackHole) DoAction() {
	switch a.componentType {
	case utils.BOMB:
	case utils.MISSILE:
	case utils.COMBOSHOT:
	case utils.MULTISHOT:
	default:
	}
}

func (a *BlackHole) Stop() {

}

// Funciones de la interfaz iComponent
func (a *BlackHole) SetPrice(typeC int64) {
	a.price = typeC
}

func (a *BlackHole) GetPrice() int64 {
	return a.price
}

func (a *BlackHole) SetPlayer(owner *Player) {

}

func (a *BlackHole) GetPlayer() *Player {
	return a.owner
}

func (a *BlackHole) GetType() int64 {
	return a.componentType
}

func (a *BlackHole) SetType(ComponentType int64) {
	a.componentType = ComponentType
}
