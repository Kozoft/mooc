package tasks

import (
	"testing"
)

func Test11(t *testing.T) {
	inout := map[string]string{
		"[]":                    "Success",
		"{}[]":                  "Success",
		"[()]":                  "Success",
		"(())":                  "Success",
		"{[]}()":                "Success",
		"{":                     "1",
		"{[}":                   "3",
		"foo(bar);":             "Success",
		"foo(bar[i);":           "10",
		"()({}":                 "3",
		"[]([]":                 "3",
		"((({[]})":              "2",
		"{}([]":                 "3",
		"(we, {all[live]} ,in)": "Success",
		"(a{wellow[submarine]}": "1",
		"yellowsubmarine]]]":    "16",
		"{{[()]}":               "1",
		"([](){([])})":          "Success",
		"()[]}":                 "5",
		"{{[()]]":               "7",
		"":                      "Success",
	}
	for in, out := range inout {
		// fmt.Println("in:", in, "out:", out)
		msg := Task11(in)
		if msg != out {
			t.Errorf("%s: %s, not %s", in, msg, out)
		} else {
			t.Logf("%s", in)
		}
	}
}
