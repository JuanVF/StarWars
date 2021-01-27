package controller

import (
	"fmt"
	"log"

	"github.com/JuanVF/StarWars/utils"
	"github.com/gorilla/websocket"
)

// Aqui definimos la manera de enviar paquetes a las personas
type NetworkPackage struct {
	Msg      *Message
	ID       int
	To       *websocket.Conn
	Response bool
}

// Aqui dependiendo el tipo de paquete que se reciba determinamos cual es la respuesta
// del servidor al siguiente mensaje
func CreatePackage(msg *NetworkPackage, client *websocket.Conn) *NetworkPackage {
	pack := NetworkPackage{}

	switch msg.ID {
	case utils.FIRST_USER:
		pack.Msg = &Message{
			IdMessage: "ADMIN",
		}
		pack.To = client
		pack.Response = true
	case utils.LOGIN:
		pack.Msg = &Message{
			IdMessage: "REQUESTDATA",
		}
		pack.To = client
		pack.Response = true
	case utils.CREATE_PLAYER:
		pack.Msg = &Message{
			IdMessage: "STARTED",
		}
		pack.To = nil
		pack.Response = true

		CreatePlayer(msg)
	default:
		return nil
	}

	return &pack
}

// Dado una respuesta del usuario determino cual es el paquete que se usara despues
func CreatePackageMsg(msg *Message, client *websocket.Conn) *NetworkPackage {
	pack := NetworkPackage{}

	switch msg.IdMessage {
	case "ADMIN":
		pack.ID = utils.LOGIN
		pack.To = client
		pack.Response = true
	case "REQUESTDATA":
		pack.ID = utils.CREATE_PLAYER
		pack.To = nil
		pack.Response = true
	default:
		return nil
	}

	return &pack
}

// Enviamos un paquete a una persona x
func SendTo(pack *NetworkPackage) {
	err := pack.To.WriteJSON(pack.Msg)

	if err != nil {
		log.Printf("Message error: %v\n", err)

		pack.To.Close()
		delete(Clients, pack.To)
	}

	fmt.Printf("Sending to: %v, message: %v\n", pack.Msg.Name, pack.Msg.IdMessage)
}

// Enviamos paquetes a todo el mundo
func SendToAll(msg *Message) {
	for client := range Clients {
		err := client.WriteJSON(msg)

		if err != nil {
			log.Printf("Message error: %v\n", err)

			client.Close()
			delete(Clients, client)
		}

		fmt.Printf("Sending to: %v, message: %v\n", msg.Name, msg.IdMessage)
	}
}
