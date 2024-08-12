package dotenvgenerator_test

import (
	"testing"

	dotenvgenerator "github.com/gendutski/g-spot/dotenv-generator"
	mock_dotenvgenerator "github.com/gendutski/g-spot/dotenv-generator/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type config1 struct {
	Port            int    `envconfig:"HTTP_PORT" default:"8080" prompt:"Enter port to serve http"`
	DefaultUsername string `envconfig:"DEFAULT_USERNAME" prompt:"Enter Login Username"`
	DefaultPassword string `envconfig:"DEFAULT_PASSWORD" prompt:"Enter Login Password" secret:"true"`
}

type config2 struct {
	MysqlHost      string `envconfig:"MYSQL_HOST" default:"localhost" prompt:"Enter mysql host"`
	MysqlPort      int    `envconfig:"MYSQL_PORT" default:"3306" prompt:"Enter mysql port"`
	MysqlDBName    string `envconfig:"MYSQL_DB_NAME" prompt:"Enter database name"`
	MysqlUsername  string `envconfig:"MYSQL_USERNAME" default:"" prompt:"Enter mysql username"`
	MysqlPassword  string `envconfig:"MYSQL_PASSWORD" default:"" prompt:"Enter mysql password" secret:"true"`
	MysqlParseTime bool   `envconfig:"MYSQL_PARSE_TIME" default:"false" prompt:"Parse mysql time to local"`
}

func Test_GenerateDotEnv(t *testing.T) {
	ctrl := gomock.NewController(t)
	p := mock_dotenvgenerator.NewMockPrompter(ctrl)

	t.Run("single struct, all field use default value", func(t *testing.T) {
		p.EXPECT().PromptString("Enter port to serve http (default:8080): ").Return("").Times(1)
		p.EXPECT().PromptString("Enter Login Username: ").Return("").Times(1)
		p.EXPECT().PromptPassword("Enter Login Password: ").Return("").Times(1)

		result, err := dotenvgenerator.GenerateDotEnv(config1{}, false, false, p)
		assert.Nil(t, err)
		assert.Equal(t, []string{
			"HTTP_PORT=\"8080\"",
			"DEFAULT_USERNAME=\"\"",
			"DEFAULT_PASSWORD=\"\"",
		}, result)
	})

	t.Run("single struct, all field value set", func(t *testing.T) {
		p.EXPECT().PromptString("Enter mysql host (default:localhost): ").Return("127.0.0.1").Times(1)
		p.EXPECT().PromptString("Enter mysql port (default:3306): ").Return("3636").Times(1)
		p.EXPECT().PromptString("Enter database name: ").Return("auth_db").Times(1)
		p.EXPECT().PromptString("Enter mysql username: ").Return("gendutski").Times(1)
		p.EXPECT().PromptPassword("Enter mysql password: ").Return("trial123!").Times(1)
		p.EXPECT().PromptString("Parse mysql time to local (default:false) (true or false): ").Return("true").Times(1)

		result, err := dotenvgenerator.GenerateDotEnv(config2{}, false, false, p)
		assert.Nil(t, err)
		assert.Equal(t, []string{
			"MYSQL_HOST=\"127.0.0.1\"",
			"MYSQL_PORT=\"3636\"",
			"MYSQL_DB_NAME=\"auth_db\"",
			"MYSQL_USERNAME=\"gendutski\"",
			"MYSQL_PASSWORD=\"trial123!\"",
			"MYSQL_PARSE_TIME=\"true\"",
		}, result)
	})

	t.Run("slice struct", func(t *testing.T) {
		// config1
		p.EXPECT().PromptString("Enter port to serve http (default:8080): ").Return("1234").Times(1)
		p.EXPECT().PromptString("Enter Login Username: ").Return("firman").Times(1)
		p.EXPECT().PromptPassword("Enter Login Password: ").Return("password123!").Times(1)
		// config2
		p.EXPECT().PromptString("Enter mysql host (default:localhost): ").Return("127.0.0.1").Times(1)
		p.EXPECT().PromptString("Enter mysql port (default:3306): ").Return("3636").Times(1)
		p.EXPECT().PromptString("Enter database name: ").Return("auth_db").Times(1)
		p.EXPECT().PromptString("Enter mysql username: ").Return("gendutski").Times(1)
		p.EXPECT().PromptPassword("Enter mysql password: ").Return("trial123!").Times(1)
		p.EXPECT().PromptString("Parse mysql time to local (default:false) (true or false): ").Return("true").Times(1)

		result, err := dotenvgenerator.GenerateDotEnv([]interface{}{&config1{}, config2{}}, false, false, p)
		assert.Nil(t, err)
		assert.Equal(t, []string{
			"#config1",
			"HTTP_PORT=\"1234\"",
			"DEFAULT_USERNAME=\"firman\"",
			"DEFAULT_PASSWORD=\"password123!\"",
			"",
			"#config2",
			"MYSQL_HOST=\"127.0.0.1\"",
			"MYSQL_PORT=\"3636\"",
			"MYSQL_DB_NAME=\"auth_db\"",
			"MYSQL_USERNAME=\"gendutski\"",
			"MYSQL_PASSWORD=\"trial123!\"",
			"MYSQL_PARSE_TIME=\"true\"",
		}, result)
	})

	t.Run("combine slice struct & non struct", func(t *testing.T) {
		// config1
		p.EXPECT().PromptString("Enter port to serve http (default:8080): ").Return("1234").Times(1)
		p.EXPECT().PromptString("Enter Login Username: ").Return("firman").Times(1)
		p.EXPECT().PromptPassword("Enter Login Password: ").Return("password123!").Times(1)

		result, err := dotenvgenerator.GenerateDotEnv([]interface{}{&config1{}, []int{10}[0]}, false, false, p)
		assert.NotNil(t, err)
		assert.Equal(t, []string{
			"#config1",
			"HTTP_PORT=\"1234\"",
			"DEFAULT_USERNAME=\"firman\"",
			"DEFAULT_PASSWORD=\"password123!\"",
		}, result)
	})

	t.Run("not struct", func(t *testing.T) {
		result, err := dotenvgenerator.GenerateDotEnv("invalid type", false, false, p)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
