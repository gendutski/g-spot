package logger

import (
	"fmt"
	"net/http"

	"github.com/gendutski/g-spot/gelo"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LogMode int

const (
	LogModeErrorOnly LogMode = iota + 1
	LogModeErrorAndWarnOnly
	LogModeAll
)

type Logger struct {
	log         *logrus.Logger
	uri         string
	method      string
	message     string
	status      int
	err         error
	logMode     LogMode
	successHook logrus.Hook
	warnHook    logrus.Hook
	errorHook   logrus.Hook
}

func Init(uri, method, message string, status int, err error, logMode LogMode, successHook, warnHook, errorHook logrus.Hook) Logger {
	glog := logrus.New()
	glog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return Logger{
		log:         glog,
		uri:         uri,
		method:      method,
		status:      status,
		message:     message,
		err:         err,
		logMode:     logMode,
		successHook: successHook,
		warnHook:    warnHook,
		errorHook:   errorHook,
	}
}

func (e Logger) Log() {
	logrusFields := logrus.Fields{
		"URI":    e.uri,
		"Method": e.method,
		"Status": e.status,
		"Error":  e.err,
	}

	if e.err == nil {
		if e.logMode != LogModeAll {
			return
		}
		delete(logrusFields, "Error")

		if e.successHook != nil {
			e.log.AddHook(e.successHook)
		}
		e.log.WithFields(logrusFields).Info(e.message)
	} else {
		// set message
		var message interface{} = fmt.Sprint(e.message)
		if _err, ok := e.err.(*echo.HTTPError); ok {
			message = _err.Message
			logrusFields["error"] = _err.Internal.Error()
			if cErr, ok := _err.Internal.(*gelo.Error); ok {
				logrusFields["code"] = cErr.Code
			}
		} else {
			logrusFields["error"] = e.err.Error()
		}

		if e.status < http.StatusInternalServerError {
			if e.logMode != LogModeErrorAndWarnOnly && e.logMode != LogModeAll {
				return
			}
			if e.warnHook != nil {
				e.log.AddHook(e.warnHook)
			}
			e.log.WithFields(logrusFields).Warn(message)
		} else {
			if e.errorHook != nil {
				e.log.AddHook(e.errorHook)
			}
			e.log.WithFields(logrusFields).Error(message)
		}
	}
}
