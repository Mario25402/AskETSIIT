package models

import (
	"errors"
	"strconv"
	"strings"
)

type MinutosPosibles string

const (
	EnPunto MinutosPosibles = "En Punto"
	YMedia  MinutosPosibles = "Y Media"
)

const (
	HoraInicioDia = 0
	HoraFinDia    = 23
)

type HoraMinutos struct {
	Hora    int // 0-23
	Minutos MinutosPosibles
}

func newHoraMinutos(tiempo string) (*HoraMinutos, error) {
	partes := strings.Split(tiempo, ":")

	if len(partes) != 2 {
		return nil, errors.New("la hora debe tener el formato HH:MM")
	}

	horas, err := strconv.Atoi(partes[0])
	if err != nil {
		return nil, err
	}

	if horas < HoraInicioDia || horas > HoraFinDia {
		return nil, errors.New("la hora debe estar entre 0 y 23")
	}

	var minutos MinutosPosibles

	if partes[1] == "00" {
		minutos = EnPunto
	} else if partes[1] == "30" {
		minutos = YMedia
	} else {
		return nil, errors.New("los minutos deben ser 00 o 30")
	}

	return &HoraMinutos{Hora: horas, Minutos: minutos}, nil
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}

func newPeriodo(tiempoInicio, tiempoFinal string) (*Periodo, error) {
	ini, err := newHoraMinutos(tiempoInicio)
	if err != nil {
		return nil, err
	}

	fin, err := newHoraMinutos(tiempoFinal)
	if err != nil {
		return nil, err
	}

	if ini.Hora > fin.Hora || (ini.Hora == fin.Hora && ini.Minutos > fin.Minutos) {
		return nil, errors.New("la hora de inicio debe ser anterior a la hora de fin")
	}

	return &Periodo{HoraInicio: *ini, HoraFin: *fin}, nil
}
