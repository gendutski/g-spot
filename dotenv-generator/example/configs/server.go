package configs

type LogMode int

const (
	LogModeErrorOnly LogMode = iota + 1
	LogModeErrorAndWarnOnly
	LogModeAll
)

type ServerConfig struct {
	Port               int     `envconfig:"HTTP_PORT" default:"8080" prompt:"Enter port to serve http"`
	EnableDebug        bool    `envconfig:"ENABLE_DEBUG" default:"false" prompt:"Enable debug to show error received"`
	LogMode            LogMode `envconfig:"LOG_MODE" default:"1" prompt:"Enter log mode (1:error, 2:error & warn, 3:all)"`
	EnableWarnFileLog  bool    `envconfig:"ENABLE_WARN_FILE_LOG" default:"false" prompt:"Enable log for warning type error (eg: http bad request error)"`
	AutoReloadTemplate bool    `envconfig:"AUTO_RELOAD_TEMPLATE" default:"false" prompt:"Auto reload template (not recommended for production)"`
	SessionKey         string  `envconfig:"SESSION_KEY" default:"session secret" prompt:"Enter http session key secret"`
	AppURL             string  `envconfig:"APP_URL" default:"http://localhost:8080" prompt:"Enter website location"`
}
