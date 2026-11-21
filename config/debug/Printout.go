package debug
import (
    "fmt"
    "sort"
    "strings"

    "main/config/decode"
)

// sloppy code
func SortedKeys[T any](m map[string]T) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    return keys
}

func PrintConfig(config decode.DConfig) {
	var b strings.Builder

	var write func(decode.DPackage, string)
	write = func(pkg decode.DPackage, indent string) {
		for _, n := range SortedKeys(pkg.Variables) {
			fmt.Fprintf(&b, "%s%s = %#v\n", indent, n, pkg.Variables[n])
		}
		for _, n := range SortedKeys(pkg.SubPackages) {
			fmt.Fprintf(&b, "%s%s\n", indent, n)
			write(pkg.SubPackages[n], indent+"  ")
		}
	}

	write(config.Root, "")
	for _, n := range SortedKeys(config.Packages) {
		fmt.Fprintln(&b, n)
		write(config.Packages[n], "  ")
	}
	fmt.Print(b.String())
}