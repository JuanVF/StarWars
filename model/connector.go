package model

import (
	"github.com/JuanVF/StarWars/utils"
)

type Connector struct {
	owner         *Player
	price         int64
	componentType int64
	Size          utils.Point
	relations     []GameObject
}

// Funciones de la interfaz iGameObject
func (c *Connector) OnStart() {
	c.SetPrice(100)

	c.Size = utils.Point{
		Width:  1,
		Height: 1,
	}

	c.SetType(utils.CONNECTOR)
}

func (c *Connector) Run() {

}

func (c *Connector) OnHit(player *Player) {
	c.owner.RemoveObject(c)
}

func (c *Connector) GetSize() utils.Point {
	return c.Size
}

// Agrega una relacion con otro objeto de la matriz
func (c *Connector) AddRelation(obj GameObject) {
	c.relations = append(c.relations, obj)
}

// Retorna las relaciones que tiene un objeto con otro de la matriz
func (c *Connector) GetRelations() []GameObject {
	return c.relations
}

// Funciones de la interfaz iComponent
func (c *Connector) SetPrice(int64) {

}

func (c *Connector) GetPrice() int64 {
	return c.price
}

func (c *Connector) SetPlayer(owner *Player) {

}

func (c *Connector) GetPlayer() *Player {
	return c.owner
}

func (c *Connector) GetType() int64 {
	return c.componentType
}

func (c *Connector) SetType(ComponentType int64) {
	c.componentType = ComponentType
}
