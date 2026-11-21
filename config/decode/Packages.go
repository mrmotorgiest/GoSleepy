package decode
import ()

func SetPackageVariable(target *DPackage, parts []string, key string, value any) {
	if len(parts) == 0 {
		target.Variables[key] = value
		return
	}
	name := parts[0]
	
	child, _ := target.SubPackages[name]

	SetPackageVariable(&child, parts[1:], key, value)
	target.SubPackages[name] = child
}

func ParsePackageSource(packages map[string]DPackage, parts []string, source string) {
	var name = parts[0]
	var pkg = packages[name]

	if len(parts) == 1 {
		ParseSource(source, &pkg)
	} else {
		ParsePackageSource(pkg.SubPackages, parts[1:], source)
	}
	packages[name] = pkg
}

// misc
// used in Decode.go
func EnsureSubPackage(target *DPackage, parts []string) {
	if len(parts) == 0 {
		return
	}// code will break if this is removed
	
	var name = parts[0]
	var child = target.SubPackages[name]
	if child.SubPackages == nil {
		child = NewDPackage()
	}

	EnsureSubPackage(&child, parts[1:])
	target.SubPackages[name] = child
}

func NewDPackage() DPackage {
	return DPackage{Variables: make(map[string]any), SubPackages: make(map[string]DPackage)}
}
