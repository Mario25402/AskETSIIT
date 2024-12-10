package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"errors"
)

type Clase struct {
	DiaSemana DiaSemana.DiaSemana
	Periodo   *Periodo
	Aula      string
	Grupo     Grupo
}

func NewClase(dia DiaSemana.DiaSemana, periodo *Periodo, aula string, grupo Grupo) (*Clase, error) {

	if dia < DiaSemana.Lunes || dia > DiaSemana.Viernes {
		return nil, errors.New("dia debe estar entre Lunes (0) y Viernes (5)")
	}

	return &Clase{
		DiaSemana: dia,
		Periodo:   periodo,
		Aula:      aula,
		Grupo:     grupo,
	}, nil
}
