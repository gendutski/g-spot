package logger

import (
	"errors"
	"testing"

	"github.com/gendutski/g-spot/gelo"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	testLog, hook := test.NewNullLogger()

	t.Run("info log, log mode all", func(t *testing.T) {
		log := Logger{
			log:     testLog,
			uri:     "/testing",
			method:  "GET",
			message: "Success",
			status:  200,
			err:     nil,
			logMode: LogModeAll,
		}
		log.Log()
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
		assert.Equal(t, "Success", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("info log, non log mode all", func(t *testing.T) {
		log := Logger{
			log:     testLog,
			uri:     "/testing",
			method:  "GET",
			message: "Success",
			status:  200,
			err:     nil,
		}
		log.Log()
		assert.Equal(t, 0, len(hook.Entries))
		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("warn log, log mode all", func(t *testing.T) {
		log := Logger{
			log:     testLog,
			uri:     "/testing",
			method:  "GET",
			message: "error 400",
			status:  400,
			err:     errors.New("some error happened"),
			logMode: LogModeAll,
		}
		log.Log()
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
		assert.Equal(t, "error 400", hook.LastEntry().Message)
		assert.Equal(t, "some error happened", hook.LastEntry().Data["error"])

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("warn log, log mode warn", func(t *testing.T) {
		log := Logger{
			log:    testLog,
			uri:    "/testing",
			method: "GET",
			status: 400,
			err: &echo.HTTPError{
				Message:  "error 400",
				Internal: gelo.Init(400001, "some error happened"),
			},
			logMode: LogModeAll,
		}
		log.Log()
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
		assert.Equal(t, "error 400", hook.LastEntry().Message)
		assert.Equal(t, 400001, hook.LastEntry().Data["code"])
		assert.Equal(t, "some error happened", hook.LastEntry().Data["error"])

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("warn log, non log mode warn & all", func(t *testing.T) {
		log := Logger{
			log:    testLog,
			uri:    "/testing",
			method: "GET",
			status: 400,
			err: &echo.HTTPError{
				Message:  "error 400",
				Internal: gelo.Init(400001, "some error happened"),
			},
		}
		log.Log()
		assert.Equal(t, 0, len(hook.Entries))
		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("error log, non echo http error", func(t *testing.T) {
		log := Logger{
			log:     testLog,
			uri:     "/testing",
			method:  "GET",
			message: "error 500",
			status:  500,
			err:     errors.New("some error happened"),
		}
		log.Log()
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "error 500", hook.LastEntry().Message)
		assert.Equal(t, "some error happened", hook.LastEntry().Data["error"])

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("error log, echo http error", func(t *testing.T) {
		log := Logger{
			log:    testLog,
			uri:    "/testing",
			method: "GET",
			status: 500,
			err: &echo.HTTPError{
				Message:  "error 500",
				Internal: gelo.Init(500001, "some error happened"),
			},
			logMode: LogModeAll,
		}
		log.Log()
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "error 500", hook.LastEntry().Message)
		assert.Equal(t, 500001, hook.LastEntry().Data["code"])
		assert.Equal(t, "some error happened", hook.LastEntry().Data["error"])

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

}
