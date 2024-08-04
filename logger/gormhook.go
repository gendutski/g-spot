package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewGormHook(db *gorm.DB, tableName string, levelField string, messageField string) *GormHook {
	return &GormHook{db, tableName, levelField, messageField}
}

type GormHook struct {
	db           *gorm.DB
	tableName    string
	levelField   string
	messageField string
}

func (hook *GormHook) Fire(entry *logrus.Entry) error {
	// insert log to database
	return hook.db.Raw(
		fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES(?, ?)", hook.tableName, hook.levelField, hook.messageField),
		entry.Level,
		entry.Message,
	).Error
}

func (hook *GormHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
