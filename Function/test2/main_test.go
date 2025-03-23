package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got:=Add(7,5)
	want:=12
	if want!=got{
		log.Fatalf("error-want %d and got %d",want,got)
	}
}