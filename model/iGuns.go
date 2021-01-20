package model

type Guns interface {
	Shot(player *Player, x int64, y int64)
	GetPrice() int64
	SetPrice(price int64)
}
