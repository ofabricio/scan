package scan

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

type Bytes []byte

// Space advances the scanner if the next rune is a space.
func (s *Bytes) Space() bool {
	return s.MatchFunc(unicode.IsSpace)
}

// Spaces advances the scanner while the next rune is a space.
func (s *Bytes) Spaces() bool {
	return s.WhileFunc(unicode.IsSpace)
}

// Until advances the scanner until the next bytes are equal to v.
// Returns true if it advanced.
func (s *Bytes) Until(v string) bool {
	m := false
	for s.More() && !s.Equal(v) {
		m = s.Next()
	}
	return m
}

// UntilChar advances the scanner until the next rune is equal to any rune in v.
// Returns true if it advanced.
func (s *Bytes) UntilChar(v string) bool {
	m := false
	for s.More() && !s.EqualChar(v) {
		m = s.Next()
	}
	return m
}

// UntilFunc advances the scanner until the next rune satisfies the function v.
// Returns true if it advanced.
func (s *Bytes) UntilFunc(v func(rune) bool) bool {
	m := false
	for s.More() && !s.EqualFunc(v) {
		m = s.Next()
	}
	return m
}

// While advances the scanner while the next bytes are equal to v.
// Returns true if it advanced.
func (s *Bytes) While(v string) bool {
	m := false
	for s.Match(v) {
		m = true
	}
	return m
}

// While advances the scanner while the next bytes are equal to v
// ignoring case. Returns true if it advanced.
func (s *Bytes) WhileFold(v string) bool {
	m := false
	for s.MatchFold(v) {
		m = true
	}
	return m
}

// WhileChar advances the scanner while the next rune is equal to
// any rune in v. Returns true if it advanced.
func (s *Bytes) WhileChar(v string) bool {
	m := false
	for s.MatchChar(v) {
		m = true
	}
	return m
}

// WhileFunc advances the scanner while the next rune satisfies
// the function v. Returns true if it advanced.
func (s *Bytes) WhileFunc(v func(rune) bool) bool {
	m := false
	for s.MatchFunc(v) {
		m = true
	}
	return m
}

// Match tells if the next bytes are equal to v.
// Advances the scanner if true.
func (s *Bytes) Match(v string) bool {
	if s.Equal(v) {
		return s.Advance(len(v))
	}
	return false
}

// MatchFold tells if the next bytes are equal to v
// ignoring case. Advances the scanner if true.
func (s *Bytes) MatchFold(v string) bool {
	if s.EqualFold(v) {
		return s.Advance(len(v))
	}
	return false
}

// MatchChar tells if the next rune is equal to any rune in v.
// Advances the scanner if true.
func (s *Bytes) MatchChar(v string) bool {
	if s.EqualChar(v) {
		return s.Next()
	}
	return false
}

// MatchFunc tells if the next rune satisfies the function v.
// Advances the scanner if true.
func (s *Bytes) MatchFunc(v func(rune) bool) bool {
	if s.EqualFunc(v) {
		return s.Next()
	}
	return false
}

// Equal tells if the next bytes are equal to v.
func (s Bytes) Equal(v string) bool {
	return bytes.HasPrefix(s, []byte(v))
}

// EqualFold tells if the next bytes are equal to v ignoring case.
func (s Bytes) EqualFold(v string) bool {
	return bytes.EqualFold(s.Peek(len(v)), []byte(v))
}

// EqualChar tells if the next rune is equal to any rune in or.
func (s Bytes) EqualChar(or string) bool {
	for _, v := range or {
		if s.Rune() == v {
			return true
		}
	}
	return false
}

// EqualFunc tells if the next rune satisfies the fuction v.
func (s Bytes) EqualFunc(v func(rune) bool) bool {
	return v(s.Rune())
}

// Take returns the next n bytes advancing the scanner by n bytes.
func (s *Bytes) Take(n int) Bytes {
	v := s.Peek(n)
	s.Advance(len(v))
	return v
}

// Peek returns the next n bytes without advancing the scanner.
func (s Bytes) Peek(n int) Bytes {
	if len(s) < n {
		return s
	}
	return s[:n]
}

// Rune returns the current rune.
// Returns [utf8.RuneError] if the scanner is empty
// or has invalid utf8.
func (s Bytes) Rune() rune {
	r, _ := utf8.DecodeRune(s)
	return r
}

// Byte returns the current byte.
// Returns 0 if the scanner is empty.
func (s Bytes) Byte() byte {
	if s.Empty() {
		return 0
	}
	return s[0]
}

// Text returns the text between the mark m
// and the current scanner position.
func (s Bytes) Text(m Bytes) string {
	return s.Delta(m).String()
}

// Delta returns the difference between the mark m
// and the current scanner position.
func (s Bytes) Delta(m Bytes) Bytes {
	n := max(0, len(m)-len(s))
	return m[:n]
}

// Next advances the scanner by one rune.
// Return true if it advanced.
func (s *Bytes) Next() bool {
	return s.AdvanceRune(s.Rune())
}

// Advance advances the scanner by the size of the rune r.
// Return true if it advanced.
func (s *Bytes) AdvanceRune(r rune) bool {
	return s.Advance(utf8.RuneLen(r))
}

// Advance advances the scanner by n bytes.
// Return true if it advanced.
func (s *Bytes) Advance(n int) bool {
	if s.Len() < n {
		return false
	}
	*s = (*s)[n:]
	return true
}

// String returns the scanner as a string.
func (s Bytes) String() string {
	return string(s)
}

// Bytes returns the scanner as a byte slice.
func (s Bytes) Bytes() []byte {
	return s
}

// Mark returns a mark to the current position.
func (s Bytes) Mark() Bytes {
	return s
}

// Move moves the scanner current position to the mark m.
func (s *Bytes) Move(m Bytes) {
	*s = m
}

// More tells if there are more bytes to scan.
// This is the opposite of [Empty].
func (s Bytes) More() bool {
	return len(s) > 0
}

// Empty tells if there are no more bytes to scan.
// This is the opposite of [More].
func (s Bytes) Empty() bool {
	return len(s) == 0
}

// Len returns the number of bytes left to scan.
func (s Bytes) Len() int {
	return len(s)
}
