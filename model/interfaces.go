package model

import "github.com/JuanVF/StarWars/utils"

// Esto representa cualquier cosa en la matriz
type GameObject interface {
	OnStart()
	Run()
	OnHit(player *Player)
	GetSize() utils.Point
	GetType() int64
	AddRelation(obj GameObject)
	GetRelations() []GameObject
	RemoveRelation(obj GameObject)
}

// Estos son los componentes de juego
type Component interface {
	SetPrice(price int64)
	GetPrice() int64
	SetPlayer(player *Player)
	GetPlayer() *Player
	GetType() int64
	SetType(componentType int64)
}

// Estos son los componentes que pueden producir cosas
type Factory interface {
	DoAction()
	Stop()
}

// Esto representa las armas en el juego
type Guns interface {
	Shot(attacker, player *Player, pos []float64) string
	GetPrice() int64
	SetPrice(price int64)
}
