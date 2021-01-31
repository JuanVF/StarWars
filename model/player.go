package model

import (
	"fmt"
	"strconv"

	"github.com/JuanVF/StarWars/utils"
)

type Player struct {
	Name           string
	Money          int64
	Steel          int64
	HasShield      int64
	IsGraphVisible bool
	GunsList       map[Guns]int
	Matrix         [15][15]GameObject
	Graph          map[utils.Point]map[utils.Point]float64
	FactoryChan    chan string
}

func (p *Player) AddObject(object GameObject) {

}

// Removemos un componente de la matriz y hacemos el grafo visible
func (p *Player) RemoveObject(object GameObject) {
	cantWorlds := 0

	if object.GetType() == utils.CONNECTOR {
		for _, rel := range object.GetRelations() {
			fmt.Println("Rel tipo: " + utils.ComponentIDToString(int(rel.GetType())))
			if rel.GetType() != utils.WORLD {
				p.IsGraphVisible = true
			}
		}
	}

	for col := 0; col < 15; col++ {
		for row := 0; row < 15; row++ {
			if p.Matrix[col][row] == nil {
				continue
			}

			if p.Matrix[col][row] == object {
				p.Matrix[col][row] = nil

				continue
			}

			// Si es un mundo aumento la cantidad de mundos
			if p.Matrix[col][row].GetType() == utils.WORLD {
				cantWorlds++
			}
		}
	}

	if !p.IsGraphVisible {
		p.IsGraphVisible = cantWorlds == 0
	}

	if p.IsGraphVisible {
		p.RemoveRelations(object)

		if object.GetType() == utils.WORLD {
			p.Money += 10000
		}

		p.FactoryChan <- "SERVER_GRAFO_VISIBLE"
	}
}

// Remueve las relaciones de un objeto
func (p *Player) RemoveRelations(object GameObject) {
	for i := 0; i < len(object.GetRelations()); i++ {
		object.GetRelations()[i].RemoveRelation(object)
	}
}

// Le agregamos acero a un jugador
func (p *Player) AddSteel(amount int64) {
	p.Steel += amount

	p.FactoryChan <- "Mina: Has recibido: " + strconv.Itoa(int(amount)) + " de acero..."
	p.FactoryChan <- "SERVER_STEEL"
}

// Generamos el grafo inicial de juego
func (p *Player) GenerateGraph() {
	p.Graph = make(map[utils.Point]map[utils.Point]float64)

	worldList := make(map[GameObject]utils.Point)
	connectorList := make(map[GameObject]utils.Point)
	componentList := make(map[GameObject]utils.Point)

	for col := 0; col < len(p.Matrix); col++ {
		for row := 0; row < len(p.Matrix[col]); row++ {
			if p.Matrix[col][row] == nil {
				continue
			}

			object := p.Matrix[col][row]
			point := utils.Point{
				X: float64(col),
				Y: float64(row),
			}

			switch object.GetType() {
			case utils.WORLD:
				worldList[object] = point
			case utils.CONNECTOR:
				connectorList[object] = point
			case utils.BLACK_HOLE:
				continue
			default:
				componentList[object] = point
			}
		}
	}

	// Generamos las relaciones de los componentes
	fmt.Println("Generando conexiones")
	p.ConnectWorld(worldList, connectorList)
	p.ConnectComponents(componentList, connectorList)
}

// Conectamos el mundo a todos los conectores
func (p *Player) ConnectWorld(worlds, connectors map[GameObject]utils.Point) {
	for connector := range connectors {
		for world := range worlds {
			// Conectamos en ambos sentidos
			fmt.Printf("Se conecta el mundo: %v con el conector %v\n", worlds[world], connectors[connector])

			weight := worlds[world].GetDistance(connectors[connector])

			// Evitamos que los hashmap no exista
			if p.Graph[worlds[world]] == nil {
				p.Graph[worlds[world]] = make(map[utils.Point]float64)
			}

			if p.Graph[connectors[connector]] == nil {
				p.Graph[connectors[connector]] = make(map[utils.Point]float64)
			}

			p.Graph[worlds[world]][connectors[connector]] = weight
			p.Graph[connectors[connector]][worlds[world]] = weight

			world.AddRelation(connector)
			connector.AddRelation(world)
		}
	}
}

// Conectamos los componentes con el conector mas cercano
func (p *Player) ConnectComponents(components, connectors map[GameObject]utils.Point) {
	for component := range components {
		var leastDistConn GameObject = nil
		var weight float64 = 0

		// Determinamos el conector que esta mas cerca y ese sera el elegido para la relacion
		for connector := range connectors {
			tmpWeight := connectors[connector].GetDistance(components[component])

			if leastDistConn == nil || tmpWeight < weight {
				leastDistConn = connector
				weight = tmpWeight
			}
		}

		// Evitamos nulos
		if leastDistConn == nil {
			continue
		}

		// Conectamos en ambos sentidos
		fmt.Printf("Se conecta el conector: %v con el componente %v\n", components[component], connectors[leastDistConn])

		// Evitamos que los hashmap no exista
		if p.Graph[components[component]] == nil {
			p.Graph[components[component]] = make(map[utils.Point]float64)
		}

		if p.Graph[connectors[leastDistConn]] == nil {
			p.Graph[connectors[leastDistConn]] = make(map[utils.Point]float64)
		}

		p.Graph[components[component]][connectors[leastDistConn]] = weight
		p.Graph[connectors[leastDistConn]][components[component]] = weight

		component.AddRelation(leastDistConn)
		leastDistConn.AddRelation(component)
	}
}

// Devuelve la matriz de un jugador
func (p *Player) GetMatrix() [][]float64 {
	var matrix [][]float64

	for i := 0; i < len(p.Matrix); i++ {
		matrix = append(matrix, []float64{})

		for j := 0; j < len(p.Matrix[i]); j++ {
			var val float64 = -1

			if p.Matrix[i][j] != nil {
				val = float64(p.Matrix[i][j].GetType())
			}

			matrix[i] = append(matrix[i], float64(val))
		}
	}

	return matrix
}

// Aqui vamos a retornar los puntos que el cliente necesita para dibujar las lineas
// Formato: {col0, fila0, col1, fila1, ..., xn, xn+1, xn+2, x+3}
func (p *Player) GetGraphPoints() []float64 {
	cache := make(map[GameObject][]float64)
	points := []float64{}

	for col := 0; col < len(p.Matrix); col++ {
		for row := 0; row < len(p.Matrix[col]); row++ {
			var object GameObject = p.Matrix[col][row]
			pos := []float64{float64(col), float64(row)}

			// Nos saltamos los espacios vacios
			if object == nil {
				continue
			}

			// Si esta en cache lo obtenemos de ahi, sino lo guardamos en cache
			// Esto para hacer busquedas mas rapidas
			if cache[object] != nil {
				pos = cache[object]
			} else {
				cache[object] = pos
			}

			// Agregamos los puntos
			relations := object.GetRelations()

			for i := 0; i < len(relations); i++ {
				tmpPos := []float64{}

				if relations[i] == nil {
					continue
				}

				// Obtenemos la posicion de un objeto en la matriz
				if cache[relations[i]] != nil {
					tmpPos = cache[relations[i]]
				} else {
					tmpPos = p.GetByPosition(relations[i])
					cache[relations[i]] = tmpPos
				}

				// Agregamos las posiciones
				points = append(points, pos...)
				points = append(points, tmpPos...)
			}
		}
	}

	fmt.Println("Puntos para el front:")
	fmt.Println(points)

	return points
}

// Dado un componente obtenemos sus puntos con respecot a sus relaciones
//func (p *Player) GetComponentPoints(col, row float64, object GameObject) []float64 {
//}

// Dado un jugador retorna en que posicion de la matriz esta
// Formato: {col0, fila0}
func (p *Player) GetByPosition(obj GameObject) []float64 {
	for col := 0; col < len(p.Matrix); col++ {
		for row := 0; row < len(p.Matrix[col]); row++ {
			if p.Matrix[col][row] == obj {
				return []float64{float64(col), float64(row)}
			}
		}
	}

	return nil
}

// Aqui generamos la matriz
func (p *Player) GenerateMatrix(Matrix [][]float64) {
	for col := 0; col < len(Matrix); col++ {
		for row := 0; row < len(Matrix[col]); row++ {
			if p.Matrix[col][row] != nil {
				continue
			}

			var object GameObject = nil

			switch Matrix[col][row] {
			case utils.WORLD:
				object = &World{
					owner: p,
				}
			case utils.ARMORY:
				object = &Armory{
					owner: p,
				}
			case utils.MARKET:
				object = &Market{
					owner: p,
				}
			case utils.CONNECTOR:
				object = &Connector{
					owner: p,
				}
			case utils.TEMPLE:
				object = &Temple{
					owner: p,
				}
			case utils.MINE:
				object = &Mine{
					owner: p,
				}
			default:
				continue
			}

			if object == nil {
				continue
			}

			p.Matrix[col][row] = object
			p.Matrix[col][row].OnStart()
			size := p.Matrix[col][row].GetSize()

			p.AssignMatrixBySize(col, row, size)
		}
	}
}

// Se asignan los campos extras con la misma referencia de memoria
func (p *Player) AssignMatrixBySize(col, row int, size utils.Point) {
	for i := col; i < col+int(size.X); i++ {
		for j := row; j < row+int(size.Y); j++ {
			p.Matrix[i][j] = p.Matrix[col][row]
		}
	}
}

// Funcion de debug solo printeamos la matriz
func (p *Player) PrintMatrix() {
	for i := 0; i < len(p.Matrix); i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(p.Matrix[i]); j++ {
			if p.Matrix[i][j] == nil {
				fmt.Printf("[%d] ", -1)
			} else {
				fmt.Printf("[%d] ", p.Matrix[i][j].GetType())
			}
		}
		fmt.Printf("]\n")
	}
}
