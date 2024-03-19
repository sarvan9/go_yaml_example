package domain

type SecurityConfigRoot struct {
	SecurityConfig SecurityConfig `yaml:"security"`
}

type SecurityConfig struct {
	SslEnabled         bool   `yaml:"sslEnabled"`
	TruststoreFilePath string `yaml:"truststoreLocation"`
	TruststorePwd      string `yaml:"truststorePassword"`
}
