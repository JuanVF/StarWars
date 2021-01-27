package model

// Esto representa cualquier cosa en la matriz
type GameObject interface {
	OnStart()
	Run()
	OnHit(player *Player)
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
}

// Esto representa las armas en el juego
type Guns interface {
	Shot(player *Player, x int64, y int64)
	GetPrice() int64
	SetPrice(price int64)
}