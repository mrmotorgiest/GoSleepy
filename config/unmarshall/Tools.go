package unmarshall
// File \\
func CleanConfig(file string) file string {
	for 
	return file
}

func ConvertString(file string) RawConfig {
	var config RawConfig
	
	for i, line := range file {
		config[i] = RawLine{Content: line, LineNumber: i}
	}

	return config
}