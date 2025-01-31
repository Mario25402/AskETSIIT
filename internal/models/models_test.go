package models

import (
	"testing"
)

func TestExtraerHorario(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario == nil {
		t.Errorf("Error al extraer el horario")
	}
}
