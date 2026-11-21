package decode
import (
	"strings"
	"strconv"
)

// main parser for values everything else is helper
func ParseValue(value string) any {
	value = strings.TrimSpace(strings.TrimSuffix(value, ";"))

	parseParts := func(parts []string) []any {
		values := make([]any, 0, len(parts))
		for _, part := range parts {
			values = append(values, ParseValue(part))
		}
		return values
	}

	if len(value) >= 2 && value[0] == '{' && value[len(value)-1] == '}' {
		value = strings.TrimSpace(value[1 : len(value)-1])
		if value == "" {
			return []any{}
		}
		return parseParts(SplitValues(value))
	}

	if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') && value[len(value)-1] == value[0] {
		if unquoted, err := strconv.Unquote(value); err == nil {
			return unquoted
		}
	}

	if value == "true" || value == "false" {
		return value == "true"
	}

	if integer, err := strconv.ParseInt(value, 0, 64); err == nil {
		return integer
	}

	if number, err := strconv.ParseFloat(value, 64); err == nil {
		return number
	}

	if parts := SplitValues(value); len(parts) > 1 {
		return parseParts(parts)
	}

	return value
}


func SplitValues(value string) []string {
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

func SplitAssignment(line string) (string, string, bool) {
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

func ParseAssignmentValue(value, declaredType string) any {//currently only used for bools
	if declaredType == "bool" {
		if value == "1" {
			return true
		}
		if value == "0" {
			return false
		}
	}
	return ParseValue(value)
}

func VariableType(name string) string {
	start := strings.Index(name, "<")
	end := strings.LastIndex(name, ">")
	if start < 0 || end <= start {
		return ""
	}
	return strings.TrimSpace(name[start+1 : end])
}
