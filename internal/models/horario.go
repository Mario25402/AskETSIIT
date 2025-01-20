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
	time, err := NewHoraMinutosStr(hora)
	if err != nil {
		return nil
	}

	return horario.Clases[dia][*time]
}

func comprobarAdicionClase(clases *[]Clase, asignatura, grupo, aula string, dia DiaSemana.DiaSemana, periodo *Periodo) {
	if asignatura != "" && dia != "" && grupo != "" && aula != "" && periodo != nil {
		clase := Clase{
			DiaSemana: dia,
			Periodo:   periodo,
			Aula:      aula,
			Grupo:     Grupo{Nombre: grupo, Asignatura: asignatura, Profesor: ""},
		}

		*clases = append(*clases, clase)
	}
}

func procesarAsignatura(linea string) string {
	expAsignatura := regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`)

	if matches := expAsignatura.FindStringSubmatch(linea); matches != nil {
		return strings.TrimSpace(matches[1])
	}

	return ""
}

func procesarDia(linea string) DiaSemana.DiaSemana {
	expDia := regexp.MustCompile(`<div class=\"clase dia-(\d)\"`)

	if matches := expDia.FindStringSubmatch(linea); matches != nil {
		if "1" < matches[1] && matches[1] > "5" {
			return ""
		}

		return DiaSemana.DiaSemana(matches[1])
	}

	return ""
}

func procesarGrupo(linea string) string {
	expGrupo := regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-Za-z]|\d{1,2})</div>`)

	if matches := expGrupo.FindStringSubmatch(linea); matches != nil {
		return matches[1]
	}

	return ""
}

func procesarAula(linea string) string {
	expAula := regexp.MustCompile(`<div>Aula:\s*(\d+)</div>`)

	if matches := expAula.FindStringSubmatch(linea); matches != nil {
		return matches[1]
	}

	return ""
}

func procesarPeriodo(linea string) *Periodo {
	expPeriodo := regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`)

	if matches := expPeriodo.FindStringSubmatch(linea); matches != nil {
		newPeriodo, err := NewPeriodoStr(matches[1], matches[2])
		if err == nil {
			return newPeriodo
		}

		return nil
	}

	return nil
}

func extraerClases(fileName string) (*[]Clase, error) {
	var clases []Clase

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		dia := procesarDia(linea)
		aula := procesarAula(linea)
		grupo := procesarGrupo(linea)
		periodo := procesarPeriodo(linea)
		asignatura := procesarAsignatura(linea)

		comprobarAdicionClase(&clases, asignatura, grupo, aula, dia, periodo)

		dia = ""
		aula = ""
		grupo = ""
		periodo = nil
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

	return errors.New("no se ha encontrado profesor")
}

func NewHorarioFromClases(clases []Clase) Horario {
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

	for i, clase := range *clases {
		var _ *Clase = &clase
		err := extraerProfesor(&clase, file)
		if err != nil {
			return nil
		}

		(*clases)[i] = clase
	}

	horario := NewHorarioFromClases(*clases)
	return &horario
}
