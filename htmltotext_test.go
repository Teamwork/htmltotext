package htmltotext

import (
	"fmt"
	"testing"
)

func TestToText(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{``, ``},
		{`Hello`, `Hello`},
		{`<span>Hello</span>`, `Hello`},
		{`<span>Hello<`, `Hello<`},
		{"<span>Hello\n\n\n</span>", "Hello\n\n\n"},
		{"<p>Hello\n\n</p><p>World</p>", "Hello\n\nWorld"},

		// This is actually an error; see
		// golang.org/x/net/html/token_test.go TestMaxBuffer()
		{`<ttttttttttttttttttttttttttt`, ``},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := ToText(tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}

func TestToDocument(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{``, ``},
		{`Hello`, `Hello`},
		{`<span>Hello</span>`, `Hello`},
		{`<span>Hello<`, `Hello<`},
		{"<span>Hello\n\n\n</span>", "Hello\n\n"},
		{"<p>Hello\n\n</p><p>World</p>", "Hello\n\nWorld\n"},
		{"<p>Hello<br></p>World<br>asd<br/>", "Hello\n\nWorld\nasd\n"},

		// This is actually an error; see
		// golang.org/x/net/html/token_test.go TestMaxBuffer()
		{`<ttttttttttttttttttttttttttt`, ``},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := ToDocument(tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}
