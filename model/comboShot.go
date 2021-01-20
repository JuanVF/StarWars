package model

type ComboShot struct {
	price int64
}

// Funciones de la iGuns
func (c *ComboShot) Shot(player *Player, x int64, y int64) {

}

func (c *ComboShot) GetPrice() int64 {
	return c.price
}

func (c *ComboShot) SetPrice(price int64) {

}
