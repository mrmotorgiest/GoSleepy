package decode

import (
	"path/filepath"
	"strings"
)

// code decodes other code \\
func DecodeFile(path string) DConfig {
	var config = DConfig{
		Root:     DPackage{Variables: make(map[string]any), SubPackages: make(map[string]DPackage)},
		Packages: make(map[string]DPackage),
		Errors:   []string{},
	} // dont remove otherwise itll brake for some reason

	var absolutePath, _ = filepath.Abs(path)
	LoadFile(absolutePath, &config)

	return config
}

func ParseSource(source string, target *DPackage) {
	var sectionPath []string
	for _, raw := range ConvertString(source) {
		var line = strings.TrimSpace(StripComment(raw.Content))
		if line == "" || strings.HasPrefix(line, "$include") {
			continue
		}

		switch {
		case line == "}":
			if len(sectionPath) > 0 {
				sectionPath = sectionPath[:len(sectionPath)-1]
			}
		case strings.HasSuffix(line, "{"):
			var name = strings.TrimSuffix(line, "{")
			name = BlockName(strings.TrimSpace(name))

			if name != "" {
				sectionPath = append(sectionPath, strings.Split(name, ".")...)
				EnsureSubPackage(target, sectionPath)
			}
		default:
			key, value, _ := SplitAssignment(line)
			SetPackageVariable(target, sectionPath, key, ParseAssignmentValue(value, VariableType(key)))
		}
	}
}
