package DiaSemana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiaStr(t *testing.T) {
	assert.Equal(t, "1", DiaStr(Lunes), "Día erroneo")
	assert.Equal(t, "2", DiaStr(Martes), "Día erroneo")
	assert.Equal(t, "3", DiaStr(Miercoles), "Día erroneo")
	assert.Equal(t, "4", DiaStr(Jueves), "Día erroneo")
	assert.Equal(t, "5", DiaStr(Viernes), "Día erroneo")
	assert.Equal(t, "", DiaStr(""), "Día erroneo")
}

func TestStrDia(t *testing.T) {
	assert.Equal(t, Lunes, StrDia("1"), "Día erroneo")
	assert.Equal(t, Martes, StrDia("2"), "Día erroneo")
	assert.Equal(t, Miercoles, StrDia("3"), "Día erroneo")
	assert.Equal(t, Jueves, StrDia("4"), "Día erroneo")
	assert.Equal(t, Viernes, StrDia("5"), "Día erroneo")
	assert.Equal(t, DiaSemana(""), StrDia(""), "Día erroneo")
}
