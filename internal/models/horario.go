package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
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
