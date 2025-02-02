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

func newClase(dia DiaSemana.DiaSemana, periodo *Periodo, aula string, grupo Grupo) *Clase {
	if dia == "" || periodo == nil || aula == "" || grupo.Asignatura == "" {
		return nil
	}

	return &Clase{
		DiaSemana: dia,
		Periodo:   periodo,
		Aula:      aula,
		Grupo:     grupo,
	}
}
