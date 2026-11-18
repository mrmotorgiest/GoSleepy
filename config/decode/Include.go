package decode

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

//longgg shiieet
var includeBlockPattern = regexp.MustCompile(`(?s)\$include\s*\{([^}]*)\}`)
var includeStringPattern = regexp.MustCompile(`"((?:\\.|[^"\\])*)"|'((?:\\.|[^'\\])*)'`)

// Including of raw files \\
func loadFile(path, rootPath string, config *DConfig, loaded map[string]bool) {
	path, err := filepath.Abs(path)
	if err != nil {
		config.Errors = append(config.Errors, err.Error())
		return
	}
	if loaded[path] {
		return
	}
	loaded[path] = true

	data, err := os.ReadFile(path)
	if err != nil {
		config.Errors = append(config.Errors, fmt.Sprintf("%s: %v", path, err))
		return
	}
	source := string(data)
	for _, includeMatch := range includeBlockPattern.FindAllStringSubmatch(source, -1) {
		for _, includeMatch := range includeStringPattern.FindAllStringSubmatch(includeMatch[1], -1) {
			includePath := includeMatch[1]
			if includePath == "" {
				includePath = includeMatch[2]
			}
			includePath = strings.ReplaceAll(includePath, `\"`, `"`)
			pattern := includePath
			if !filepath.IsAbs(pattern) {
				pattern = filepath.Join(filepath.Dir(path), pattern)
			}
			matches, globErr := filepath.Glob(pattern)
			if globErr != nil {
				config.Errors = append(config.Errors, fmt.Sprintf("%s: invalid include %q: %v", path, includePath, globErr))
				continue
			}
			sort.Strings(matches)
			for _, match := range matches {
				loadIncludedPath(match, rootPath, config, loaded)
			}
		}
	}

	cleanSource := includeBlockPattern.ReplaceAllString(source, "")
	if strings.HasSuffix(strings.ToLower(path), ".gsc.package") {
		packagePath, relErr := filepath.Rel(filepath.Dir(rootPath), path)
		if relErr != nil {
			config.Errors = append(config.Errors, relErr.Error())
			return
		}
		packageParts := strings.Split(filepath.ToSlash(filepath.Dir(packagePath)), "/")
		if packageParts[0] == "." {
			packageParts = packageParts[1:]
		}
		if len(packageParts) > 0 && packageParts[0] == "packages" {
			packageParts = packageParts[1:]
		}
		packageParts = append(packageParts, strings.TrimSuffix(filepath.Base(path), ".gsc.package"))
		parsePackageSource(config.Packages, packageParts, cleanSource)
		return
	}
	parseSource(cleanSource, &config.Root)
}

func loadIncludedPath(path, rootPath string, config *DConfig, loaded map[string]bool) {
	info, err := os.Stat(path)
	if err != nil {
		config.Errors = append(config.Errors, fmt.Sprintf("%s: %v", path, err))
		return
	}
	if !info.IsDir() {
		loadFile(path, rootPath, config, loaded)
		return
	}
	err = filepath.Walk(path, func(currentPath string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			config.Errors = append(config.Errors, fmt.Sprintf("%s: %v", currentPath, walkErr))
			return nil
		}
		if !info.IsDir() {
			loadFile(currentPath, rootPath, config, loaded)
		}
		return nil
	})
	if err != nil {
		config.Errors = append(config.Errors, fmt.Sprintf("%s: %v", path, err))
	}
}
