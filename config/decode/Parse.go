package decode
import (
	"strings"
)

// cleanup
func StripComment(line string) string {
	var quoted = rune(0)
	var escaped = false
	
	for index, char := range line {
		switch char {
			case '"', '\'' :
				if escaped { continue }

				if quoted == 0 {
					quoted = char
				} else if quoted == char {
					quoted = 0
				}
			case '\\':
				if quoted != 0 {
					escaped = !escaped
				}
			case '/':
				if !escaped && quoted == 0 && strings.HasPrefix(line[index:], "//") {
					return line[:index]
				}
			default:
				escaped = false
		}
		if quoted != 0 && char == quoted {
			quoted = 0
		}
	}
	return line
}

// misc
func ConvertString(file string) RawConfig {
	var lines = strings.Split(file, "\n")
	var config = make(RawConfig, 0, len(lines))
	
	for index, line := range lines {
		config = append(config, RawLine{Content: strings.TrimSuffix(line, "\r"), LineNumber: index + 1})
	}
	return config
}

func BlockName(name string) string {
	var index = strings.IndexAny(name, "(<") 
	if index >= 0 { 
		name = name[:index] 
	}
	return strings.Trim(strings.TrimSpace(name), ".")
}
