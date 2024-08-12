package configs

type SmtpConfig struct {
	MailHost          string `envconfig:"MAIL_HOST" default:"smtp.gmail.com" prompt:"Enter smtp server host"`
	MailPort          int    `envconfig:"MAIL_PORT" default:"465" prompt:"Enter smtp server port"`
	MailUser          string `envconfig:"MAIL_USER" prompt:"Enter smtp server user"`
	MailPassword      string `envconfig:"MAIL_PASSWORD" prompt:"Enter smtp server password" secret:"true"`
	MailSenderName    string `envconfig:"MAIL_SENDER_NAME" prompt:"Enter email sender name"`
	MailSenderAddress string `envconfig:"MAIL_SENDER_ADDRESS" prompt:"Enter email sender address"`
}
