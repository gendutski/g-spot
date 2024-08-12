package main

import (
	dotenvgenerator "github.com/gendutski/g-spot/dotenv-generator"
	"github.com/gendutski/g-spot/dotenv-generator/example/configs"
)

func main() {
	dotenvgenerator.GenerateDotEnv(configs.ServerConfig{}, true, true, nil)
}
