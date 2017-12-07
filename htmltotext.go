// Package htmltotext converts HTML to plain text.
package htmltotext

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ToText removes HTML tags. In general you only want to use this for short
// strings (e.g. a line or less). use ToDocument() for longer text with
// paragraphs.
//
// For example:
//
//    <b>&iexcl;Hi!</b> <script>...</script>
//
// becomes:
//
//    &iexcl;Hi!
func ToText(html string) string {
	if strings.TrimSpace(html) == "" {
		return html
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		// Don't return the original HTML, since that might introduce security
		// risks if we rely on it for XSS escaping.
		return fmt.Sprintf("could not parse document: %v", err.Error())
	}

	return doc.Text()
}

var (
	reParagraphs = regexp.MustCompile(`(\<\/?((br\s?\/?)|(p))\>)`)
	reWhitespace = regexp.MustCompile(`(?m)^[ \t]*`)
	reNewlines   = regexp.MustCompile(`\n{3,}`)
)

// ToDocument removes HTML tags from a document and attempts to format it so
// it's relatively readable.
func ToDocument(html string) string {
	// Convert any newline type
	plaintext := reParagraphs.ReplaceAllString(html, "\r\n")

	// Remove all tags but leave nice
	plaintext = ToText(plaintext)

	// Remove leading whitespace from lines
	plaintext = reWhitespace.ReplaceAllString(plaintext, "")

	// We never want more than two consecutive newlines.
	plaintext = reNewlines.ReplaceAllString(plaintext, "\n\n")

	return plaintext
}
