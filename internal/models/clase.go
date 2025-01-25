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

func newClase(dia DiaSemana.DiaSemana, periodo *Periodo, aula string, grupo Grupo) (*Clase, error) {
	if dia == "" || periodo == nil || aula == "" || grupo.Asignatura == "" {
		return nil, errors.New("no se puede crear la clase")
	}

	return &Clase{
		DiaSemana: dia,
		Periodo:   periodo,
		Aula:      aula,
		Grupo:     grupo,
	}, nil
}
