package main

import "testing"

func TestMain(t *testing.T) {
	total := Somar(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: o Resultado %d,"+
			"não é o valor esperado", total)
	}
}
