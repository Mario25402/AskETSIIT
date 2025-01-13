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

func strMinutos(mins string) (MinutosPosibles, error) {
	if mins == "00" {
		return EnPunto, nil
	} else if mins == "30" {
		return YMedia, nil
	} else {
		return "", errors.New("los minutos deben ser 00 o 30")
	}
}

func minutosStr(mins MinutosPosibles) string {
	if mins == EnPunto {
		return "00"
	} else if mins == YMedia {
		return "30"
	} else {
		return ""
	}
}

type HoraMinutos struct {
	Hora    int // 0-23
	Minutos MinutosPosibles
}

func GetHoraMinutosStr(hm *HoraMinutos) string {
	return strconv.Itoa(hm.Hora) + ":" + minutosStr(hm.Minutos)
}

func newHoraMinutos(horas int, mins MinutosPosibles) (*HoraMinutos, error) {
	if horas < 0 || horas > 23 {
		return nil, errors.New("la hora debe estar entre 0 y 23")
	}
	return &HoraMinutos{Hora: horas, Minutos: mins}, nil
}

func NewHoraMinutosStr(tiempo string) (*HoraMinutos, error) {
	partes := strings.Split(tiempo, ":")
	if len(partes) != 2 {
		return nil, errors.New("la hora debe tener el formato HH:MM")
	}

	horas, err := strconv.Atoi(partes[0])
	if err != nil {
		return nil, err
	}

	return NewHoraMinutosSplit(horas, partes[1])
}

func NewHoraMinutosSplit(horas int, mins string) (*HoraMinutos, error) {
	minutos, err := strMinutos(mins)
	if err != nil {
		return nil, err
	}

	return newHoraMinutos(horas, minutos)
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}

func newPeriodo(horaInicio, horaFinal HoraMinutos) (*Periodo, error) {
	if horaInicio.Hora > horaFinal.Hora || (horaInicio.Hora == horaFinal.Hora && horaInicio.Minutos > horaFinal.Minutos) {
		return nil, errors.New("la hora de inicio debe ser anterior a la hora de fin")
	}

	return &Periodo{HoraInicio: horaInicio, HoraFin: horaFinal}, nil
}

func NewPeriodoStr(tiempoInicio, tiempoFinal string) (*Periodo, error) {
	ini, err := NewHoraMinutosStr(tiempoInicio)
	if err != nil {
		return nil, err
	}

	fin, err := NewHoraMinutosStr(tiempoFinal)
	if err != nil {
		return nil, err
	}

	return newPeriodo(*ini, *fin)
}

func NewPeriodoSplit(horasIni int, minsIni string, horasFin int, minsFin string) (*Periodo, error) {
	ini, err := NewHoraMinutosSplit(horasIni, minsIni)
	if err != nil {
		return nil, err
	}

	fin, err := NewHoraMinutosSplit(horasFin, minsFin)
	if err != nil {
		return nil, err
	}

	return newPeriodo(*ini, *fin)
}
