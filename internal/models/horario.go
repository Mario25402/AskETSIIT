package models

import "errors"

type Horario struct {
	Clases []Clase
}

func (h *Horario) AgregarClase(clase Clase) error {
	if err := h.validarSolapamiento(clase); err != nil {
		return err
	}
	h.Clases = append(h.Clases, clase)
	return nil
}

// Valida que no haya solapamiento entre clases
func (h *Horario) validarSolapamiento(nuevaClase Clase) error {

	for _, clase := range h.Clases {
		if clase.Aula == nuevaClase.Aula && clase.Dia == nuevaClase.Dia {
			if clase.Periodo.HoraInicio.Before(nuevaClase.Periodo.HoraFin) &&
				clase.Periodo.HoraFin.After(nuevaClase.Periodo.HoraInicio) {
				return errors.New("no se ha podido añadir, se solapan clases")
			}
		}
	}
	return nil
}
