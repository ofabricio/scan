# scan
A text scanner

## Example

[Go Playground](https://go.dev/play/p/zuT2MvVDph0)

```go
package main

import "github.com/ofabricio/scan"

func main() {

	var s scan.Bytes = []byte(`
		KeyA = ValA
		KeyB = ValB
	`)

	Word := func(s *scan.Bytes) bool {
		return s.WhileFunc(unicode.IsLetter)
	}

	for s.Spaces(); s.More(); s.Spaces() {
		var k, v string
		if s.TextWith(&k, Word) && s.Spaces() && s.Match("=") && s.Spaces() && s.TextWith(&v, Word) {
			fmt.Println(k, v)
			continue
		}
		fmt.Println("Syntax error", s)
		break
	}

	// Output:
	// KeyA ValA
	// KeyB ValB
}
```
