package model

import (
	"strconv"

	"github.com/JuanVF/StarWars/utils"
)

type ComboShot struct {
	price int64
}

// Funciones de la iGuns
func (c *ComboShot) Shot(attacker, player *Player, pos []float64) string {
	attackPoints := (utils.Point{}).Parse(pos)
	log := ""

	// Atacamos los puntos
	for _, point := range attackPoints {
		col := int(point.X)
		row := int(point.Y)
		obj := player.Matrix[col][row]

		// El ataque fallo
		if obj == nil {
			log += "El ataque en: {" + strconv.Itoa(col) + ", " + strconv.Itoa(row) + " } fallo...\n"
			continue
		}

		// El ataque no fallo
		log += "El ataque en: {" + strconv.Itoa(col) + ", " + strconv.Itoa(row) + " } acerto...\n"
		log += "Componente: " + utils.ComponentIDToString(int(obj.GetType()))

		obj.OnHit(attacker)
	}

	return log
}

func (c *ComboShot) GetPrice() int64 {
	return 5000
}

func (c *ComboShot) SetPrice(price int64) {
	c.price = 5000
}
