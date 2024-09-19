package scan

import (
	"fmt"
	"unicode"
)

func ExampleBytes_TextWith() {

	s := Bytes("abc")

	Word := func(s *Bytes) bool {
		return s.WhileFunc(unicode.IsLetter)
	}

	var v string
	fmt.Println(s.TextWith(&v, Word))
	fmt.Println(v)
	fmt.Println(s.TextWith(&v, Word))

	// Output:
	// true
	// abc
	// false
	//
}

func ExampleBytes_Edge() {

	s := Bytes("a\nb\r\nc")

	fmt.Println(s.Edge())
	s.Match("a")
	fmt.Println(s.Edge())
	s.Match("\nb")
	fmt.Println(s.Edge())
	s.Match("\r\nc")
	fmt.Println(s.Edge())

	// Output:
	// false
	// true
	// true
	// true
}

func ExampleBytes_Spaces() {

	s := Bytes("  ")

	fmt.Println(s.Spaces())
	fmt.Println(s.String() == "")
	fmt.Println(s.Spaces())

	// Output:
	// true
	// true
	// false
}

func ExampleBytes_Space() {

	s := Bytes(" ")

	fmt.Println(s.Space())
	fmt.Println(s.String() == "")
	fmt.Println(s.Space())

	// Output:
	// true
	// true
	// false
}

func ExampleBytes_Until() {

	s := Bytes("abc..def..")

	fmt.Println(s.Until(".."))
	fmt.Println(s)
	fmt.Println(s.Until(".."))
	fmt.Println(s)

	// Output:
	// true
	// ..def..
	// false
	// ..def..
}

func ExampleBytes_UntilChar() {

	s := Bytes("abc.def")

	fmt.Println(s.UntilChar(".,"))
	fmt.Println(s)

	s = Bytes("abc,def")

	fmt.Println(s.UntilChar(".,"))
	fmt.Println(s)

	fmt.Println(s.UntilChar(".,"))
	fmt.Println(s)

	// Output:
	// true
	// .def
	// true
	// ,def
	// false
	// ,def
}

func ExampleBytes_UntilFunc() {

	s := Bytes("abc.def")

	fmt.Println(s.UntilFunc(unicode.IsPunct))
	fmt.Println(s)
	fmt.Println(s.UntilFunc(unicode.IsPunct))
	fmt.Println(s)

	// Output:
	// true
	// .def
	// false
	// .def
}

func ExampleBytes_While() {

	s := Bytes("ababc")

	fmt.Println(s.While("ab"))
	fmt.Println(s.While("d"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// c
}

func ExampleBytes_WhileFold() {

	s := Bytes("ababc")

	fmt.Println(s.WhileFold("AB"))
	fmt.Println(s.WhileFold("D"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// c
}

func ExampleBytes_WhileChar() {

	s := Bytes("ababc")

	fmt.Println(s.WhileChar("ba"))
	fmt.Println(s.WhileChar("d"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// c
}
func ExampleBytes_WhileFunc() {

	s := Bytes("abab.")

	fmt.Println(s.WhileFunc(unicode.IsLetter))
	fmt.Println(s.WhileFunc(unicode.IsNumber))
	fmt.Println(s)

	// Output:
	// true
	// false
	// .
}

func ExampleBytes_Match() {

	s := Bytes("ab")

	fmt.Println(s.Match("a"))
	fmt.Println(s.Match("c"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// b
}

func ExampleBytes_MatchFold() {

	s := Bytes("AB")

	fmt.Println(s.MatchFold("a"))
	fmt.Println(s.MatchFold("c"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// B
}

func ExampleBytes_MatchChar() {

	s := Bytes("ab")

	fmt.Println(s.MatchChar("ba"))
	fmt.Println(s.MatchChar("c"))
	fmt.Println(s)

	// Output:
	// true
	// false
	// b
}

func ExampleBytes_MatchFunc() {

	s := Bytes("ab")

	fmt.Println(s.MatchFunc(unicode.IsLetter))
	fmt.Println(s.MatchFunc(unicode.IsDigit))
	fmt.Println(s)

	// Output:
	// true
	// false
	// b
}

func ExampleBytes_Equal() {

	s := Bytes("a")

	fmt.Println(s.Equal("a"))
	fmt.Println(s.Equal("b"))

	// Output:
	// true
	// false
}

func ExampleBytes_EqualChar() {

	s := Bytes("a")

	fmt.Println(s.EqualChar("a"))
	fmt.Println(s.EqualChar("b"))
	fmt.Println(s.EqualChar("ba"))

	// Output:
	// true
	// false
	// true
}

func ExampleBytes_EqualFold() {

	s := Bytes("A")

	fmt.Println(s.EqualFold("a"))
	fmt.Println(s.EqualFold("b"))

	// Output:
	// true
	// false
}

func ExampleBytes_EqualFunc() {

	s := Bytes("a")

	fmt.Println(s.EqualFunc(unicode.IsLetter))
	fmt.Println(s.EqualFunc(unicode.IsNumber))

	// Output:
	// true
	// false
}

func ExampleBytes_Take() {

	s := Bytes("abc")

	fmt.Println(s.Take(2))
	fmt.Println(s)

	// Output:
	// ab
	// c
}

func ExampleBytes_Peek() {

	s := Bytes("abc")

	fmt.Println(s.Peek(2))
	fmt.Println(s)
	fmt.Println(s.Peek(4))

	// Output:
	// ab
	// abc
	// abc
}

func ExampleBytes_Rune() {

	a := Bytes("世")
	b := Bytes("a")
	c := Bytes("")

	fmt.Println(a.Rune())
	fmt.Println(b.Rune())
	fmt.Println(c.Rune())

	// Output:
	// 19990
	// 97
	// 65533
}

func ExampleBytes_Byte() {

	a := Bytes("世")
	b := Bytes("a")
	c := Bytes("")

	fmt.Println(a.Byte())
	fmt.Println(b.Byte())
	fmt.Println(c.Byte())

	// Output:
	// 228
	// 97
	// 0
}

func ExampleBytes_Text() {

	s := Bytes("abc")

	m := s.Mark()
	s.Advance(2)
	fmt.Println(s.Text(m))

	// Output:
	// ab
}

func ExampleBytes_Delta() {

	s := Bytes("abc")

	m := s.Mark()
	s.Advance(2)
	fmt.Println(s.Delta(m))

	// Output:
	// ab
}

func ExampleBytes_Mark() {

	s := Bytes("abc")

	a := s.Mark()
	s.Advance(2)
	b := s.Mark()

	fmt.Println(a)
	fmt.Println(b)

	// Output:
	// abc
	// c
}

func ExampleBytes_Move() {

	s := Bytes("abc")

	m := s.Mark()
	s.Advance(2)
	s.Move(m)

	fmt.Println(s)

	// Output:
	// abc
}

func ExampleBytes_Next() {

	s := Bytes("世.界")

	fmt.Println(s.Next())
	fmt.Println(s)
	fmt.Println(s.Next())
	fmt.Println(s)
	fmt.Println(s.Next())
	fmt.Println(s)
	fmt.Println(s.Next())

	// Output:
	// true
	// .界
	// true
	// 界
	// true
	//
	// false
}
func ExampleBytes_AdvanceRune() {

	s := Bytes("世.界")

	fmt.Println(s.AdvanceRune('世'))
	fmt.Println(s)
	fmt.Println(s.AdvanceRune('.'))
	fmt.Println(s)
	fmt.Println(s.AdvanceRune('界'))
	fmt.Println(s)
	fmt.Println(s.AdvanceRune('.'))

	// Output:
	// true
	// .界
	// true
	// 界
	// true
	//
	// false
}

func ExampleBytes_Advance() {

	s := Bytes("abc")

	fmt.Println(s)
	fmt.Println(s.Advance(2))
	fmt.Println(s)
	fmt.Println(s.Advance(2))
	fmt.Println(s)

	// Output:
	// abc
	// true
	// c
	// false
	// c
}

func ExampleBytes_String() {

	s := Bytes("abc")

	fmt.Println(s.String())
	fmt.Println(s)

	// Output:
	// abc
	// abc
}

func ExampleBytes_Bytes() {

	s := Bytes("abc")

	fmt.Println(s.Bytes())

	// Output:
	// [97 98 99]
}

func ExampleBytes_More() {

	a := Bytes("a")
	b := Bytes("")

	fmt.Println(a.More())
	fmt.Println(b.More())

	// Output:
	// true
	// false
}

func ExampleBytes_Empty() {

	a := Bytes("")
	b := Bytes("a")

	fmt.Println(a.Empty())
	fmt.Println(b.Empty())

	// Output:
	// true
	// false
}

func ExampleBytes_Len() {

	a := Bytes("a")
	b := Bytes("")

	fmt.Println(a.Len())
	fmt.Println(b.Len())

	// Output:
	// 1
	// 0
}
