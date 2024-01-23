package main

import "testing"

func TestSoma(t *testing.T) {
	total := Somar(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: o Resultado %d,"+
			"não é o valor esperado", total)
	}
}

func TestMain(t *testing.T) {
}

/**
	func TestSub(t *testing.T) {
	total := Subtracao(15, 15)
	if total != 0 {

		t.Errorf("O valor esperado era de  0,"+
			"porém o valor recebido é de  %d", total)
	}
}

func TestMult(t *testing.T) {
	total := Multi(15, 2)
	if total != 30 {

		t.Errorf("O valor esperado era de  0,"+
			"porém o valor recebido é de  %d", total)
	}
}
**/
