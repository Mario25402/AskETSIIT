package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
)

type Clase struct {
	DiaSemana DiaSemana.DiaSemana
	Periodo   *Periodo
	Aula      string
	Grupo     Grupo
}
