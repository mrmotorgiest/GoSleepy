package decode

import "strings"

func ConvertString(file string) RawConfig {
	lines := strings.Split(file, "\n")
	config := make(RawConfig, 0, len(lines))
	for index, line := range lines {
		config = append(config, RawLine{Content: strings.TrimSuffix(line, "\r"), LineNumber: index + 1})
	}
	return config
}
