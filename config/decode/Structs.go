package decode

// decoded
type DConfig struct {
	Package  DPackage
	Root     DPackage
	Packages map[string]DPackage
	Errors   []string
}

type DPackage struct {
	SubPackages map[string]DPackage
	Variables   map[string]any
}

// raw
type RawConfig []RawLine
type RawLine struct {
	Content    string
	LineNumber int
}

