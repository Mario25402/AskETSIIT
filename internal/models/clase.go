package models

import (
	DiaSemana "askETSIIT/internal/diasemana"
	"fmt"
)

type Clase struct {
	Dia     DiaSemana.DiaSemana
	Periodo *Periodo
	Aula    string
	Grupo   Grupo
}

func (c Clase) String() string {
	return fmt.Sprintf("- %s %s:  %s (%s)", c.Dia, c.Periodo, c.Grupo.Asignatura, c.Aula)
}
