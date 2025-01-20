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

func comprobarAdicionClase(clases *[]Clase, asignatura, grupo, aula *string, dia *DiaSemana.DiaSemana, periodo *Periodo) {
	if *asignatura != "" && *dia != "" && *grupo != "" && *aula != "" && periodo != nil {
		clase := Clase{
			DiaSemana: *dia,
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

func procesadorClase(dia *DiaSemana.DiaSemana, asignatura, grupo, aula *string, periodo **Periodo) map[*regexp.Regexp]func([]string) error {
	return map[*regexp.Regexp]func([]string) error{
		regexp.MustCompile(`<h1 class=\"page-title\">([^<]+)</h1>`): func(matches []string) error {
			*asignatura = strings.TrimSpace(matches[1])
			return nil
		},
		regexp.MustCompile(`<div class=\"clase dia-(\d)\"`): func(matches []string) error {
			*dia = DiaSemana.DiaSemana(matches[1])
			if *dia == "" {
				return errors.New("dia no válido")
			}
			return nil
		},
		regexp.MustCompile(`<div class=\"grupo\"><span>Grupo:</span>\s*([A-Za-z]|\d{1,2})</div>`): func(matches []string) error {
			*grupo = matches[1]
			return nil
		},
		regexp.MustCompile(`<div>Aula:\s*(\d+)</div>`): func(matches []string) error {
			*aula = matches[1]
			return nil
		},
		regexp.MustCompile(`<div>Horario:\s*De\s*(\d{2}:\d{2})\s*a\s*(\d{2}:\d{2})</div>`): func(matches []string) error {
			var err error
			*periodo, err = NewPeriodoStr(matches[1], matches[2])
			return err
		},
	}
}

func extraerClases(fileName string) (*[]Clase, error) {
	var clases []Clase
	var periodo *Periodo
	var dia DiaSemana.DiaSemana
	var aula, grupo, asignatura string

	file, _ := os.Open(fileName)
	defer file.Close()

	procesador := procesadorClase(&dia, &asignatura, &grupo, &aula, &periodo)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		for exp, proc := range procesador {
			if matches := exp.FindStringSubmatch(linea); matches != nil {
				if err := proc(matches); err != nil {
					return nil, err
				}
			}
		}

		comprobarAdicionClase(&clases, &asignatura, &grupo, &aula, &dia, periodo)
	}

	if len(clases) == 0 {
		return nil, errors.New("no se han encontrado clases")
	} else {
		return &clases, nil
	}
}

func establecerProfesor(clase *Clase, prof, cadena *string) {
	if clase != nil && prof != nil && cadena != nil {
		var grupos []string

		*cadena = strings.ReplaceAll(*cadena, " y ", ",")
		grupos = append(grupos, strings.Split(*cadena, ",")...)

		for _, grupo := range grupos {
			if strings.TrimSpace(grupo) == clase.Grupo.Nombre {
				clase.Grupo.setProfesor(*prof)
				break
			}
		}
	}
}

func procesadorProfesor(prof *string) map[*regexp.Regexp]func([]string) (clase *Clase, profesor, cadena *string) {
	var leer bool

	return map[*regexp.Regexp]func([]string) (clase *Clase, profesor, cadena *string){
		regexp.MustCompile(`<a href=\"https://www.ugr.es/personal/[^>]*\">([^<]+)</a>`): func(matches []string) (clase *Clase, profesor, cadena *string) {
			*prof = strings.TrimSpace(matches[1])
			return nil, nil, nil
		},

		regexp.MustCompile(`Grupos?&nbsp;`): func(matches []string) (clase *Clase, profesor, cadena *string) {
			leer = true
			return nil, nil, nil
		},

		regexp.MustCompile(`([A-Z]|\d{1,2})(,\s*([A-Z]|\d{1,2}))*\s*(y\s*([A-Z]|\d{1,2}))?$`): func(matches []string) (clase *Clase, profesor, cadena *string) {
			if leer {
				leer = false
				return clase, prof, &matches[0]
			}

			return nil, nil, nil
		},
	}
}

func extraerProfesor(clase *Clase, fileName string) error {
	var prof string
	file, _ := os.Open(fileName)
	defer file.Close()

	procesador := procesadorProfesor(&prof)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		for expReg, procesado := range procesador {
			if matches := expReg.FindStringSubmatch(linea); matches != nil {
				clase, prof, match := procesado(matches)
				establecerProfesor(clase, prof, match)
				break
			}
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
