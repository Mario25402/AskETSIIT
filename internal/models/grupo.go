package models

type Grupo struct {
	Nombre     string
	Asignatura string
	Profesor   string
}

func (grupo *Grupo) setProfesor(prof string) {
	grupo.Profesor = prof
}
