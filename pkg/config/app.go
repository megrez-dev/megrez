package config

type Config struct {
	Database Database `yaml:"database"`
	Debug    bool     `yaml:"debug"`
}

type Database struct {
	Bolt   Bolt   `yaml:"bolt"`
	MySQL  MySQL  `yaml:"mysql"`
	SQLite SQLite `yaml:"sqlite"`
}

type Bolt struct {
	Path string `yaml:"path"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type SQLite struct {
	Path string `yaml:"path"`
}
