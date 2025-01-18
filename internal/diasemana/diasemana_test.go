package DiaSemana

import (
	"testing"
)

func TestDiaStr(t *testing.T) {
	tests := []struct {
		dia      DiaSemana
		expected string
	}{
		{Lunes, "1"},
		{Martes, "2"},
		{Miercoles, "3"},
		{Jueves, "4"},
		{Viernes, "5"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if res := DiaStr(test.dia); res != test.expected {
				t.Errorf("DiaStr(%v) = %v, want %v", test.dia, res, test.expected)
			}
		})
	}
}

func TestStrDia(t *testing.T) {
	tests := []struct {
		expected string
		dia      DiaSemana
	}{
		{"1", Lunes},
		{"2", Martes},
		{"3", Miercoles},
		{"4", Jueves},
		{"5", Viernes},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if res := StrDia(test.expected); res != test.dia {
				t.Errorf("StrDia(%v) = %v, want %v", test.expected, res, test.dia)
			}
		})
	}
}
