package config

type MenuBar struct {
	padding string `yaml:"padding"`
	spacing string `yaml:"spacing"`

	orentation string `yaml:"orientation"`

	items []MenuItem `yaml:"items"`
}

type MenuItem struct {
	label string `yaml:"label"`
	action string `yaml:"action"`
	subitems []MenuItem `yaml:"subitems"`
}