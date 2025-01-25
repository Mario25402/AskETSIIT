package models

import (
	"testing"
)

func TestGetClaseDia(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.DiaSemana != "4" {
		t.Errorf("Día incorrecto: se esperaba %v, se obtuvo %v", "4", clase.DiaSemana)
	}
}

func TestGetClaseHoraInicio(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Periodo.HoraInicio != (HoraMinutos{Hora: 8, Minutos: YMedia}) {
		t.Errorf("Hora de inicio incorrecta: se esperaba %v, se obtuvo %v", HoraMinutos{Hora: 8, Minutos: YMedia}, clase.Periodo.HoraInicio)
	}
}

func TestGetClaseHoraFin(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Periodo.HoraFin != (HoraMinutos{Hora: 10, Minutos: YMedia}) {
		t.Errorf("Hora de fin incorrecta: se esperaba %v, se obtuvo %v", HoraMinutos{Hora: 10, Minutos: YMedia}, clase.Periodo.HoraFin)
	}
}

func TestGetClaseAula(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Aula != "-1.2" {
		t.Errorf("Aula incorrecta: se esperaba %q, se obtuvo %q", "-1.2", clase.Aula)
	}
}

func TestGetClaseNombreGrupo(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Grupo.Nombre != "1" {
		t.Errorf("Nombre del grupo incorrecto: se esperaba %q, se obtuvo %q", "1", clase.Grupo.Nombre)
	}
}

func TestGetClaseAsignatura(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Grupo.Asignatura != "Infraestructura Virtual (Especialidad Tecnologías de la Información)" {
		t.Errorf("Nombre del grupo incorrecto: se esperaba %q, se obtuvo %q", "Infraestructura Virtual (Especialidad Tecnologías de la Información)", clase.Grupo.Asignatura)
	}
}

func TestGetClaseProfesor(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")
	clase := horario.GetClase("4", "08:30")

	if clase.Grupo.Profesor != "Juan Julián Merelo Guervos" {
		t.Errorf("Profesor incorrecto: se esperaba %q, se obtuvo %q", "Juan Julián Merelo Guervos", clase.Grupo.Profesor)
	}
}
