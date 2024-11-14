package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
)

type Clase struct {
	Dia     DiaSemana.DiaSemana
	Periodo *Periodo
	Aula    string
	Grupo   Grupo
}
