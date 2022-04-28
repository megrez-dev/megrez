package config

type ItemType string

const (
	ItemTypeInput    ItemType = "input"
	ItemTypeTextarea ItemType = "textarea"
	ItemTypeRadio    ItemType = "radio"
	ItemTypeCheckbox ItemType = "checkbox"
	ItemTypeList     ItemType = "list"
	ItemTypeSelect   ItemType = "select"
	ItemTypeImage    ItemType = "image"
)

type ThemeConfig struct {
	Tabs []Tab `yaml:"tabs"`
}

type Tab struct {
	Name  string `yaml:"name"`
	Label string `yaml:"label"`
	Items []Item `yaml:"items"`
}

type Item struct {
	Name        string   `yaml:"name"`
	Label       string   `yaml:"label"`
	Type        string   `yaml:"type"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description"`
	Value       string   `yaml:"value"`
	Options     []string `yaml:"options"`
}
