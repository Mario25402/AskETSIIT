package DiaSemana

type DiaSemana int

const (
	Lunes     DiaSemana = iota // 0
	Martes                     // 1
	Miercoles                  // 2
	Jueves                     // 3
	Viernes                    // 4
)

func (d DiaSemana) String() string {
	switch d {
	case Lunes:
		return "Lunes"
	case Martes:
		return "Martes"
	case Miercoles:
		return "Miércoles"
	case Jueves:
		return "Jueves"
	case Viernes:
		return "Viernes"
	default:
		return "Día desconocido"
	}
}
