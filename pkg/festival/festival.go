package festival

type Band struct {
	Name  string `yaml:"name"`
	Start string `yaml:"start"`
	Stop  string `yaml:"stop"`
}

type Day struct {
	Date  string `yaml:"day"`
	Bands []Band `yaml:"bands"`
}

type Location struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Days        []Day  `yaml:"days"`
}

type Festival struct {
	Name      string     `yaml:"name"`
	Locations []Location `yaml:"locations"`
}
