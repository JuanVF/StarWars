package model

import "github.com/JuanVF/StarWars/utils"

type Connector struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
}

// Funciones de la interfaz iGameObject
func (c Connector) OnStart() {
	c.SetPrice(100)

	c.Size = utils.Point{
		Width:  1,
		Height: 1,
	}

	c.SetType(utils.CONNECTOR)
}

func (c Connector) Run() {

}

func (c Connector) OnHit(player *Player) {
	c.owner.RemoveObject(c)
}

// Funciones de la interfaz iComponent
func (c Connector) SetPrice(int64) {

}

func (c Connector) GetPrice() int64 {
	return c.price
}

func (c Connector) SetPlayer(owner *Player) {

}

func (c Connector) GetPlayer() *Player {
	return c.owner
}

func (c Connector) GetType() int64 {
	return c.componentType
}

func (c Connector) SetType(ComponentType int64) {

}
