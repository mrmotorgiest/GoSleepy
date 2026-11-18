package debug

import (
	"fmt"
	"sort"
	"strings"

	"main/config/decode"
)

// PrintConfig prints the decoded config with stable ordering.
func PrintConfig(config decode.DConfig) {
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
	fmt.Print(strings.TrimRight(builder.String(), "\n"))
}

func writePackage(builder *strings.Builder, pkg decode.DPackage, indent string) {
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
