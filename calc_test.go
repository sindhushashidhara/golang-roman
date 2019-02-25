package roman_test

import (
	"testing"
)


func TestToRoman(t *testing.T) {
	uu := []struct {
		arabic int
		glyph string
	} {
		{1, "I"} ,
		{2, "II"},
		{3, "III"},
		{5, "V"},
		{10, "X"},
		{9, "IX"},
		{11, "XI"},
	}

	for _, u := range uu {
		g := toRoman(u.arabic)
		e := u.glyph
		if e != g {
			t.Fatalf("Expecting %s but got %s", e, g)
		}
	}
}

func TestToRoman1(t *testing.T){
	output := toRoman(1)
	expected := "I"

	if expected != output {
		t.Fatal("Expecting %s but got %s", expected, output)
	}
}