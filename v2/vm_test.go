package vm

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
)

func TestVM_Valid(t *testing.T) {
	type testProgram struct {
		Name    string
		Program Program
	}

	// Each test is a set of programs recognizing the same set of inputs.
	tests := []struct {
		Name     string
		Programs []testProgram
		Inputs   []string
	}{
		{
			"Concat",
			[]testProgram{
				{"abc1", String("abc")},
				{"abc2a", Concat(String("ab"), String("c"))},
				{"abc2b", Concat(String("a"), String("bc"))},
				{"abc3", Concat(String("a"), String("b"), String("c"))},
			},
			[]string{
				"hello world",
				"world hello",
				"helloworld",
				"hhello world",
				"",
				"hello",
			},
		},

		{
			"ZeroOrMore",
			[]testProgram{
				{"abc1", ZeroOrMore(String("abc"))},
				{"abc2", ZeroOrMore(Concat(String("a"), String("b"), String("c")))},
			},
			[]string{
				"",
				"abc",
				"abcabc",
				"abcabcabc",
				"ab",
				"def",
			},
		},

		{
			"ZeroOrOne",
			[]testProgram{
				{"abc1", ZeroOrOne(String("abc"))},
				{"abc2", ZeroOrOne(Concat(String("a"), String("b"), String("c")))},
			},
			[]string{
				"",
				"abc",
				"abcabc",
				"abcabcabc",
				"ab",
				"def",
			},
		},

		{
			"OneOrMore",
			[]testProgram{
				{"abc1", OneOrMore(String("abc"))},
				{"abc2", OneOrMore(Concat(String("a"), String("b"), String("c")))},
			},
			[]string{
				"",
				"abc",
				"abcabc",
				"abcabcabc",
				"ab",
				"def",
			},
		},

		{
			"Alt",
			[]testProgram{
				{"abcdefghj1", Alt(String("abc"), String("def"), String("ghi"))},
				{"abcdefghi2", Alt(
					Concat(String("a"), String("bc")),
					Concat(String("de"), String("f")),
					Concat(String("g"), String("h"), String("i")),
				)},
			},
			[]string{
				"",
				"abc",
				"def",
				"ghi",
				"abcabc",
				"adg",
			},
		},

		{
			"Repeat",
			[]testProgram{
				{"abc1", Repeat(1, 3, String("abc"))},
				{"abc3", Repeat(1, 3, Concat(String("a"), String("b"), String("c")))},
			},
			[]string{
				"",
				"ab",
				"abc",
				"abcabc",
				"abcabcabc",
				"abcabcabcabc",
				"adcd",
			},
		},
	}

	gf := gc.GoldenFixtures("vm_valid")
	for _, test := range tests {
		buf := bytes.NewBuffer(nil)
		fmt.Fprintf(buf, "# %s\n\n", test.Name)
		fmt.Fprintf(buf, "## Programs\n\n")

		for _, program := range test.Programs {
			fmt.Fprintf(buf, "### Program %s\n\n```\n%s```\n\n", program.Name, program.Program)
		}

		fmt.Fprintf(buf, "## Inputs\n\n")

		for i, input := range test.Inputs {
			var table [][]string
			table = append(table, []string{"Program", "Valid", "n", "err", "ops", "forks"})

			inputID := fmt.Sprintf("%d", i+1)
			t.Run(inputID, func(t *testing.T) {
				fmt.Fprintf(buf, "## Input %s: %q\n\n", inputID, input)

				for _, program := range test.Programs {
					t.Run(program.Name, func(t *testing.T) {
						vm := NewVM(program.Program)
						n, err := vm.Write([]byte(input))
						valid := vm.Valid()
						stats := vm.Stats()

						table = append(table, []string{
							program.Name,
							fmt.Sprintf("%t", valid),
							fmt.Sprintf("%d", n),
							fmt.Sprintf("%v", err),
							fmt.Sprintf("%d", stats.Ops),
							fmt.Sprintf("%d", stats.Forks),
						})
					})
				}
				fmt.Fprintf(buf, "%s\n", markdownTable(table))
			})

		}
		gf.Add(buf.Bytes(), test.Name+".md")
	}
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}

func TestVM_Complexity(t *testing.T) {
	buf := &bytes.Buffer{}
	var table [][]string
	table = append(table, []string{"n", "err", "ops", "forks"})

	max := 30
	for n := 1; n <= max; n++ {
		var p Program
		input := strings.Repeat("a", n)
		for i := 0; i < n; i++ {
			p = Concat(p, ZeroOrOne(String("a")))
		}
		for i := 0; i < n; i++ {
			p = Concat(p, String("a"))
		}

		if n == max {
			fmt.Fprintf(buf, "# Program\n\n```\n%s```\n\n", p)
		}

		v := NewVM(p)
		n, err := v.Write([]byte(input))
		if !v.Valid() {
			t.Fatal("invalid")
		}
		stats := v.Stats()

		table = append(table, []string{
			fmt.Sprintf("%d", n),
			fmt.Sprintf("%v", err),
			fmt.Sprintf("%d", stats.Ops),
			fmt.Sprintf("%d", stats.Forks),
		})
	}

	fmt.Fprintf(buf, "# Complexities\n\n%s\n", markdownTable(table))
	gc.GoldenFixture(buf.Bytes(), "vm_complexity.md")
}

func markdownTable(data [][]string) string {
	buf := &bytes.Buffer{}
	table := tablewriter.NewWriter(buf)
	table.SetHeader(data[0])
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data[1:])
	table.Render()
	return buf.String()
}
