package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
)

type Horario struct {
	Clases map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase
}
