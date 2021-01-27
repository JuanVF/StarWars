package model

import "github.com/JuanVF/StarWars/utils"

type Armory struct {
	owner         *Player
	price         int64
	gunType       int64
	componentType int64
}

// Funciones de la interfaz iGameObject
func (a *Armory) OnStart() {

}

func (a *Armory) Run() {

}

func (a *Armory) OnHit(player *Player) {

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

// Funciones de la interfaz iComponent
func (a *Armory) SetPrice(int64) {

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
}
