package decode
import (
	"strings"
)

func splitValues(value string) []string {
	var parts []string
	start := 0
	quoted := rune(0)
	depth := 0
	for index, char := range value {
		if quoted != 0 {
			if char == quoted && (index == 0 || value[index-1] != '\\') {
				quoted = 0
			}
			continue
		}
		switch char {
		case '"', '\'':
			quoted = char
		case '{', '[', '(':
			depth++
		case '}', ']', ')':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, strings.TrimSpace(value[start:index]))
				start = index + 1
			}
		}
	}
	return append(parts, strings.TrimSpace(value[start:]))
}

func splitAssignment(line string) (string, string, bool) {
	quoted := rune(0)
	for index, char := range line {
		if quoted != 0 {
			if char == quoted && (index == 0 || line[index-1] != '\\') {
				quoted = 0
			}
			continue
		}
		if char == '"' || char == '\'' {
			quoted = char
			continue
		}
		if char == ':' || char == '=' {
			return strings.TrimSpace(line[:index]), strings.TrimSpace(line[index+1:]), true
		}
	}
	return "", "", false
}

func parseAssignmentValue(value, declaredType string) any {
	if declaredType == "bool" {
		if value == "1" {
			return true
		}
		if value == "0" {
			return false
		}
	}
	return parseValue(value)
}

func variableType(name string) string {
	start := strings.Index(name, "<")
	end := strings.LastIndex(name, ">")
	if start < 0 || end <= start {
		return ""
	}
	return strings.TrimSpace(name[start+1 : end])
}

func variableName(name string) string {
	name = strings.TrimSpace(name)
	if index := strings.Index(name, "<"); index >= 0 {
		name = name[:index]
	}
	return strings.TrimSpace(name)
}

