package configs

type JwtConfig struct {
	JwtSecret              string `envconfig:"JWT_SECRET" default:"secret" prompt:"Enter secret to generate JWT token"`
	JwtExpirationInMinutes int    `envconfig:"JWT_EXPIRATION_IN_MINUTES" default:"60" prompt:"Enter token expired in minute"`
	JwtRememberInDays      int    `envconfig:"JWT_REMEMBER_IN_DAYS" default:"30" prompt:"Enter token remember(for remember login) in days"`
}
