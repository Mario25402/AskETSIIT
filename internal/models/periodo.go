package models

import (
	"errors"
	"fmt"
	"time"
)

type Periodo struct {
	HoraInicio time.Time
	HoraFin    time.Time
}

func (p Periodo) String() string {
	return fmt.Sprintf("%s-%s", p.HoraInicio.Format("15:04"), p.HoraFin.Format("15:04"))
}

func esHoraValida(hora time.Time) bool {
	return hora.Minute() == 0 || hora.Minute() == 30
}

func NuevoPeriodo(horaInicioStr, horaFinStr string) (*Periodo, error) {
	// Formato esperado "15:04" para horas en formato de 24 horas
	const formato = "15:04"

	horaInicio, err := time.Parse(formato, horaInicioStr)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la hora de inicio: %v", err)
	}

	horaFin, err := time.Parse(formato, horaFinStr)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la hora de fin: %v", err)
	}

	// Verificar que la hora de inicio es anterior a la hora de fin
	if horaInicio.After(horaFin) {
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
