// nolint: lll
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
		{`<span>Hello<!-- world --> test</span>`, `Hello test`},

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

func TestToLine(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{``, ``},
		{`Hello`, `Hello`},
		{`<span>Hello</span>`, `Hello`},
		{`<span>Hello<`, `Hello<`},
		{"<span>Hello\n\n\n</span>", "Hello"},
		{"<p>Hello\n\n</p><p>World</p>", "Hello World"},
		{`<span>Hello<!-- world --> test</span>`, `Hello test`},
		{`<html><body style=""><div id="bloop_customfont" style="">Hey Pat,</div><div id="bloop_customfont" style="">This is a test with newlines.</div><div id="bloop_customfont" style="">Should be fixed.</div></body></html>`, `Hey Pat, This is a test with newlines. Should be fixed.`},

		// This is actually an error; see
		// golang.org/x/net/html/token_test.go TestMaxBuffer()
		{`<ttttttttttttttttttttttttttt`, ``},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := ToLine(tc.in)
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
		{`<span>Hello<!-- world --> test</span>`, `Hello test`},

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
