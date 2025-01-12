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

///////////////////////////////////////

func (horario Horario) GetClase(dia DiaSemana.DiaSemana, hora string) *Clase {
	time, err := NewHoraMinutosStr(hora)
	if err != nil {
		return nil
	}

	return horario.Clases[dia][*time]
}

///////////////////

func extraerClases(fileName string) (*[]Clase, error) {
	// Inicialización de variables
	var err error
	var clases []Clase
	var periodo *Periodo
	var dia DiaSemana.DiaSemana

	var aula string
	var grupo string
	var asignatura string

	// Abir y cerrar el archivo
	file, _ := os.Open(fileName)
	defer file.Close()

	// Expresiones regulares para buscar la información
	expAsignatura := regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`)
	expDia := regexp.MustCompile(`<div class=\"clase dia-(\d)\"`)
	expGrupo := regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-Za-z]|\d{1,2})</div>`)
	expAula := regexp.MustCompile(`<div>Aula:\s*(\d+)</div>`)
	expHorario := regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`)

	// Leer el archivo
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

		if asignatura != "" && dia != "" && grupo != "" && aula != "" && periodo != nil {
			clase := Clase{
				DiaSemana: dia,
				Periodo:   periodo,
				Aula:      aula,
				Grupo:     Grupo{Nombre: grupo, Asignatura: asignatura, Profesor: ""},
			}

			clases = append(clases, clase)

			// Reiniciar variables
			asignatura = ""
			dia = ""
			grupo = ""
			aula = ""
			periodo = nil
		}
	}

	if len(clases) == 0 {
		return nil, errors.New("no se han encontrado clases")
	} else {
		return &clases, nil
	}
}

///////////////////

func extraerProfesor(fileName string, clase *Clase) (*Clase, error) {
	var prof string
	var grupos []string
	var leer bool

	// Abrir el archivo
	file, _ := os.Open(fileName)
	defer file.Close()

	// Expresiones regulares para buscar la inforamción
	expProf := regexp.MustCompile(`<a href=\"https://www.ugr.es/personal/[^>]*\">([^<]+)</a>`)
	expGrupos := regexp.MustCompile(`Grupos?&nbsp;`)
	expNumGp := regexp.MustCompile(`([A-Z]|\d{1,2})(,\s*([A-Z]|\d{1,2}))*\s*(y\s*([A-Z]|\d{1,2}))?$`)

	// Leer el archivo
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
				grupo := matches[0]

				// Limpiar y dividir los grupos
				grupo = strings.ReplaceAll(grupo, " y ", ",")
				grupos = append(grupos, strings.Split(grupo, ",")...)

				for _, grupo := range grupos {
					if strings.TrimSpace(grupo) == clase.Grupo.Nombre {
						clase.Grupo.setProfesor(prof)
						return clase, nil
					}
				}

				// Reiniciar variables
				leer = false
				grupos = grupos[:0]
			}
		}
	}

	// Si no se encontró el grupo
	return nil, errors.New("no se ha encontrado profesor para el grupo buscado")
}

///////////////////

func NewHorarioFromClases(clases []Clase) Horario {
	horario := Horario{Clases: make(map[DiaSemana.DiaSemana]map[HoraMinutos]*Clase)}

	for _, clase := range clases {
		dia := clase.DiaSemana

		// Inicializar el mapa interno (por hora) si aún no existe
		if horario.Clases[dia] == nil {
			horario.Clases[dia] = make(map[HoraMinutos]*Clase)
		}

		// Insertar la clase en la hora correspondiente
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
