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

func comprobarAdicionClase(clases *[]Clase, asignatura, dia, grupo, aula *string, periodo *Periodo) {
	if *asignatura != "" && *dia != "" && *grupo != "" && *aula != "" && periodo != nil {
		clase := Clase{
			DiaSemana: DiaSemana.StrDia(*dia),
			Periodo:   periodo,
			Aula:      *aula,
			Grupo:     Grupo{Nombre: *grupo, Asignatura: *asignatura, Profesor: ""},
		}

		*clases = append(*clases, clase)

		*asignatura = ""
		*dia = ""
		*grupo = ""
		*aula = ""
		periodo = nil
	}
}

func extraerClases(fileName string) (*[]Clase, error) {
	var err error
	var clases []Clase
	var periodo *Periodo
	var dia DiaSemana.DiaSemana

	var aula string
	var grupo string
	var asignatura string

	file, _ := os.Open(fileName)
	defer file.Close()

	expAsignatura := regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`)
	expDia := regexp.MustCompile(`<div class=\"clase dia-(\d)\"`)
	expGrupo := regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-Za-z]|\d{1,2})</div>`)
	expAula := regexp.MustCompile(`<div>Aula:\s*(\d+)</div>`)
	expHorario := regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		if matches := expAsignatura.FindStringSubmatch(linea); matches != nil {
			asignatura = strings.TrimSpace(matches[1])
			continue
		} else if matches := expDia.FindStringSubmatch(linea); matches != nil {
			dia = DiaSemana.StrDia(matches[1])
			if dia == "" {
				return nil, errors.New("dia no válido")
			}
			continue
		} else if matches := expGrupo.FindStringSubmatch(linea); matches != nil {
			grupo = matches[1]
			continue
		} else if matches := expAula.FindStringSubmatch(linea); matches != nil {
			aula = matches[1]
			continue
		} else if matches := expHorario.FindStringSubmatch(linea); matches != nil {
			periodo, err = NewPeriodoStr(matches[1], matches[2])
			if err != nil {
				return nil, err
			}
		}

		diaStr := DiaSemana.DiaStr(dia)
		comprobarAdicionClase(&clases, &asignatura, &diaStr, &grupo, &aula, periodo)

	}

	if len(clases) == 0 {
		return nil, errors.New("no se han encontrado clases")
	} else {
		return &clases, nil
	}
}

func establecerProfesor(clase *Clase, prof, cadena string) *Clase {
	var grupos []string

	cadena = strings.ReplaceAll(cadena, " y ", ",")
	grupos = append(grupos, strings.Split(cadena, ",")...)

	for _, grupo := range grupos {
		if strings.TrimSpace(grupo) == clase.Grupo.Nombre {
			clase.Grupo.setProfesor(prof)
			return clase
		}
	}

	return nil
}

func extraerProfesor(fileName string, clase *Clase) (*Clase, error) {
	var prof string
	var leer bool

	file, _ := os.Open(fileName)
	defer file.Close()

	expProf := regexp.MustCompile(`<a href=\"https://www.ugr.es/personal/[^>]*\">([^<]+)</a>`)
	expGrupos := regexp.MustCompile(`Grupos?&nbsp;`)
	expNumGp := regexp.MustCompile(`([A-Z]|\d{1,2})(,\s*([A-Z]|\d{1,2}))*\s*(y\s*([A-Z]|\d{1,2}))?$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		if matches := expProf.FindStringSubmatch(linea); matches != nil {
			prof = strings.TrimSpace(matches[1])
			continue
		}

		if matches := expGrupos.FindStringSubmatch(linea); matches != nil {
			leer = true
			continue
		}

		if leer {
			if matches := expNumGp.FindStringSubmatch(linea); matches != nil {
				clase := establecerProfesor(clase, prof, matches[0])
				if clase != nil {
					return clase, nil
				}

				leer = false
			}
		}
	}

	return nil, errors.New("no se ha encontrado profesor para el grupo buscado")
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
		clasePtr, err := extraerProfesor(file, &clase)
		if err != nil {
			return nil
		}

		(*clases)[i] = *clasePtr
	}

	horario := NewHorarioFromClases(*clases)
	return &horario
}
