package model

type Component interface {
	SetPrice(price int64)
	GetPrice() int64
	SetPlayer(player *Player)
	GetPlayer() *Player
	GetType() int64
	SetType(componentType int64)
}
