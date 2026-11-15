package decode

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Code decodes other code \\
func Decode(source string) DConfig {
	config := newDConfig()
	parseSource(source, &config.Root)
	return config
}


func DecodeFile(path string) DConfig {
	config := newDConfig()
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		config.Errors = append(config.Errors, err.Error())
		return config
	}
	loadFile(absolutePath, absolutePath, &config, make(map[string]bool))
	return config
}

func parseSource(source string, target *DPackage) {
	var sectionPath []string
	for _, raw := range ConvertString(source) {
		line := strings.TrimSpace(stripComment(raw.Content))

		if line == "" || strings.HasPrefix(line, "$include") {
			continue
		}

		switch {
			case line == "}":
				if len(sectionPath) > 0 {
					sectionPath = sectionPath[:len(sectionPath)-1]
				}
		case strings.HasSuffix(line, "{"):
			if name := blockName(strings.TrimSpace(strings.TrimSuffix(line, "{"))); name != "" {
				sectionPath = append(sectionPath, strings.Split(name, ".")...)
				ensureSubPackage(target, sectionPath)
			}
		default:
			key, value, ok := splitAssignment(line)
			if !ok {
				continue
			}

			pkg := packageAt(target, sectionPath)
			pkg.Variables[key] = parseAssignmentValue(value, variableType(key))
		}
	}
}

func newDConfig() DConfig {
	return DConfig{Root: newDPackage(), Packages: make(map[string]DPackage)}
}

func (config DConfig) String() string {
	var builder strings.Builder
	builder.WriteString("root\n")
	writePackage(&builder, config.Root, "  ")
	packageNames := make([]string, 0, len(config.Packages))
	for name := range config.Packages {
		packageNames = append(packageNames, name)
	}
	sort.Strings(packageNames)
	for _, name := range packageNames {
		fmt.Fprintf(&builder, "package %s\n", name)
		writePackage(&builder, config.Packages[name], "  ")
	}
	for _, err := range config.Errors {
		fmt.Fprintf(&builder, "error: %s\n", err)
	}
	return strings.TrimRight(builder.String(), "\n")
}
