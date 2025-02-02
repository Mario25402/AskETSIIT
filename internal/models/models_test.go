package models

import (
	"testing"
)

var horario = newHorarioFromFile("../../docs/fuentes/page.html")

func TestClasesDia(t *testing.T) {
	clases := horario.Clases["4"]

	if clases == nil {
		t.Errorf("Error al obtener las clases")
	}
}

func TestClaseHora(t *testing.T) {
	clase := horario.Clases["4"][HoraMinutos{Hora: 8, Minutos: YMedia}]

	if clase == nil {
		t.Errorf("Error al obtener la clase")
	}
}

func TestHoraAsignatura(t *testing.T) {
	clase := horario.Clases["4"][HoraMinutos{Hora: 8, Minutos: YMedia}]

	if clase.Grupo.Asignatura == "" {
		t.Errorf("Error al obtener la asignatura")
	}
}

func TestProfesorAsignatura(t *testing.T) {
	clase := horario.Clases["4"][HoraMinutos{Hora: 8, Minutos: YMedia}]

	if clase.Grupo.Profesor == "" {
		t.Errorf("Error al obtener el profesor")
	}
}

func TestAulaAsignatura(t *testing.T) {
	clase := horario.Clases["4"][HoraMinutos{Hora: 8, Minutos: YMedia}]

	if clase.Aula == "" {
		t.Errorf("Error al obtener el aula")
	}
}
