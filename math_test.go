package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(20, 15)

	if total != 35 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 35)
	}
}
