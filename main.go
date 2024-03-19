package main

import (
	"fmt"

	"example.com/yaml_parser/loader"
)

func main() {
	fmt.Println("DB Config loaded : %v", loader.DbConfig)
	fmt.Println("Security Config loaded : %v", loader.SecurityConfig)
	fmt.Println("Server Config loaded : %v", loader.ServerConfig)
}
