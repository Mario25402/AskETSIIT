package models

import (
	"strconv"
	"strings"
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

func newHoraMinutos(tiempo string) *HoraMinutos {
	partes := strings.Split(tiempo, ":")
	horas, _ := strconv.Atoi(partes[0])

	var minutos MinutosPosibles
	if partes[1] == "00" {
		minutos = EnPunto
	} else {
		minutos = YMedia
	}

	return &HoraMinutos{Hora: horas, Minutos: minutos}
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}
