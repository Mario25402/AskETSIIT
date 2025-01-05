package models

import (
	"errors"
)

type MinutosPosibles string

const (
	EnPunto MinutosPosibles = "En Punto"
	YMedia  MinutosPosibles = "Y Media"
)

type HoraMinutos struct {
	Hora    int // 0-23
	Minutos MinutosPosibles
}

func NewHoraMinutos(hora int, minutos MinutosPosibles) (*HoraMinutos, error) {
	if hora < 0 || hora > 23 {
		return nil, errors.New("la hora debe estar entre 0 y 23")
	}
	return &HoraMinutos{Hora: hora, Minutos: minutos}, nil
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}

func NewPeriodo(horaInicio, horaFin HoraMinutos) (*Periodo, error) {
	// Verificar que la hora de inicio es anterior a la hora de fin
	if horaInicio.Hora > horaFin.Hora || (horaInicio.Hora == horaFin.Hora && horaInicio.Minutos > horaFin.Minutos) {
		return nil, errors.New("la hora de inicio debe ser anterior a la hora de fin")
	}
	// Crear y devolver el Periodo si todo es válido
	periodo := &Periodo{
		HoraInicio: horaInicio,
		HoraFin:    horaFin,
	}

	return periodo, nil
}
