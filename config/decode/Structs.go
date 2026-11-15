package decode

type DConfig struct {
	RawFiles []RawFile
	Package  DPackage
	Root     DPackage
	Packages map[string]DPackage
	Errors   []string
}

type RawLine struct {
	Content    string
	LineNumber int
}

type RawFile struct {
	Path     string
	RawLines RawConfig
}

type RawConfig []RawLine

type DPackage struct {
	SubPackages map[string]DPackage
	Variables   map[string]any
}

type SubPackage struct {
	Variables map[string]any
}

type Package struct {
	SubPackages map[string]SubPackage
	Variables   map[string]any
}
