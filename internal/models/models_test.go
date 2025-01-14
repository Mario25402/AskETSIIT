package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func crearFichTmp(html string) string {
	tmpFile, _ := os.CreateTemp("", "test.html")
	_, _ = tmpFile.WriteString(html)
	_ = tmpFile.Close()

	return tmpFile.Name()
}
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

func TestNewPeriodoStrError(t *testing.T) {
	_, err := NewPeriodoStr("hola", "12:00")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoStr("12:00", "12:30:00")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoStr("30:00", "12:30")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoStr("12:", "")
	assert.Error(t, err, "Error al crear Periodo")

	_, err = NewPeriodoStr("12:00", "ho:la")
	assert.Error(t, err, "Error al crear Periodo")
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
			Aula:  "1",
			Grupo: *NewGrupo("A", "Matemáticas", "Juan Pérez"),
		},
	}

	horario := NewHorarioFromClases(clases)
	lunes := horario.GetClase("Lunes", "incorrecto")

	assert.Nil(t, lunes, "Clase incorrecta")
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

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	_, err := extraerClases(tmpFile)
	assert.Error(t, err, "Periodo incorrecto")
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

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	clase, _ := NewClase(DiaSemana.Lunes, &Periodo{HoraInicio: HoraMinutos{Hora: 10, Minutos: EnPunto}, HoraFin: HoraMinutos{Hora: 12, Minutos: EnPunto}}, "23", *NewGrupo("", "Sistemas Operativos", ""))

	_, err := extraerProfesor(tmpFile, clase)
	assert.Error(t, err, "Error al extraer profesor")
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
	assert.NotNil(t, horario, "El horario no debería ser nulo")

	assert.Len(t, horario.Clases[DiaSemana.Lunes], 2, "Debería haber 2 clase el lunes")
	assert.Equal(t, "1", horario.GetClase("Lunes", "8:30").Aula)
	assert.Equal(t, "Juan Pérez", horario.GetClase("Lunes", "8:30").Grupo.Profesor)

	assert.Equal(t, "2", horario.GetClase("Lunes", "10:30").Aula)
	assert.Equal(t, "Ana López", horario.GetClase("Lunes", "10:30").Grupo.Profesor)

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

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

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

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	horario := NewHorarioFromFile(tmpFile)
	assert.Nil(t, horario, "Error al extraer el profesor")
}

func TestExtraerHorarioFromFileErrorClases(t *testing.T) {
	html := ``

	tmpFile := crearFichTmp(html)
	defer os.Remove(tmpFile)

	horario := NewHorarioFromFile(tmpFile)
	assert.Nil(t, horario, "Error al extraer las clases")
}
