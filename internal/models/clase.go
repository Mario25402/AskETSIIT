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

func NewClase(dia DiaSemana.DiaSemana, periodo *Periodo, aula string, grupo Grupo) (*Clase, error) {
	return &Clase{
		DiaSemana: dia,
		Periodo:   periodo,
		Aula:      aula,
		Grupo:     grupo,
	}, nil
}
