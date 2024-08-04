package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewFileHook(filePath string, logLevels []logrus.Level) *FileHook {
	return &FileHook{
		filePath:  filePath,
		logLevels: logLevels,
	}
}

// logrus file hook, will hook log to file
type FileHook struct {
	filePath  string
	logLevels []logrus.Level
}

// Fire write log to file
func (hook *FileHook) Fire(entry *logrus.Entry) error {
	theLog, err := os.OpenFile(hook.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer theLog.Close()

	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = theLog.Write([]byte(line))
	return err
}

// Levels returns the levels logged by this hook
func (hook *FileHook) Levels() []logrus.Level {
	return hook.logLevels
}
