package main

import (
	dotenvgenerator "github.com/gendutski/g-spot/dotenv-generator"
	"github.com/gendutski/g-spot/dotenv-generator/example/configs"
)

func main() {
	dotenvgenerator.GenerateDotEnv([]interface{}{
		configs.ServerConfig{},
		configs.JwtConfig{},
		configs.SmtpConfig{},
		configs.DBConfig{},
	}, true, true, nil)
}
