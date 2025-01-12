package models

type Grupo struct {
	Nombre     string
	Asignatura string
	Profesor   string
}

func (grupo *Grupo) setProfesor(prof string) {
	grupo.Profesor = prof
}

func NewGrupo(nombre, asignatura, profesor string) *Grupo {
	return &Grupo{Nombre: nombre, Asignatura: asignatura, Profesor: profesor}
}
