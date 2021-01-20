package model

type GameObject interface {
	OnStart()
	Run()
	OnHit(player *Player)
}
