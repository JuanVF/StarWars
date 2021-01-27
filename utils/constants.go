package utils

// Enumerables para los componentes
const (
	WORLD     = iota
	CONNECTOR = iota
	MARKET    = iota
	MINE      = iota
	ARMORY    = iota
	TEMPLE    = iota
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
)
