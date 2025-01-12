package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////
// Funciones auxiliares

func crearFichTmp(html string) string {
	tmpFile, _ := os.CreateTemp("", "test.html")
	_, _ = tmpFile.WriteString(html)
	_ = tmpFile.Close()

	return tmpFile.Name()
}

// /////////////////////////////////////
// Grupo

func TestNewGrupo(t *testing.T) {
	grupo := NewGrupo("9", "Sistemas Operativos", "Pedro Martín Cuevas")
	assert.Equal(t, Grupo{Nombre: "9", Asignatura: "Sistemas Operativos", Profesor: "Pedro Martín Cuevas"}, *grupo, "Grupo erroneo")
}

// /////////////////////////////////////
// DiaSemana

func TestDiaStr(t *testing.T) {
	assert.Equal(t, "1", DiaSemana.DiaStr(DiaSemana.Lunes), "Día erroneo")
	assert.Equal(t, "2", DiaSemana.DiaStr(DiaSemana.Martes), "Día erroneo")
	assert.Equal(t, "3", DiaSemana.DiaStr(DiaSemana.Miercoles), "Día erroneo")
	assert.Equal(t, "4", DiaSemana.DiaStr(DiaSemana.Jueves), "Día erroneo")
	assert.Equal(t, "5", DiaSemana.DiaStr(DiaSemana.Viernes), "Día erroneo")
	assert.Equal(t, "", DiaSemana.DiaStr(""), "Día erroneo")
}

func TestStrDia(t *testing.T) {
	assert.Equal(t, DiaSemana.Lunes, DiaSemana.StrDia("1"), "Día erroneo")
	assert.Equal(t, DiaSemana.Martes, DiaSemana.StrDia("2"), "Día erroneo")
	assert.Equal(t, DiaSemana.Miercoles, DiaSemana.StrDia("3"), "Día erroneo")
	assert.Equal(t, DiaSemana.Jueves, DiaSemana.StrDia("4"), "Día erroneo")
	assert.Equal(t, DiaSemana.Viernes, DiaSemana.StrDia("5"), "Día erroneo")
	assert.Equal(t, DiaSemana.DiaSemana(""), DiaSemana.StrDia(""), "Día erroneo")
}

///////////////////////////////////////
// MinutosPosibles

func TestStrMinutos(t *testing.T) {
	minutos, err := strMinutos("00")
	assert.Nil(t, err, "Error al obtener los minutos")
	assert.Equal(t, EnPunto, minutos, "Minutos erroneos")

	minutos, err = strMinutos("30")
	assert.Nil(t, err, "Error al obtener los minutos")
	assert.Equal(t, YMedia, minutos, "Minutos erroneos")
}

func TestStrMinutosError(t *testing.T) {
	_, err := strMinutos("45")
	assert.Error(t, err, "Error al obtener los minutos")
}

func TestMinutosStr(t *testing.T) {
	assert.Equal(t, "00", minutosStr(EnPunto), "Minutos erroneos")
	assert.Equal(t, "30", minutosStr(YMedia), "Minutos erroneos")
	assert.Equal(t, "", minutosStr("hola"), "Minutos erroneos")
}

///////////////////
// HoraMinutos

func TestNewHoraMinutos(t *testing.T) {
	horMin, err := newHoraMinutos(01, EnPunto)
	assert.Equal(t, HoraMinutos{Hora: 01, Minutos: EnPunto}, *horMin, "HoraMinutos erroneo")
	assert.Nil(t, err, "Error al crear HoraMinutos")
}

func TestNewHoraMinutosError(t *testing.T) {
	_, err := newHoraMinutos(31, EnPunto)
	assert.Error(t, err, "Error al crear HoraMinutos")
}

func TestNewHoraMinutosStr(t *testing.T) {
	horMin, err := NewHoraMinutosStr("12:30")
	assert.Equal(t, HoraMinutos{Hora: 12, Minutos: YMedia}, *horMin, "HoraMinutos erroneo")
	assert.Nil(t, err, "Error al crear HoraMinutos")
}

func TestNewHoraMinutosStrError(t *testing.T) {
	_, err := NewHoraMinutosStr("12:")
	assert.Error(t, err, "Error al crear HoraMinutos")

	_, err = NewHoraMinutosStr("hola:30")
	assert.Error(t, err, "Error al crear HoraMinutos")
}

func TestNewHoraMinutosSplit(t *testing.T) {
	horMin, err := NewHoraMinutosSplit(23, "30")
	assert.Equal(t, HoraMinutos{Hora: 23, Minutos: YMedia}, *horMin, "HoraMinutos erroneo")
	assert.Nil(t, err, "Error al crear HoraMinutos")
}

func TestNewHoraMinutosSplitError(t *testing.T) {
	_, err := NewHoraMinutosSplit(17, "45")
	assert.Error(t, err, "Error al crear HoraMinutos")
}

///////////////////
// Periodo

func TestNewPeriodo(t *testing.T) {
	ini, _ := newHoraMinutos(10, EnPunto)
	fin, _ := newHoraMinutos(12, EnPunto)

	periodo, err := newPeriodo(*ini, *fin)
	assert.Equal(t, Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, *periodo, "Periodo erroneo")
	assert.Nil(t, err, "Error al crear Periodo")
}

func TestNewPeriodoError(t *testing.T) {
	ini, _ := newHoraMinutos(10, EnPunto)
	fin, _ := newHoraMinutos(12, EnPunto)

	_, err := newPeriodo(*fin, *ini)
	assert.Error(t, err, "Error al crear Periodo")
}

func TestNewPeriodoStr(t *testing.T) {
	periodo, err := NewPeriodoStr("10:00", "12:00")
	assert.Equal(t, Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, *periodo, "Periodo erroneo")
	assert.Nil(t, err, "Error al crear Periodo")
}

func TestNewPeriodoStrError(t *testing.T) {
	_, err := NewPeriodoStr("hola", "12:00")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoStr("12:00", "12:30:00")
	assert.Error(t, err, "Error al crear Periodo")
}

func TestNewPeriodoSplit(t *testing.T) {
	periodo, err := NewPeriodoSplit(10, "00", 12, "00")
	assert.Equal(t, Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, *periodo, "Periodo erroneo")
	assert.Nil(t, err, "Error al crear Periodo")
}

func TestNewPeriodoSplitError(t *testing.T) {
	_, err := NewPeriodoSplit(10, "ab", 12, "30")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoSplit(10, "00", 12, "12:30")
	assert.Error(t, err, "Error al crear Periodo")
}

///////////////////////////////////////
// Clase

func TestNewClase(t *testing.T) {
	periodo, _ := NewPeriodoStr("10:00", "12:00")
	clase, err := NewClase(DiaSemana.Lunes, periodo, "23", Grupo{Asignatura: "Sistemas Operativos", Nombre: "9"})

	assert.Equal(t, Clase{DiaSemana: DiaSemana.Lunes, Periodo: periodo, Aula: "23", Grupo: Grupo{Asignatura: "Sistemas Operativos", Nombre: "9"}}, *clase, "Clase erronea")
	assert.Nil(t, err, "Error al crear la clase")
}

///////////////////////////////////////
// Horario

func TestGetClase(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula: "1",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Matemáticas",
				Profesor:   "Juan Pérez",
			},
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "08:30")

	assert.Equal(t, clases[0], *lunes, "Día incorrecto")
}

func TestGetClaseError(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula: "1",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Matemáticas",
				Profesor:   "Juan Pérez",
			},
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "incorrecto")

	assert.Nil(t, lunes, "Clase incorrecta")
}

/*func TestGetDia(t *testing.T) {
	clases := []Clase{
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 8, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula: "1",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Matemáticas",
				Profesor:   "Juan Pérez",
			},
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetDia("Lunes")

	assert.Len(t, *lunes, 1, "Número incorrecto de clases")
	assert.Equal(t, clases[0], (*lunes)[HoraMinutos{Hora: 8, Minutos: YMedia}], "Clase incorrecta")
}*/

func TestGetDiaError(t *testing.T) {
}

////////////////////

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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clases, err := extraerClases(tmpFile)
	assert.Nil(t, err, "Error al extraer las clases")
	assert.Len(t, *clases, 1, "Número incorrecto de clases")

	assert.Equal(t, "Sistemas Operativos", (*clases)[0].Grupo.Asignatura, "Asignatura incorrecta")
	assert.Equal(t, DiaSemana.Lunes, (*clases)[0].DiaSemana, "Día incorrecto")
	assert.Equal(t, "9", (*clases)[0].Grupo.Nombre, "Grupo incorrecto")
	assert.Equal(t, "23", (*clases)[0].Aula, "Aula incorrecta")
	assert.Equal(t, "9:30", GetHoraMinutosStr(&(*clases)[0].Periodo.HoraInicio), "Hora de inicio incorrecta")
	assert.Equal(t, "11:30", GetHoraMinutosStr(&(*clases)[0].Periodo.HoraFin), "Hora de fin incorrecta")
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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	_, err := extraerClases(tmpFile)
	assert.Error(t, err, "Día incorrecto")
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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	_, err := extraerClases(tmpFile)
	assert.Error(t, err, "Periodo incorrecto")
}

///////////////////

func TestExtraerProfesor(t *testing.T) {
	html := `
    <li class="profesor">
        <a href="https://www.ugr.es/personal/8e43782373ab33f672b24a92f8eb9e10">Pedro Martín Cuevas</a>
        <span class="grupos">
              Grupo&nbsp;
            4
        </span>
    </li>`

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	// Clase sin profesor
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", Grupo{Asignatura: "Sistemas Operativos", Nombre: "4", Profesor: ""})

	clase, err := extraerProfesor(tmpFile, clase)
	assert.Equal(t, "Pedro Martín Cuevas", clase.Grupo.Profesor, "Profesor incorrecto")
	assert.Nil(t, err, "Error al extraer profesor")
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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	// Clase sin profesor
	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", Grupo{Asignatura: "Sistemas Operativos", Nombre: "", Profesor: ""})

	_, err := extraerProfesor(tmpFile, clase)
	assert.Error(t, err, "Error al extraer profesor")
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
			Aula: "1",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Matemáticas",
				Profesor:   "Juan Pérez",
			},
		},
		{
			DiaSemana: DiaSemana.Lunes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 10, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 12, Minutos: YMedia},
			},
			Aula: "2",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Física",
				Profesor:   "Ana López",
			},
		},
		{
			DiaSemana: DiaSemana.Martes,
			Periodo: &Periodo{
				HoraInicio: HoraMinutos{Hora: 9, Minutos: YMedia},
				HoraFin:    HoraMinutos{Hora: 10, Minutos: YMedia},
			},
			Aula: "2",
			Grupo: Grupo{
				Nombre:     "A",
				Asignatura: "Química",
				Profesor:   "Carlos Gómez",
			},
		},
	}

	// Llamar a la función
	horario := NewHorarioFromClases(clases)

	// Verificar la estructura del horario
	assert.NotNil(t, horario, "El horario no debería ser nulo")

	assert.Len(t, horario.Clases[DiaSemana.Lunes], 2, "Debería haber 2 clase el lunes")
	assert.Equal(t, "1", horario.GetClase("Lunes", "8:30").Aula)
	assert.Equal(t, "Juan Pérez", horario.GetClase("Lunes", "8:30").Grupo.Profesor)

	assert.Equal(t, "2", horario.GetClase("Lunes", "10:30").Aula)
	assert.Equal(t, "Ana López", horario.GetClase("Lunes", "10:30").Grupo.Profesor)

	// Verificar las clases en el martes
	assert.Len(t, horario.Clases[DiaSemana.Martes], 1, "Debería haber 1 clase el martes")
	assert.Equal(t, "2", horario.GetClase("Martes", "9:30").Aula)
	assert.Equal(t, "Carlos Gómez", horario.GetClase("Martes", "9:30").Grupo.Profesor)
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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	// Extraer el horario
	horario := NewHorarioFromFile(tmpFile)
	assert.NotNil(t, horario, "Error al extraer el horario")
	assert.Len(t, horario.Clases[DiaSemana.Lunes], 1, "Número incorrecto de clases")
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

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	// Extraer el horario
	horario := NewHorarioFromFile(tmpFile)
	assert.Nil(t, horario, "Error al extraer el profesor")
}

func TestExtraerHorarioFromFileErrorClases(t *testing.T) {
	html := ``

	// Crear fichero temporal
	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	// Extraer el horario
	horario := NewHorarioFromFile(tmpFile)
	assert.Nil(t, horario, "Error al extraer las clases")
}

///////////////////////////////////////
