package payload

import (
	"errors"
	"testing"
)

func TestParseFrontmatterAcceptsLFAndCRLF(t *testing.T) {
	cases := []struct {
		name    string
		data    string
		newline string
		body    string
	}{
		{name: "lf", data: "---\nname: demo\ndescription: d\n---\nBody text\n", newline: "\n", body: "Body text\n"},
		{name: "crlf", data: "---\r\nname: demo\r\ndescription: d\r\n---\r\nBody text\r\n", newline: "\r\n", body: "Body text\r\n"},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			document, err := ParseFrontmatter([]byte(testCase.data))
			if err != nil {
				t.Fatalf("ParseFrontmatter: %v", err)
			}
			if document.Newline != testCase.newline {
				t.Fatalf("newline = %q, want %q", document.Newline, testCase.newline)
			}
			lines := document.Lines()
			if len(lines) != 2 || lines[0] != "name: demo" || lines[1] != "description: d" {
				t.Fatalf("unexpected frontmatter lines %q", lines)
			}
			if got := testCase.data[document.BodyStart:]; got != testCase.body {
				t.Fatalf("body = %q, want %q", got, testCase.body)
			}
			if got := testCase.data[document.CloseStart:document.BodyStart]; got != testCase.newline+"---"+testCase.newline {
				t.Fatalf("closing delimiter = %q", got)
			}
		})
	}
}

func TestParseFrontmatterRejectsMalformedDocuments(t *testing.T) {
	if _, err := ParseFrontmatter([]byte("name: demo\n")); !errors.Is(err, ErrMissingOpeningFrontmatter) {
		t.Fatalf("missing opening delimiter error = %v", err)
	}
	if _, err := ParseFrontmatter([]byte("---\nname: demo\n")); !errors.Is(err, ErrMissingClosingFrontmatter) {
		t.Fatalf("missing closing delimiter error = %v", err)
	}
}

func TestValidateFrontmatterAcceptsCRLFSkills(t *testing.T) {
	data := []byte("---\r\nname: demo\r\ndescription: d\r\n---\r\nBody\r\n")
	if err := validateFrontmatter("demo", data); err != nil {
		t.Fatalf("validateFrontmatter: %v", err)
	}
	if err := validateFrontmatter("other", data); err == nil {
		t.Fatal("expected mismatched name to be rejected")
	}
}
