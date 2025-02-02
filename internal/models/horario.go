package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
)

type Horario struct {
	Clases map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase
}

func (horario Horario) GetClaseDia(dia DiaSemana.DiaSemana) map[HoraMinutos]*Clase {
	return horario.Clases[dia]
}

func (horario Horario) GetClaseHora(dia DiaSemana.DiaSemana, hora string) string {
	time := newHoraMinutos(hora)
	clases := horario.GetClaseDia(dia)
	return clases[*time].Grupo.Asignatura
}

func (horario Horario) GetHoraAsignatura(asignatura string, dia DiaSemana.DiaSemana) *HoraMinutos {
	clases := horario.GetClaseDia(dia)

	for _, clase := range clases {
		if clase.Grupo.Asignatura == asignatura {
			return &clase.Periodo.HoraInicio
		}
	}

	return nil
}

func (horario Horario) GetProfesorAsignatura(asignatura, hora string, dia DiaSemana.DiaSemana) string {
	time := newHoraMinutos(hora)
	clases := horario.GetClaseDia(dia)

	for _, clase := range clases {
		if clase.Grupo.Asignatura == asignatura && clase.Periodo.HoraInicio == *time {
			return clase.Grupo.Profesor
		}
	}

	return ""
}

func (horario Horario) GetAulaAsignatura(asignatura, hora string, dia DiaSemana.DiaSemana) string {
	time := newHoraMinutos(hora)
	clases := horario.GetClaseDia(dia)

	for _, clase := range clases {
		if clase.Grupo.Asignatura == asignatura && clase.Periodo.HoraInicio == *time {
			return clase.Aula
		}
	}

	return ""
}
