package unmarshall
include ()

type RawConfig struct {
	map[RawLine]RawLines
}

type RawLine struct {
	Content string
	LineNumber int
}

type Package struct {
	
}