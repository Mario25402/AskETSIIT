package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"errors"
)

type Horario struct {
	Clases map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase
}

// NewHorarioFromClases crea un horario a partir de un conjunto de clases.
func NewHorarioFromClases(clases []Clase) (*Horario, error) {
	horario := &Horario{
		Clases: make(map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase),
	}

	for _, clase := range clases {
		if clase.Periodo == nil {
			return nil, errors.New("cada clase debe tener un periodo definido")
		}

		if _, ok := horario.Clases[clase.DiaSemana]; !ok {
			horario.Clases[clase.DiaSemana] = make(map[HoraMinutos]*Clase)
		}

		if _, conflict := horario.Clases[clase.DiaSemana][clase.Periodo.HoraInicio]; conflict {
			return nil, errors.New("conflicto de horario: dos clases no pueden empezar el mismo día a la misma hora")
		}

		horario.Clases[clase.DiaSemana][clase.Periodo.HoraInicio] = &clase
	}

	return horario, nil
}
