package models

type MinutosPosibles string

const (
	EnPunto MinutosPosibles = "En Punto"
	YMedia  MinutosPosibles = "Y Media"
)

const (
	HoraInicioDia = 0
	HoraFinDia    = 23
)

type HoraMinutos struct {
	Hora    int // 0-23
	Minutos MinutosPosibles
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}
