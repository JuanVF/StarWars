package model

type Bomb struct {
	price int64
}

// Funciones de la iGuns
func (b *Bomb) Shot(player *Player, x int64, y int64) {

}

func (b *Bomb) GetPrice() int64 {
	return b.price
}

func (b *Bomb) SetPrice(price int64) {

}
