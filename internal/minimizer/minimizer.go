package minimizer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Minimizer struct {
	stringsPreserver *StringsPreserver
}

func New() *Minimizer {
	return &Minimizer{
		stringsPreserver: NewStringsPreserver(),
	}
}

func (m *Minimizer) MinimizeFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	code, err := m.minimize(content)
	if err != nil {
		return fmt.Errorf("failed to minimize code: %w", err)
	}

	fileName := filepath.Base(filePath)
	outputFileName := fmt.Sprintf("minimized_%s", fileName)
	err = os.WriteFile(outputFileName, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("failed to write minimized file: %w", err)
	}

	fmt.Printf("Minimized code written to %s\n", outputFileName)
	return nil
}

func (m *Minimizer) minimize(src []byte) (string, error) {
	code := string(src)

	code = m.stringsPreserver.Preserve(code)

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if !strings.Contains(line, "\"") && !strings.Contains(line, "'") {
			lines[i] = m.applyRegex(line, `#.*$`, "")
		}
	}
	code = strings.Join(lines, "\n")
	code = m.applyRegex(code, `/\*[\s\S]*?\*/`, "")

	code = m.applyRegex(code, `^(\s*)[ ]{4}(\s*)`, "$1\\t$2")
	code = m.applyRegex(code, `^(\t*)[ ]+`, "$1")
	code = m.applyRegex(code, `(?m)^\s+`, "")

	code = m.applyRegex(code, `\s*==\s*|\n+==\s*|\s*==\n+`, "==")
	code = m.applyRegex(code, `\s*!=\s*|\n+!=\s*|\s*!=\n+`, "!=")
	code = m.applyRegex(code, `\s*>\s*|\n*>\s*|\s*>+\n*`, ">")
	code = m.applyRegex(code, `\s*<\s*|\n+<\s*|\s*<\n+`, "<")
	code = m.applyRegex(code, `\s*>=\s*|\n+>=\s*|\s*>=\n+`, ">=")
	code = m.applyRegex(code, `\s*<=\s*|\n+<=\s*|\s*<=\n+`, "<=")

	code = m.applyRegex(code, `\s*&&\s*|\n+&&\s*|\s*&&\n+`, "&&")
	code = m.applyRegex(code, `\s*\|\|\s*|\n+\|\|\s*|\s*\|\|\n+`, "||")

	code = m.applyRegex(code, `\s*\+=\s*|\n*\+=\s*|\s*\+=\n+`, "+=")
	code = m.applyRegex(code, `\s*-\=\s*|\n*-\=\s*|\s*-\=\n+`, "-=")
	code = m.applyRegex(code, `\s*\*=\s*|\n*\*=\s*|\s*\*=\n+`, "*=")
	code = m.applyRegex(code, `\s*/=\s*|\n*/=\s*|\s*/=\n+`, "/=")

	code = m.applyRegex(code, `\s*=\s*|\n*=\s*|\s*=\n+`, "=")
	// BUG: Disabled until bug is fixed
	// https://discord.com/channels/681641241125060652/1300029742040219722
	//code = ApplyRegex(code, `\s*-\s*|\n+-\s*|\s*-\n+`, "-")
	code = m.applyRegex(code, `\n-\s*|\s*-\n`, " - ")
	code = m.applyRegex(code, `\s*\+\s*|\n+\s*\+\s*|\s*\+\s*\n+`, "+")
	code = m.applyRegex(code, `\s*\*\s*|\n+\s*\*\s*|\s*\*\s*\n+`, "*")
	code = m.applyRegex(code, `\s*/\s*|\n+/s*|\s*/\s*\n+`, "/")
	code = m.applyRegex(code, `\s*,\s*|\n*,\s*|\s*,\n+`, ",")

	code = m.applyRegex(code, `;\s*\n|\s*;\n+`, ";")
	code = m.applyRegex(code, `\s*\(\s*|\n*\(\s*|\s*\(\n+`, "(")
	code = m.applyRegex(code, `\s*\)\s*|\n*\)\s*|\s*\)\n+`, ")")
	code = m.applyRegex(code, `\s*\{\s*|\n*\{\s*|\s*\{\n+`, "{")
	code = m.applyRegex(code, `\s*\}\s*|\n*\}\s*|\s*\}\n+`, "}")

	code = m.stringsPreserver.Restore(code)

	code = strings.TrimSpace(code)

	return code, nil
}

func (m *Minimizer) applyRegex(input string, pattern string, replacement string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(input, replacement)
}
