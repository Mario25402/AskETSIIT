package DiaSemana

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miércoles"
	Jueves    DiaSemana = "Jueves "
	Viernes   DiaSemana = "Viernes"
)

func DiaStr(dia DiaSemana) string {
	switch dia {
	case Lunes:
		return "1"
	case Martes:
		return "2"
	case Miercoles:
		return "3"
	case Jueves:
		return "4"
	case Viernes:
		return "5"
	}
	return ""
}

func StrDia(dia string) DiaSemana {
	switch dia {
	case "1":
		return Lunes
	case "2":
		return Martes
	case "3":
		return Miercoles
	case "4":
		return Jueves
	case "5":
		return Viernes
	}
	return ""
}
