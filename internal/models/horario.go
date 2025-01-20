package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

type Horario struct {
	Clases map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase
}

func (horario Horario) GetClase(dia DiaSemana.DiaSemana, hora string) *Clase {
	time, err := newHoraMinutos(hora)
	if err != nil {
		return nil
	}

	clases := horario.GetDia(dia)
	return clases[*time]
}

func (horario Horario) GetDia(dia DiaSemana.DiaSemana) map[HoraMinutos]*Clase {
	if dia < DiaSemana.Lunes || dia > DiaSemana.Viernes {
		return nil
	}

	return horario.Clases[dia]
}

func comprobarAdicionClase(clases *[]Clase, asignatura, grupo, aula *string, dia *DiaSemana.DiaSemana, periodo *Periodo) {
	if *asignatura != "" && *dia != "" && *grupo != "" && *aula != "" && periodo != nil {
		clase := Clase{
			DiaSemana: *dia,
			Periodo:   periodo,
			Aula:      *aula,
			Grupo:     Grupo{Nombre: *grupo, Asignatura: *asignatura, Profesor: ""},
		}

		*clases = append(*clases, clase)

		*dia = ""
		*aula = ""
		*grupo = ""
		periodo = nil
	}
}

func procesarAsignatura(linea string, asignatura *string) {
	expAsignatura := regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`)

	if matches := expAsignatura.FindStringSubmatch(linea); matches != nil {
		*asignatura = strings.TrimSpace(matches[1])
	}
}

func procesarDia(linea string, dia *DiaSemana.DiaSemana) {
	expDia := regexp.MustCompile(`<div class=\"clase dia-(\d)\"`)

	if matches := expDia.FindStringSubmatch(linea); matches != nil {
		if "1" < matches[1] && matches[1] > "5" {
			return
		}

		*dia = DiaSemana.DiaSemana(matches[1])
	}
}

func procesarGrupo(linea string, grupo *string) {
	expGrupo := regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-Za-z]|\d{1,2})</div>`)

	if matches := expGrupo.FindStringSubmatch(linea); matches != nil {
		*grupo = matches[1]
	}
}

func procesarAula(linea string, aula *string) {
	expAula := regexp.MustCompile(`<div>Aula:\s*(\d+)</div>`)

	if matches := expAula.FindStringSubmatch(linea); matches != nil {
		*aula = matches[1]
	}
}

func procesarPeriodo(linea string, periodo **Periodo) {
	expPeriodo := regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`)

	if matches := expPeriodo.FindStringSubmatch(linea); matches != nil {
		newPeriodo, err := newPeriodo(matches[1], matches[2])
		if err != nil {
			return
		}

		*periodo = newPeriodo
	}
}

func extraerClases(fileName string) (*[]Clase, error) {
	var clases []Clase
	var periodo *Periodo
	var dia DiaSemana.DiaSemana
	var aula, grupo, asignatura string

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		procesarDia(linea, &dia)
		procesarAula(linea, &aula)
		procesarGrupo(linea, &grupo)
		procesarPeriodo(linea, &periodo)
		procesarAsignatura(linea, &asignatura)

		comprobarAdicionClase(&clases, &asignatura, &grupo, &aula, &dia, periodo)
	}

	if len(clases) == 0 {
		return nil, errors.New("no se han encontrado clases")
	}

	return &clases, nil
}

func establecerProfesor(clase *Clase, profesor, cadena *string) {
	if clase != nil && profesor != nil && cadena != nil {
		var grupos []string

		grupos = append(grupos, strings.Split(*cadena, ",")...)

		for _, grupo := range grupos {
			if strings.TrimSpace(grupo) == clase.Grupo.Nombre {
				clase.Grupo.setProfesor(*profesor)
				*profesor = ""
				break
			}
		}
	}
}

func procesadorProfesor(linea string, profesor *string, leer *bool) *string {
	expNombre := regexp.MustCompile(`<a href=\"https://www.ugr.es/personal/[^>]*\">([^<]+)</a>`)

	if matches := expNombre.FindStringSubmatch(linea); matches != nil {
		*profesor = strings.TrimSpace(matches[1])
		return nil
	}

	expTitulo := regexp.MustCompile(`Grupos?&nbsp;`)
	if expTitulo.MatchString(linea) {
		*leer = true
		return nil
	}

	if *leer {
		expGrupos := regexp.MustCompile(`([A-Z]|\d{1,2})(,\s*([A-Z]|\d{1,2}))*\s*(y\s*([A-Z]|\d{1,2}))?$`)

		if matches := expGrupos.FindStringSubmatch(linea); matches != nil {
			*leer = false
			result := strings.ReplaceAll(matches[0], " y ", ",")
			return &result
		}
	}

	return nil
}

func extraerProfesor(clase *Clase, fileName string) error {
	var profesor string
	var leer bool

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		if grupos := procesadorProfesor(linea, &profesor, &leer); grupos != nil {
			establecerProfesor(clase, &profesor, grupos)
		}
	}

	if clase.Grupo.Profesor == "" {
		return errors.New("no se ha encontrado profesor")
	}

	return nil
}

func newHorarioFromClases(clases []Clase) Horario {
	horario := Horario{Clases: make(map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase)}

	for _, clase := range clases {
		dia := clase.DiaSemana

		if horario.Clases[dia] == nil {
			horario.Clases[dia] = make(map[HoraMinutos]*Clase)
		}

		horaInicio := clase.Periodo.HoraInicio
		horario.Clases[dia][horaInicio] = &clase
	}

	return horario
}

func NewHorarioFromFile(file string) *Horario {
	clases, err := extraerClases(file)
	if err != nil {
		return nil
	}

	for iteracion, clase := range *clases {
		err := extraerProfesor(&clase, file)
		if err != nil {
			return nil
		}

		(*clases)[iteracion] = clase
	}

	horario := newHorarioFromClases(*clases)
	return &horario
}
