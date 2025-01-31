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

func TestGetClaseHora(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetClaseHora("4", "08:30") != "Infraestructura Virtual (Especialidad Tecnologías de la Información)" {
		t.Errorf("Error al obtener la asignatura")
	}
}

func TestGetHoraAsignatura(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetHoraAsignatura("Infraestructura Virtual (Especialidad Tecnologías de la Información)", "4") == nil {
		t.Errorf("Error al obtener la hora")
	}
}

func TestGetHoraAsignaturaError(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetHoraAsignatura("MAL", "4") != nil {
		t.Errorf("Error al obtener la hora")
	}
}

func TestGetProfesorAsignatura(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetProfesorAsignatura("Infraestructura Virtual (Especialidad Tecnologías de la Información)", "08:30", "4") != "Juan Julián Merelo Guervos" {
		t.Errorf("Error al obtener el profesor")
	}
}

func TestGetProfesorAsignaturaError(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetProfesorAsignatura("Infraestructura Virtual (Especialidad Tecnologías de la Información)", "8:30", "MAL") != "" {
		t.Errorf("Error al obtener el profesor")
	}
}

func TestGetAulaAsignatura(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetAulaAsignatura("Infraestructura Virtual (Especialidad Tecnologías de la Información)", "08:30", "4") != "-1.2" {
		t.Errorf("Error al obtener el aula")
	}
}

func TestGetAulaAsignaturaError(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if horario.GetAulaAsignatura("MAL", "8:30", "4") != "" {
		t.Errorf("Error al obtener el aula")
	}
}
