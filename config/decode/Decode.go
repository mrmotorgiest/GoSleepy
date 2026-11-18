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
	}// dont remove otherwise itll brake for some reason

	var absolutePath, _ = filepath.Abs(path)
	loadFile(absolutePath, absolutePath, &config, make(map[string]bool))
	return config
}

func parseSource(source string, target *DPackage) {
	var sectionPath []string
	for _, raw := range ConvertString(source) {
		var line = strings.TrimSpace(stripComment(raw.Content))

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
				name = blockName(strings.TrimSpace(name))

				if name != "" {
					sectionPath = append(sectionPath, strings.Split(name, ".")...)
					ensureSubPackage(target, sectionPath)
				}
			default:
				key, value, _ := splitAssignment(line)
				setPackageVariable(target, sectionPath, key, parseAssignmentValue(value, variableType(key)))
		}
	}
}