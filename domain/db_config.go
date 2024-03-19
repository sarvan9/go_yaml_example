package domain

type DbConfigRoot struct {
	DbConfig DbConfig `yaml:"dbConfig"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	UserName string `yaml: "username"`
	Password string `yaml:"password"`
}
