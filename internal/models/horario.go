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

func procesarAsignatura(linea string, asignatura *string) {
	expAsignatura := regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`)

	if matches := expAsignatura.FindStringSubmatch(linea); matches != nil {
		*asignatura = strings.TrimSpace(matches[1])
	}
}

func procesarDia(linea string, dia *DiaSemana.DiaSemana) {
	expDia := regexp.MustCompile(`<div class=\"clase dia-(\d)\"`)

	if matches := expDia.FindStringSubmatch(linea); matches != nil {
		*dia = DiaSemana.DiaSemana(matches[1])
	}
}

func procesarGrupo(linea string, grupo *string) {
	expGrupo := regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-F]|[A-F][1-3]|\d{1,2})</div>`)

	if matches := expGrupo.FindStringSubmatch(linea); matches != nil {
		*grupo = matches[1]
	}
}

func procesarAula(linea string, aula *string) {
	expAula := regexp.MustCompile(`<div>Aula:\s*(-?[A-Za-z0-9][A-Za-z0-9.][A-Za-z0-9])</div>`)

	if matches := expAula.FindStringSubmatch(linea); matches != nil {
		*aula = matches[1]
	}
}

func procesarPeriodo(linea string, periodo **Periodo) {
	expPeriodo := regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`)

	if matches := expPeriodo.FindStringSubmatch(linea); matches != nil {
		*periodo = newPeriodo(matches[1], matches[2])
	}
}

func procesadorProfesor(linea string, profesor *string, leer *bool) []string {
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

			cadena := strings.TrimSpace(matches[0])
			grupos := regexp.MustCompile(`\s*(,|y)\s*`)
			return grupos.Split(cadena, -1)
		}
	}

	return nil
}

func extraerClases(fileName string) (*[]Clase, error) {
	var clases []Clase
	var periodo *Periodo
	var dia DiaSemana.DiaSemana
	var aula, grupo, asignatura string

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for ok := scanner.Scan(); ok; ok = scanner.Scan() {
		linea := scanner.Text()

		procesarDia(linea, &dia)
		procesarAula(linea, &aula)
		procesarGrupo(linea, &grupo)
		procesarPeriodo(linea, &periodo)
		procesarAsignatura(linea, &asignatura)

		clase, err := newClase(dia, periodo, aula, Grupo{Nombre: grupo, Asignatura: asignatura, Profesor: ""})

		if err == nil {
			clases = append(clases, *clase)

			dia = ""
			aula = ""
			grupo = ""
			asignatura = ""
			periodo = nil
		}
	}

	if len(clases) == 0 {
		return nil, errors.New("no se han encontrado clases")
	}

	return &clases, nil
}

func establecerProfesor(clase *Clase, profesor *string, grupos []string) {
	if clase != nil && profesor != nil && grupos != nil {
		for _, grupo := range grupos {
			if grupo == clase.Grupo.Nombre {
				clase.Grupo.Profesor = *profesor
				*profesor = ""
				break
			}
		}
	}
}

func extraerProfesor(clase *Clase, fileName string) {
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
		extraerProfesor(&clase, file)
		(*clases)[iteracion] = clase
	}

	horario := newHorarioFromClases(*clases)
	return &horario
}

func (horario Horario) GetClase(dia DiaSemana.DiaSemana, hora string) *Clase {
	time := newHoraMinutos(hora)
	clases := horario.GetDia(dia)
	return clases[*time]
}

func (horario Horario) GetDia(dia DiaSemana.DiaSemana) map[HoraMinutos]*Clase {
	return horario.Clases[dia]
}
