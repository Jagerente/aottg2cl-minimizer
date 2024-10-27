package minimizer

import (
	"fmt"
	"strings"
)

type StringsPreserver struct {
	stringsList []string
}

func NewStringsPreserver() *StringsPreserver {
	return &StringsPreserver{
		stringsList: make([]string, 0),
	}
}

func (p *StringsPreserver) Preserve(code string) string {
	var preservedCode strings.Builder
	var insideString bool
	var currentString strings.Builder

	for _, char := range code {
		if char == '"' {
			if insideString {
				currentString.WriteRune(char)
				p.stringsList = append(p.stringsList, currentString.String())
				preservedCode.WriteString(fmt.Sprintf("%%STR%d%%", len(p.stringsList)-1))
				currentString.Reset()
				insideString = false
			} else {
				insideString = true
				currentString.WriteRune(char)
			}
		} else if insideString {
			currentString.WriteRune(char)
		} else {
			preservedCode.WriteRune(char)
		}
	}

	return preservedCode.String()
}

func (p *StringsPreserver) Restore(code string) string {
	for i, str := range p.stringsList {
		placeholder := fmt.Sprintf("%%STR%d%%", i)
		code = strings.ReplaceAll(code, placeholder, str)
	}
	p.stringsList = make([]string, 0)
	return code
}
