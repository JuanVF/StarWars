package model

import (
	"math/rand"
	"strconv"

	"github.com/JuanVF/StarWars/utils"
)

type MultiShot struct {
	price       int64
	alreadyShot bool
}

// Funciones de la iGuns
func (m *MultiShot) Shot(attacker, player *Player, pos []float64) string {
	attackPoints := (utils.Point{}).Parse(pos)
	log := ""

	hit := false

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

		hit = true

		// El ataque no fallo
		log += "El ataque en: {" + strconv.Itoa(col) + ", " + strconv.Itoa(row) + " } acerto...\n"
		log += "Componente: " + utils.ComponentIDToString(int(obj.GetType()))

		obj.OnHit(attacker)
	}

	if hit && !m.alreadyShot {
		log += "Acerto! Generando 4 disparos mas...\n"
		m.alreadyShot = true

		for i := 0; i < 4; i++ {
			log += m.Shot(attacker, player, []float64{float64(rand.Intn(15)), float64(rand.Intn(15))})
		}
	}

	return log
}

func (m *MultiShot) GetPrice() int64 {
	return 1000
}

func (m *MultiShot) SetPrice(price int64) {
	m.price = 1000
}
