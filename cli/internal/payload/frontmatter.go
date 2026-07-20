package payload

import (
	"errors"
	"strings"
)

// FrontmatterDocument describes the frontmatter block that must open a
// SKILL.md file. Offsets are byte positions into the original document so
// callers can splice metadata into the frontmatter or body without
// re-deriving the delimiter layout.
type FrontmatterDocument struct {
	Frontmatter string // frontmatter body between the delimiters
	Newline     string // line ending used by the document ("\n" or "\r\n")
	CloseStart  int    // offset of the newline that begins the closing delimiter
	BodyStart   int    // offset of the first byte after the closing delimiter
}

var (
	ErrMissingOpeningFrontmatter = errors.New("SKILL.md has no opening frontmatter delimiter")
	ErrMissingClosingFrontmatter = errors.New("SKILL.md has no closing frontmatter delimiter")
)

// ParseFrontmatter splits a SKILL.md document into its frontmatter block and
// body. Both LF and CRLF line endings are accepted; the document's own ending
// is reported so insertions can match it.
func ParseFrontmatter(data []byte) (FrontmatterDocument, error) {
	text := string(data)
	newline := "\n"
	if strings.HasPrefix(text, "---\r\n") {
		newline = "\r\n"
	} else if !strings.HasPrefix(text, "---\n") {
		return FrontmatterDocument{}, ErrMissingOpeningFrontmatter
	}
	open := len("---") + len(newline)
	closing := newline + "---" + newline
	end := strings.Index(text[open:], closing)
	if end < 0 {
		return FrontmatterDocument{}, ErrMissingClosingFrontmatter
	}
	return FrontmatterDocument{
		Frontmatter: text[open : open+end],
		Newline:     newline,
		CloseStart:  open + end,
		BodyStart:   open + end + len(closing),
	}, nil
}

// Lines returns the frontmatter body split into lines without line endings.
func (document FrontmatterDocument) Lines() []string {
	return strings.Split(document.Frontmatter, document.Newline)
}
