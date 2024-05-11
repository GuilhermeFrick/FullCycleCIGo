package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(20, 20)
	esperado := 40
	if resultado != esperado {
		t.Errorf("Soma(20, 20) = %d; esperado %d", resultado, esperado)
	}
}
