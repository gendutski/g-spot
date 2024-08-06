package main

import dotenvgenerator "github.com/gendutski/g-spot/dotenv-generator"

func main() {
	var cfg config
	dotenvgenerator.GenerateDotEnv(cfg)
}

type config struct {
	Port        int    `envconfig:"HTTP_PORT" default:"8080" prompt:"Enter port to serve http" comment:"main config"`
	EnableDebug bool   `envconfig:"ENABLE_DEBUG" default:"false" prompt:"Enable debug to show error received"`
	LogMode     int    `envconfig:"LOG_MODE" default:"1" prompt:"Enter log mode (1:error, 2:error & warn, 3:all)"`
	SessionKey  string `envconfig:"SESSION_KEY" default:"session secret" prompt:"Enter http session key secret"`
	AppURL      string `envconfig:"APP_URL" default:"http://localhost:8080" prompt:"Enter website location"`

	// jwt
	JwtSecret              string `envconfig:"JWT_SECRET" default:"secret" prompt:"Enter secret to generate JWT token" comment:"jwt config"`
	JwtExpirationInMinutes int    `envconfig:"JWT_EXPIRATION_IN_MINUTES" default:"60" prompt:"Enter token expired in minute"`
	JwtRememberInDays      int    `envconfig:"JWT_REMEMBER_IN_DAYS" default:"30" prompt:"Enter token remember(for remember login) in days"`

	// smtp
	MailHost     string `envconfig:"MAIL_HOST" default:"smtp.gmail.com" prompt:"Enter smtp server host" comment:"smtp config"`
	MailPort     int    `envconfig:"MAIL_PORT" default:"465" prompt:"Enter smtp server port"`
	MailUser     string `envconfig:"MAIL_USER" prompt:"Enter smtp server user"`
	MailPassword string `envconfig:"MAIL_PASSWORD" prompt:"Enter smtp server password" secret:"true"`

	// mysql
	MysqlHost                  string `envconfig:"MYSQL_HOST" default:"localhost" prompt:"Enter mysql host" comment:"mysql config"`
	MysqlPort                  int    `envconfig:"MYSQL_PORT" default:"3306" prompt:"Enter mysql port"`
	MysqlDBName                string `envconfig:"MYSQL_DB_NAME" default:"" prompt:"Enter database name"`
	MysqlUsername              string `envconfig:"MYSQL_USERNAME" default:"" prompt:"Enter mysql username"`
	MysqlPassword              string `envconfig:"MYSQL_PASSWORD" default:"" prompt:"Enter mysql password" secret:"true"`
	MysqlLogMode               int    `envconfig:"MYSQL_LOG_MODE" default:"1" prompt:"Enter gorm log mode 1-4"`
	MysqlParseTime             bool   `envconfig:"MYSQL_PARSE_TIME" default:"true" prompt:"Parse mysql time to local"`
	MysqlCharset               string `envconfig:"MYSQL_CHARSET" default:"utf8mb4" prompt:"Enter mysql database charset"`
	MysqlLoc                   string `envconfig:"MYSQL_LOC" default:"Local" prompt:"Enter mysql local time"`
	MysqlMaxLifetimeConnection int    `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION" default:"10" prompt:"Enter mysql maximum amount of time a connection may be reused, in minute"`
	MysqlMaxOpenConnection     int    `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"50" prompt:"Enter mysql maximum number of open connections to the database"`
	MysqlMaxIdleConnection     int    `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"10" prompt:"Enter mysql maximum number of connections in the idle connection pool"`
}
