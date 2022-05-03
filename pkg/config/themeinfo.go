package config

// ThemeInfo is the struct for theme.yaml under theme folder
type ThemeInfo struct {
	ID          string     `yaml:"id" json:"id"`
	Name        string     `yaml:"name" json:"name"`
	Author      AuthorInfo `yaml:"author" json:"author"`
	Cover       string     `yaml:"cover" json:"cover"`
	Description string     `yaml:"description" json:"description"`
	Version     string     `yaml:"version" json:"version"`
	Repository  string     `yaml:"repository" json:"repository"`
}

type AuthorInfo struct {
	Name    string `yaml:"name" json:"name"`
	Website string `yaml:"website" json:"website"`
	Avatar  string `yaml:"avatar" json:"avatar"`
}
