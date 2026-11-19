package debug

import "main/config/decode"

type DPackage = decode.DPackage

func NewDPackage() DPackage {
	return DPackage{Variables: make(map[string]any), SubPackages: make(map[string]DPackage)}
}

func EnsureSubPackage(target *DPackage, parts []string) {
	if len(parts) == 0 {
		return
	}
	name := parts[0]
	child, ok := target.SubPackages[name]
	if !ok {
		child = NewDPackage()
	}
	if len(parts) > 1 {
		EnsureSubPackage(&child, parts[1:])
	}
	target.SubPackages[name] = child
}

func EnsurePackage(packages *map[string]DPackage, parts []string) DPackage {
	if len(parts) == 0 {
		return NewDPackage()
	}
	current := *packages
	name := parts[0]
	pkg, ok := current[name]
	if !ok {
		pkg = NewDPackage()
	}
	if pkg.SubPackages == nil {
		pkg.SubPackages = make(map[string]DPackage)
	}
	if len(parts) > 1 {
		childPackages := pkg.SubPackages
		leaf := EnsurePackage(&childPackages, parts[1:])
		pkg.SubPackages = childPackages
		current[name] = pkg
		*packages = current
		return leaf
	}
	if pkg.Variables == nil {
		pkg.Variables = make(map[string]any)
	}
	current[name] = pkg
	*packages = current
	return pkg
}
