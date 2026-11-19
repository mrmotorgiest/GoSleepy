package decode
import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"main/config/debug/errors"
)

var InBlock = regexp.MustCompile(`(?s)\$include\s*\{([^}]*)\}`)
var InString = regexp.MustCompile(`"((?:\\.|[^"\\])*)"|'((?:\\.|[^'\\])*)'`)

// including of raw individualy \\
func LoadFile(path string, config *DConfig) {
	data, err := os.ReadFile(path)
	errors.ReceiveError(err, "LoadFile: os.ReadFile")

	var source = string(data)
	for _, block := range InBlock.FindAllStringSubmatch(source, -1) {
		for _, match := range InString.FindAllStringSubmatch(block[1], -1) {
			var includePath = match[1]

			includePath = strings.ReplaceAll(includePath, `\"`, `"`)
			if !filepath.IsAbs(includePath) {
				includePath = filepath.Join(filepath.Dir(path), includePath)
			}

			matches, _ := filepath.Glob(includePath)
			sort.Strings(matches) 
			for _, match := range matches {
				LoadIncludedPath(match, config)
			}
		}
	}

	cleanSource := InBlock.ReplaceAllString(source, "")
	if strings.HasSuffix(strings.ToLower(path), ".gsc.package") {
		var packageParts = strings.Split(filepath.ToSlash(filepath.Dir(path)), "/")

		if packageParts[0] == "." {
			packageParts = packageParts[1:]
		}else if len(packageParts) > 0 && packageParts[0] == "packages" {
			packageParts = packageParts[1:]
		}

		packageParts = append(packageParts, strings.TrimSuffix(filepath.Base(path), ".gsc.package"))
		ParsePackageSource(config.Packages, packageParts, cleanSource)
		return
	}
	ParseSource(cleanSource, &config.Root)
}

// inlcudes files as "path/*" \\
func LoadIncludedPath(path string, config *DConfig) {
	err := filepath.Walk(path, func(CurrentPath string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			LoadFile(CurrentPath, config)
		}
		errors.ReceiveError(err, fmt.Sprintf("failed at LoadIncludedPath: %q", CurrentPath))
		return nil
	})
	errors.ReceiveError(err, fmt.Sprintf("LoadIncludedPath: filepath.Walk %q", path))
}
