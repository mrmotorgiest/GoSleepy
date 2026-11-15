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

func packageAt(target *DPackage, parts []string) *DPackage {
	if len(parts) == 0 {
		return target
	}
	child := target.SubPackages[parts[0]]
	return packageAt(&child, parts[1:])
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

func ensurePackage(packages *map[string]DPackage, parts []string) DPackage {
	if len(parts) == 0 {
		return newDPackage()
	}
	current := *packages
	name := parts[0]
	pkg, ok := current[name]
	if !ok {
		pkg = newDPackage()
	}
	if len(parts) > 1 {
		children := pkg.SubPackages
		leaf := ensurePackage(&children, parts[1:])
		_ = leaf
		pkg.SubPackages = children
	}
	current[name] = pkg
	*packages = current
	return pkg
}

func newDPackage() DPackage {
	return DPackage{Variables: make(map[string]any), SubPackages: make(map[string]DPackage)}
}
