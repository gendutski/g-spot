package configs

type DBConfig struct {
	MysqlHost                  string `envconfig:"MYSQL_HOST" default:"localhost" prompt:"Enter mysql host"`
	MysqlPort                  int    `envconfig:"MYSQL_PORT" default:"3306" prompt:"Enter mysql port"`
	MysqlDBName                string `envconfig:"MYSQL_DB_NAME" prompt:"Enter database name"`
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
