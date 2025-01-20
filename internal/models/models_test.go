package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"testing"
)

func TestStrMinutos(t *testing.T) {
	minutos, err := strMinutos("00")
	if err != nil {
		t.Fatalf("Error inesperado al obtener los minutos: %v", err)
	}
	if minutos != EnPunto {
		t.Errorf("Minutos incorrectos: se esperaba %v, se obtuvo %v", EnPunto, minutos)
	}

	minutos, err = strMinutos("30")
	if err != nil {
		t.Fatalf("Error inesperado al obtener los minutos: %v", err)
	}
	if minutos != YMedia {
		t.Errorf("Minutos incorrectos: se esperaba %v, se obtuvo %v", YMedia, minutos)
	}
}

func TestStrMinutosError(t *testing.T) {
	_, err := strMinutos("45")
	if err == nil {
		t.Errorf("Se esperaba un error al obtener los minutos, pero no se produjo ninguno")
	}
}

func TestMinutosStr(t *testing.T) {
	if minutosStr(EnPunto) != "00" {
		t.Errorf("Minutos incorrectos: se esperaba %v, se obtuvo %v", "00", minutosStr(EnPunto))
	}
	if minutosStr(YMedia) != "30" {
		t.Errorf("Minutos incorrectos: se esperaba %v, se obtuvo %v", "30", minutosStr(YMedia))
	}
	if minutosStr("hola") != "" {
		t.Errorf("Minutos incorrectos: se esperaba %v, se obtuvo %v", "", minutosStr("hola"))
	}
}

func TestNewPeriodoStrError(t *testing.T) {
	pruebas := []struct {
		inicio  string
		fin     string
		mensaje string
	}{
		{"hola", "12:00", "Inicio inválido"},
		{"12:00", "12:30:00", "Formato inválido de fin"},
		{"30:00", "12:30", "Hora inicial inválida"},
		{"12:", "", "Formato incompleto"},
		{"12:00", "ho:la", "Hora final inválida"},
	}

	for _, prueba := range pruebas {
		_, err := NewPeriodoStr(prueba.inicio, prueba.fin)
		if err == nil {
			t.Errorf("Se esperaba un error para el caso %q -> %q, pero no se produjo ninguno", prueba.inicio, prueba.fin)
		}
	}
}

func TestGetClase(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: *NewGrupo("A", "Matemáticas", "Juan Pérez"),
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "08:30")

	if lunes == nil {
		t.Fatalf("No se encontró la clase esperada")
	}
	if *lunes != clases[0] {
		t.Errorf("Clase incorrecta: se esperaba %+v, se obtuvo %+v", clases[0], *lunes)
	}
}

func TestGetClaseError(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: *NewGrupo("A", "Matemáticas", "Juan Pérez"),
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "incorrecto")

	if lunes != nil {
		t.Errorf("Se esperaba nil, pero se obtuvo %+v", *lunes)
	}
}

func TestExtraerClases(t *testing.T) {
	clases, err := extraerClases("../../docs/fuentes/correcto.html")

	if err != nil {
		t.Fatalf("Error inesperado al extraer las clases: %v", err)
	}

	if len(*clases) != 1 {
		t.Errorf("Número incorrecto de clases: se esperaba 1, se obtuvo %d", len(*clases))
	}
}

func TestExtraerClasesErrorDia(t *testing.T) {
	_, err := extraerClases("../../docs/fuentes/errorDia.html")

	if err == nil {
		t.Errorf("Se esperaba un error por día incorrecto, pero no se produjo ninguno")
	}
}

func TestExtraerClasesErrorPeriodo(t *testing.T) {
	_, err := extraerClases("../../docs/fuentes/errorPeriodo.html")

	if err == nil {
		t.Errorf("Se esperaba un error por periodo incorrecto, pero no se produjo ninguno")
	}
}

func TestExtraerProfesor(t *testing.T) {
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("4", "Sistemas Operativos", ""))

	err := extraerProfesor(clase, "../../docs/fuentes/correcto.html")
	if err != nil {
		t.Fatalf("Error inesperado al extraer profesor: %v", err)
	}

	if clase.Grupo.Profesor != "Pedro Martín Cuevas" {
		t.Errorf("Profesor incorrecto: se esperaba %q, se obtuvo %q", "Pedro Martín Cuevas", clase.Grupo.Profesor)
	}
}

func TestExtraerProfesorSinGrupo(t *testing.T) {
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("", "Sistemas Operativos", ""))

	err := extraerProfesor(clase, "../../docs/fuentes/correcto.html")
	if err == nil {
		t.Errorf("Se esperaba un error al extraer profesor, pero no se produjo ninguno")
	}
}

func TestExtraerClasesErrorSintaxis(t *testing.T) {
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("4", "Sistemas Operativos", ""))
	err := extraerProfesor(clase, "../../docs/fuentes/errorSintaxisGrupo.html")

	if err == nil {
		t.Errorf("Se esperaba un error al extraer profesor del fichero malformado, pero no se produjo ninguno")
	}
}

func TestNewHorarioFromClases(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: *NewGrupo("A", "Matemáticas", "Juan Pérez"),
		},
		{
			DiaSemana: DiaSemana.Martes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 9, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "2",
			Grupo: *NewGrupo("A", "Química", "Carlos Gómez"),
		},
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 10, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 12, Minutos: YMedia},
			},
			Aula:  "2",
			Grupo: *NewGrupo("A", "Física", "Ana López"),
		},
	}

	horario := NewHorarioFromClases(clases)

	if len(horario.Clases[DiaSemana.Lunes]) != 2 {
		t.Errorf("Número incorrecto de clases el lunes: se esperaba %d, se obtuvo %d", 2, len(horario.Clases[DiaSemana.Lunes]))
	}
}

func TestExtraerHorarioFromFile(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/correcto.html")

	if len(horario.Clases["1"]) != 1 {
		t.Errorf("Número incorrecto de clases: se esperaba %d, se obtuvo %d", 1, len(horario.Clases["1"]))
	}
}

func TestExtraerHorarioFromFileErrorProfesor(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/errorSinProfesor.html")

	if horario != nil {
		t.Errorf("Se esperaba un horario nulo debido a la falta de un profesor, pero se obtuvo: %+v", horario)
	}
}

func TestExtraerHorarioFromFileErrorClases(t *testing.T) {
	horario := NewHorarioFromFile("../../docs/fuentes/vacio.html")

	if horario != nil {
		t.Errorf("Se esperaba un horario nulo debido a la falta de un profesor, pero se obtuvo: %+v", horario)
	}
}
