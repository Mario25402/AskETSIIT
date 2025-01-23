package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"testing"
)

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
		_, err := newPeriodo(prueba.inicio, prueba.fin)
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
			Grupo: Grupo{"A", "Matemáticas", "Juan Pérez"},
		},
	}

	horario := newHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "08:30")

	if lunes == nil {
		t.Fatalf("No se encontró la clase esperada")
	}

	if lunes.DiaSemana != DiaSemana.Lunes {
		t.Errorf("Día incorrecto: se esperaba %v, se obtuvo %v", DiaSemana.Lunes, lunes.DiaSemana)
	}

	if lunes.Periodo.HoraInicio != (HoraMinutos{Hora: 8, Minutos: YMedia}) {
		t.Errorf("Hora de inicio incorrecta: se esperaba %v, se obtuvo %v", HoraMinutos{Hora: 8, Minutos: YMedia}, lunes.Periodo.HoraInicio)
	}

	if lunes.Periodo.HoraFin != (HoraMinutos{Hora: 10, Minutos: YMedia}) {
		t.Errorf("Hora de fin incorrecta: se esperaba %v, se obtuvo %v", HoraMinutos{Hora: 10, Minutos: YMedia}, lunes.Periodo.HoraFin)
	}

	if lunes.Aula != "1" {
		t.Errorf("Aula incorrecta: se esperaba %q, se obtuvo %q", "1", lunes.Aula)
	}

	if lunes.Grupo.Nombre != "A" {
		t.Errorf("Nombre del grupo incorrecto: se esperaba %q, se obtuvo %q", "A", lunes.Grupo.Nombre)
	}

	if lunes.Grupo.Asignatura != "Matemáticas" {
		t.Errorf("Nombre del grupo incorrecto: se esperaba %q, se obtuvo %q", "Matemáticas", lunes.Grupo.Asignatura)
	}

	if lunes.Grupo.Profesor != "Juan Pérez" {
		t.Errorf("Profesor incorrecto: se esperaba %q, se obtuvo %q", "Juan Pérez", lunes.Grupo.Profesor)
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
			Grupo: Grupo{"A", "Matemáticas", "Juan Pérez"},
		},
	}

	horario := newHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "incorrecto")

	if lunes != nil {
		t.Errorf("Se esperaba nil, pero se obtuvo %+v", *lunes)
	}
}

func TestGetDia(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: Grupo{"A", "Matemáticas", "Juan Pérez"},
		},
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 10, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 12, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: Grupo{"A", "Física", "Ana López"},
		},
	}

	horario := newHorarioFromClases(clases)
	lunes := horario.GetDia("Lunes")

	if len(lunes) != 2 {
		t.Errorf("Número incorrecto de clases: se esperaba 2, se obtuvo %d", len(lunes))
	}
}

func TestGetDiaError(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "1",
			Grupo: Grupo{"A", "Matemáticas", "Juan Pérez"},
		},
	}

	horario := newHorarioFromClases(clases)
	lunes := horario.GetDia("incorrecto")

	if lunes != nil {
		t.Errorf("Se esperaba nil, pero se obtuvo %+v", lunes)
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
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", Grupo{"4", "Sistemas Operativos", ""})

	err := extraerProfesor(clase, "../../docs/fuentes/correcto.html")
	if err != nil {
		t.Fatalf("Error inesperado al extraer profesor: %v", err)
	}

	if clase.Grupo.Profesor != "Sofía García Pascual" {
		t.Errorf("Profesor incorrecto: se esperaba %q, se obtuvo %q", "Sofía García Pascual", clase.Grupo.Profesor)
	}
}

func TestExtraerProfesorSinGrupo(t *testing.T) {
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", Grupo{"", "Sistemas Operativos", ""})

	err := extraerProfesor(clase, "../../docs/fuentes/correcto.html")
	if err == nil {
		t.Errorf("Se esperaba un error al extraer profesor, pero no se produjo ninguno")
	}
}

func TestExtraerClasesErrorSintaxis(t *testing.T) {
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", Grupo{"4", "Sistemas Operativos", ""})
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
			Grupo: Grupo{"A", "Matemáticas", "Juan Pérez"},
		},
		{
			DiaSemana: DiaSemana.Martes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 9, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula:  "2",
			Grupo: Grupo{"A", "Química", "Carlos Gómez"},
		},
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 10, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 12, Minutos: YMedia},
			},
			Aula:  "2",
			Grupo: Grupo{"A", "Física", "Ana López"},
		},
	}

	horario := newHorarioFromClases(clases)

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
