package utils

// Enumerables para los componentes
const (
	ARMORY     = iota
	CONNECTOR  = iota
	MARKET     = iota
	MINE       = iota
	WORLD      = iota
	TEMPLE     = iota
	BLACK_HOLE = iota
)

// Enumerables para los disparos
const (
	MISSILE   = iota
	MULTISHOT = iota
	BOMB      = iota
	COMBOSHOT = iota
)

// Enumerables para la distribucion de paquetes
const (
	FIRST_USER    = iota
	LOGIN         = iota
	CREATE_PLAYER = iota
	SEND_MATRIX   = iota
	SEND_GRAPH    = iota
	REQUEST_TURN  = iota
	CHAT          = iota
	ENEMY_MATRIX  = iota
	ENEMY_GRAPH   = iota
	ENEMY_INIT    = iota
	PLAYER_INIT   = iota
	ATTACK        = iota
	BUY_ARMORY    = iota
	SKIP          = iota
	NEXT_ENEMY    = iota
)

func ComponentIDToString(cType int) string {
	switch cType {
	case CONNECTOR:
		return "conector"
	case MINE:
		return "mina"
	case MARKET:
		return "mercado"
	case TEMPLE:
		return "templo"
	case ARMORY:
		return "armeria"
	case WORLD:
		return "mundo"
	case BLACK_HOLE:
		return "agujero negro"
	}

	return ""
}
