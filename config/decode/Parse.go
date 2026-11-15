package decode
import (
	"strings"
	"strconv"
)

func stripComment(line string) string {
	quoted := rune(0)
	escaped := false
	for index, char := range line {
		if escaped {
			escaped = false
			continue
		}
		if char == '\\' && quoted != 0 {
			escaped = true
			continue
		}
		if quoted != 0 {
			if char == quoted {
				quoted = 0
			}
			continue
		}
		if char == '"' || char == '\'' {
			quoted = char
			continue
		}
		if char == '/' && strings.HasPrefix(line[index:], "//") {
			return line[:index]
		}
	}
	return line
}

func parseValue(value string) any {
	value = strings.TrimSpace(strings.TrimSuffix(value, ";"))
	if len(value) >= 2 && value[0] == '{' && value[len(value)-1] == '}' {
		value = strings.TrimSpace(value[1 : len(value)-1])
		if value == "" {
			return []any{}
		}
		parts := splitValues(value)
		values := make([]any, 0, len(parts))
		for _, part := range parts {
			values = append(values, parseValue(part))
		}
		return values
	}
	if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		if unquoted, err := strconv.Unquote(value); err == nil {
			return unquoted
		}
	}
	if len(value) >= 2 && value[0] == '\'' && value[len(value)-1] == '\'' {
		if unquoted, err := strconv.Unquote(value); err == nil {
			return string(unquoted)
		}
	}
	if value == "true" {
		return true

	}
	if value == "false" {
		return false
	}
	if integer, err := strconv.ParseInt(value, 0, 64); err == nil {
		return integer
	}
	if number, err := strconv.ParseFloat(value, 64); err == nil {
		return number
	}
	parts := splitValues(value)
	if len(parts) > 1 {
		values := make([]any, 0, len(parts))
		for _, part := range parts {
			values = append(values, parseValue(part))
		}
		return values
	}
	return value
}



func blockName(name string) string {
	if index := strings.IndexAny(name, "(<"); index >= 0 {
		name = name[:index]
	}
	return strings.Trim(strings.TrimSpace(name), ".")
}

