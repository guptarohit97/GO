package main

import "testing"

func TestParadise(t *testing.T) {
	got:=paradise("Manali")
	want :="My idea of paradise is Manali"
	if want!=got{
		t.Fatalf("error - want %s and got %s",want,got)
	}
}