package decode

import (
	"fmt"
	"sort"
	"strings"
)

func writePackage(builder *strings.Builder, pkg DPackage, indent string) {
	variableNames := make([]string, 0, len(pkg.Variables))
	for name := range pkg.Variables {
		variableNames = append(variableNames, name)
	}
	sort.Strings(variableNames)
	for _, name := range variableNames {
		fmt.Fprintf(builder, "%s%s = %#v\n", indent, name, pkg.Variables[name])
	}
	packageNames := make([]string, 0, len(pkg.SubPackages))
	for name := range pkg.SubPackages {
		packageNames = append(packageNames, name)
	}
	sort.Strings(packageNames)
	for _, name := range packageNames {
		fmt.Fprintf(builder, "%s%s\n", indent, name)
		writePackage(builder, pkg.SubPackages[name], indent+"  ")
	}
}

func ensureSubPackage(target *DPackage, parts []string) {
	if len(parts) == 0 {
		return
	}
	name := parts[0]
	child, ok := target.SubPackages[name]
	if !ok {
		child = newDPackage()
	}
	if len(parts) > 1 {
		ensureSubPackage(&child, parts[1:])
	}
	target.SubPackages[name] = child
}

func setPackageVariable(target *DPackage, parts []string, key string, value any) {
	if len(parts) == 0 {
		target.Variables[key] = value
		return
	}
	name := parts[0]
	child, ok := target.SubPackages[name]
	if !ok {
		child = newDPackage()
	}
	setPackageVariable(&child, parts[1:], key, value)
	target.SubPackages[name] = child
}

func parsePackageSource(packages map[string]DPackage, parts []string, source string) {
	if len(parts) == 0 {
		return
	}
	name := parts[0]
	pkg, ok := packages[name]
	if !ok {
		pkg = newDPackage()
	}
	if len(parts) == 1 {
		parseSource(source, &pkg)
	} else {
		parsePackageSource(pkg.SubPackages, parts[1:], source)
	}
	packages[name] = pkg
}

func newDPackage() DPackage {
	return DPackage{Variables: make(map[string]any), SubPackages: make(map[string]DPackage)}
}
