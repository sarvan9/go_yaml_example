package loader

import (
	"fmt"
	"log"
	"os"

	"example.com/yaml_parser/domain"
	"gopkg.in/yaml.v3"
)

var ServerConfig domain.ServerRoot
var SecurityConfig domain.SecurityConfigRoot
var DbConfig domain.DbConfigRoot

func init() {
	loadConfigs()
}

func loadConfigs() {

	yamlData, err := os.ReadFile("app_config.yml")

	if err != nil {
		log.Fatal("Error while reading app config file", err)
	}

	laodServerConfig(yamlData)
	laodDBConfig(yamlData)
	laodSecurityConfig(yamlData)

}

func laodServerConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &ServerConfig)
	fmt.Println("Loaded Server Config")
}

func laodDBConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &DbConfig)
	fmt.Println("Loaded DB Config")
}

func laodSecurityConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &SecurityConfig)
	fmt.Println("Loaded Security Config")
}
