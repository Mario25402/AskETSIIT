package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"os"
	"testing"
)

func crearFichTmp(html string) string {
	tmpFile, _ := os.CreateTemp("", "test.html")
	_, _ = tmpFile.WriteString(html)
	_ = tmpFile.Close()

	return tmpFile.Name()
}
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
	html := `
	<body>
		<h1 class="page-title">Sistemas Operativos</h1>
		<div class="clase dia-1" style="height: 76px; min-height: 76px; width:100%; left:0%; top: 50%; >
			<div class="grupo"><span>Grupo:</span> 9</div>

			<div class="otros-datos">
				<div>Aula: 23</div>
				<div>Fecha de inicio: 16/09/2024</div>
				<div>Fecha final: 20/12/2024</div>
				<div>Horario: De 09:30 a 11:30</div>
			</div>
		</div>
	</body>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clases, err := extraerClases(tmpFile)
	if err != nil {
		t.Fatalf("Error inesperado al extraer las clases: %v", err)
	}

	if len(*clases) != 1 {
		t.Errorf("Número incorrecto de clases: se esperaba 1, se obtuvo %d", len(*clases))
	}

	clase := (*clases)[0]
	if clase.Grupo.Asignatura != "Sistemas Operativos" {
		t.Errorf("Asignatura incorrecta: se esperaba %q, se obtuvo %q", "Sistemas Operativos", clase.Grupo.Asignatura)
	}
	if clase.DiaSemana != "1" {
		t.Errorf("Día incorrecto: se esperaba %v, se obtuvo %v", "1", clase.DiaSemana)
	}
	if clase.Grupo.Nombre != "9" {
		t.Errorf("Grupo incorrecto: se esperaba %q, se obtuvo %q", "9", clase.Grupo.Nombre)
	}
	if clase.Aula != "23" {
		t.Errorf("Aula incorrecta: se esperaba %q, se obtuvo %q", "23", clase.Aula)
	}
	if GetHoraMinutosStr(&clase.Periodo.HoraInicio) != "9:30" {
		t.Errorf("Hora de inicio incorrecta: se esperaba %q, se obtuvo %q", "9:30", GetHoraMinutosStr(&clase.Periodo.HoraInicio))
	}
	if GetHoraMinutosStr(&clase.Periodo.HoraFin) != "11:30" {
		t.Errorf("Hora de fin incorrecta: se esperaba %q, se obtuvo %q", "11:30", GetHoraMinutosStr(&clase.Periodo.HoraFin))
	}
}

func TestExtraerClasesErrorDia(t *testing.T) {
	html := `
	<body>
		<h1 class="page-title">Sistemas Operativos</h1>
		<div class="clase dia-9" style="height: 76px; min-height: 76px; width:100%; left:0%; top: 50%; >
			<div class="grupo"><span>Grupo:</span> 9</div>

			<div class="otros-datos">
				<div>Aula: 23</div>
				<div>Fecha de inicio: 16/09/2024</div>
				<div>Fecha final: 20/12/2024</div>
				<div>Horario: De 09:30 a 11:30</div>
			</div>
		</div>
	</body>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	_, err := extraerClases(tmpFile)
	if err == nil {
		t.Errorf("Se esperaba un error por día incorrecto, pero no se produjo ninguno")
	}
}

func TestExtraerClasesErrorPeriodo(t *testing.T) {
	html := `
	<body>
		<h1 class="page-title">Sistemas Operativos</h1>
		<div class="clase dia-1" style="height: 76px; min-height: 76px; width:100%; left:0%; top: 50%; >
			<div class="grupo"><span>Grupo:</span> 9</div>

			<div class="otros-datos">
				<div>Aula: 23</div>
				<div>Fecha de inicio: 16/09/2024</div>
				<div>Fecha final: 20/12/2024</div>
				<div>Horario: De 11:30 a 10:30</div>
			</div>
		</div>
	</body>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	_, err := extraerClases(tmpFile)
	if err == nil {
		t.Errorf("Se esperaba un error por periodo incorrecto, pero no se produjo ninguno")
	}
}

func TestExtraerProfesor(t *testing.T) {
	html := `
    <li class="profesor">
        <a href="https://www.ugr.es/personal/8e43782373ab33f672b24a92f8eb9e10">Pedro Martín Cuevas</a>
        <span class="grupos">
              Grupo&nbsp;
            4
        </span>
    </li>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("4", "Sistemas Operativos", ""))

	clase, err := extraerProfesor(tmpFile, clase)
	if err != nil {
		t.Fatalf("Error inesperado al extraer profesor: %v", err)
	}

	if clase.Grupo.Profesor != "Pedro Martín Cuevas" {
		t.Errorf("Profesor incorrecto: se esperaba %q, se obtuvo %q", "Pedro Martín Cuevas", clase.Grupo.Profesor)
	}
}

func TestExtraerProfesorSinGrupo(t *testing.T) {
	html := `
    <li class="profesor">
        <a href="https://www.ugr.es/personal/8e43782373ab33f672b24a92f8eb9e10">Pedro Martín Cuevas</a>
        <span class="grupos">
              Grupo&nbsp;
            4
        </span>
    </li>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("", "Sistemas Operativos", ""))

	_, err := extraerProfesor(tmpFile, clase)
	if err == nil {
		t.Errorf("Se esperaba un error al extraer profesor, pero no se produjo ninguno")
	}
}

func TestExtraerClasesErrorSintaxis(t *testing.T) {
	html := `
    <li class="profesor">
        <a href="https://www.ugr.es/personal/8e43782373ab33f672b24a92f8eb9e10">Pedro Martín Cuevas</a>
        <span class="grupos">
            4
        </span>
    </li>`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("4", "Sistemas Operativos", ""))

	_, err := extraerProfesor(tmpFile, clase)
	if err == nil {
		t.Errorf("Se esperaba un error al extraer profesor del fichero malformado, pero no se produjo ninguno")
	}
}

func TestNewHorarioFromClases(t *testing.T) {
	// Crear clases de prueba
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
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 10, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 12, Minutos: YMedia},
			},
			Aula:  "2",
			Grupo: *NewGrupo("A", "Física", "Ana López"),
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
	}

	horario := NewHorarioFromClases(clases)

	if len(horario.Clases[DiaSemana.Lunes]) != 2 {
		t.Errorf("Número incorrecto de clases el lunes: se esperaba %d, se obtuvo %d", 2, len(horario.Clases[DiaSemana.Lunes]))
	}

	if clase := horario.GetClase("Lunes", "8:30"); clase.Aula != "1" || clase.Grupo.Profesor != "Juan Pérez" {
		t.Errorf("Clase incorrecta para las 8:30 del lunes: Aula esperada %q, Profesor esperado %q", "1", "Juan Pérez")
	}

	if clase := horario.GetClase("Lunes", "10:30"); clase.Aula != "2" || clase.Grupo.Profesor != "Ana López" {
		t.Errorf("Clase incorrecta para las 10:30 del lunes: Aula esperada %q, Profesor esperado %q", "2", "Ana López")
	}

	if len(horario.Clases[DiaSemana.Martes]) != 1 {
		t.Errorf("Número incorrecto de clases el martes: se esperaba %d, se obtuvo %d", 1, len(horario.Clases[DiaSemana.Martes]))
	}

	if clase := horario.GetClase("Martes", "9:30"); clase.Aula != "2" || clase.Grupo.Profesor != "Carlos Gómez" {
		t.Errorf("Clase incorrecta para las 9:30 del martes: Aula esperada %q, Profesor esperado %q", "2", "Carlos Gómez")
	}
}

func TestExtraerHorarioFromFile(t *testing.T) {
	html := `
	<body>
		<h1 class="page-title">Sistemas Operativos</h1>
		<div class="clase dia-1" style="height: 76px; min-height: 76px; width:100%; left:0%; top: 50%; >
			<div class="grupo"><span>Grupo:</span> 4</div>

			<div class="otros-datos">
				<div>Aula: 23</div>
				<div>Fecha de inicio: 16/09/2024</div>
				<div>Fecha final: 20/12/2024</div>
				<div>Horario: De 09:30 a 11:30</div>
			</div>

			<ul>
				<li class="profesor">
					<a href="https://www.ugr.es/personal/8e43782373ab33f672b24a92f8eb9e10">Pedro Martín Cuevas</a>
					<span class="grupos">
							Grupo&nbsp;
						4
					</span>
				</li>
			</ul>
		</div>
	</body>
	`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	horario := NewHorarioFromFile(tmpFile)
	if len(horario.Clases[DiaSemana.Lunes]) != 1 {
		t.Errorf("Número incorrecto de clases: se esperaba %d, se obtuvo %d", 1, len(horario.Clases[DiaSemana.Lunes]))
	}
}

func TestExtraerHorarioFromFileErrorProfesor(t *testing.T) {
	html := `
	<body>
		<h1 class="page-title">Sistemas Operativos</h1>
		<div class="clase dia-1" style="height: 76px; min-height: 76px; width:100%; left:0%; top: 50%; >
			<div class="grupo"><span>Grupo:</span> 4</div>

			<div class="otros-datos">
				<div>Aula: 23</div>
				<div>Fecha de inicio: 16/09/2024</div>
				<div>Fecha final: 20/12/2024</div>
				<div>Horario: De 09:30 a 11:30</div>
			</div>
		</div>
	</body>
	`

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	horario := NewHorarioFromFile(tmpFile)
	if horario != nil {
		t.Errorf("Se esperaba un horario nulo debido a la falta de un profesor, pero se obtuvo: %+v", horario)
	}
}

func TestExtraerHorarioFromFileErrorClases(t *testing.T) {
	html := ``

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	horario := NewHorarioFromFile(tmpFile)
	if horario != nil {
		t.Errorf("Se esperaba un horario nulo debido a la falta de un profesor, pero se obtuvo: %+v", horario)
	}
}
