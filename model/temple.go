package model

import (
	"math/rand"
	"strconv"

	"github.com/JuanVF/StarWars/utils"
)

type Temple struct {
	owner         *Player
	price         int64
	componentType int64
	time          int64
	currentTime   int64
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
	t.currentTime = utils.GetCurrentTime()
	t.time = 60000 * 5 //5 minutos

	go t.DoAction()
}

func (t *Temple) Run() {

}

func (t *Temple) OnHit(player *Player) {
	t.owner.RemoveObject(t)
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

// Remueve una relacion
func (t *Temple) RemoveRelation(obj GameObject) {
	index := -1

	for i := 0; i < len(t.relations); i++ {
		if t.relations[i] == obj {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	t.relations = append(t.relations[:index], t.relations[:index+1]...)
}

// Funciones de la interfaz iFactory
func (t *Temple) DoAction() {
	for {
		if t.currentTime+t.time-utils.GetCurrentTime() < 0 {
			t.owner.HasShield += int64(rand.Intn(6)) + 6 // Generamos una cantidad de escudos

			t.owner.FactoryChan <- "Templo: ahora tienes: " + strconv.Itoa(int(t.owner.HasShield)) + " de escudos..."
			t.owner.FactoryChan <- "SERVER_STEEL"

			t.currentTime = utils.GetCurrentTime()
		}
	}
}

func (t *Temple) Stop() {

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
