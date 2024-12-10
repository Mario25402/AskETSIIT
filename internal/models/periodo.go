package models

import (
	"errors"
	"strconv"
	"strings"
)

type HoraMinutos struct {
	Hora    int // 0-23
	Minutos int // 0-59
}

// NewHoraMinutosFromString toma una cadena "xx:yy" y devuelve un objeto HoraMinutos validado
func NewHoraMinutosFromString(horaStr string) (*HoraMinutos, error) {
	// Separar la cadena en horas y minutos
	parts := strings.Split(horaStr, ":")
	if len(parts) != 2 {
		return nil, errors.New("el formato de la hora debe ser 'hh:mm'")
	}

	// Convertir las partes a enteros
	hora, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.New("la hora debe ser un número entero")
	}

	minuto, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("los minutos deben ser un número entero")
	}

	// Usar la función NewHoraMinutos para validar y crear el objeto
	return NewHoraMinutos(hora, minuto)
}

func NewHoraMinutos(hora, minutos int) (*HoraMinutos, error) {
	if hora < 0 || hora > 23 {
		return nil, errors.New("la hora debe estar entre 0 y 23")
	}
	if minutos != 0 && minutos != 30 {
		return nil, errors.New("las horas deben ser a en punto o a y media")
	}
	return &HoraMinutos{Hora: hora, Minutos: minutos}, nil
}

type Periodo struct {
	HoraInicio HoraMinutos
	HoraFin    HoraMinutos
}

func esHoraValida(hora HoraMinutos) bool {
	return hora.Minutos == 0 || hora.Minutos == 30
}

func NewPeriodo(horaInicio, horaFin HoraMinutos) (*Periodo, error) {
	// Verificar que la hora de inicio es anterior a la hora de fin
	if horaInicio.Hora > horaFin.Hora || (horaInicio.Hora == horaFin.Hora && horaInicio.Minutos > horaFin.Minutos) {
		return nil, errors.New("la hora de inicio debe ser anterior a la hora de fin")
	}

	// Verificar que las horas sean en "en punto" o "y media"
	if !esHoraValida(horaInicio) || !esHoraValida(horaFin) {
		return nil, errors.New("las horas deben ser a en punto o a y media")
	}

	// Crear y devolver el Periodo si todo es válido
	periodo := &Periodo{
		HoraInicio: horaInicio,
		HoraFin:    horaFin,
	}

	return periodo, nil
}
