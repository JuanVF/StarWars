package controller

import (
	"fmt"

	"github.com/JuanVF/StarWars/utils"
	"github.com/gorilla/websocket"
)

// Asignamos el admin
func AssignAdmin(client *websocket.Conn) {
	fmt.Printf("Servidor: Asignando administrador...\n")

	pack := NetworkPackage{
		ID: utils.FIRST_USER,
	}

	SendTo(CreatePackage(&pack, client))
}

// Creamos un nuevo jugador
func CreatePlayer(pack *NetworkPackage) {
	fmt.Println("Servidor: Creando nuevo perfil de jugador...")

}
