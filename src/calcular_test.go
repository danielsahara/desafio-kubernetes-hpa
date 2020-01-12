package main

import "testing"

func TestCalcular(t *testing.T){
	got := calcular()
	want := 2.499962871007386e+11

	if got != want{
		t.Errorf("calcular got: %v want: %v", got, want)
	}
}