package main

import "testing"

func TestShowWelcome(t *testing.T) {
	got := ShowWelcome()
	want := "Bem vindo ao programa para turistas"

	if got != want {
		t.Errorf("Erro \n got: %v, \n want: %v", got, want)
	}
}
