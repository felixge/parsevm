package vm

import (
	"testing"

	"github.com/felixge/goldy"
)

var gc = goldy.DefaultConfig()

func TestProgram_String(t *testing.T) {
	tests := []struct {
		Name    string
		Program Program
	}{
		{
			Name: "empty",
		},
		{
			Name:    "one_string",
			Program: String("hello world"),
		},
		{
			Name:    "two_strings",
			Program: Concat(String("hello"), String("world")),
		},
		{
			Name: "eleven_strings",
			Program: Concat(
				String("0"),
				String("1"),
				String("2"),
				String("3"),
				String("4"),
				String("5"),
				String("6"),
				String("7"),
				String("8"),
				String("9"),
				String("10"),
			),
		},
	}

	gf := gc.GoldenFixtures("program_string")
	for _, test := range tests {
		got := test.Program.String()
		gf.Add([]byte(got), test.Name+".txt")
	}
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}
