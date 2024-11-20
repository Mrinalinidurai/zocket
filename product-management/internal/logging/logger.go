package logging

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

func LogRequest(path string, method string, status int) {
	logger.WithFields(logrus.Fields{
		"method": method,
		"path":   path,
		"status": status,
	}).Info("Request processed")
}

func LogError(err error, context string) {
	logger.WithFields(logrus.Fields{
		"context": context,
		"error":   err.Error(),
	}).Error("Error occurred")
}
