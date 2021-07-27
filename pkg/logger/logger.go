package logger

import (
	"io"
	"os"
	"strings"

	"github.com/bino1490/crypto-svc/pkg/config"

	"github.com/sirupsen/logrus"
)

// BootstrapLogger is the logger handle for logging service
// bootstrap events.
var BootstrapLogger *logrus.Logger

// AccessLogger is the logger handle for logging service
// access events.
var AccessLogger *logrus.Logger

// Logger is the logger handle for logging all service events.
var Logger *logrus.Logger

// timeFormat is the format of timestamp to log on all events.
var timeFormat = "2006-01-02 15:04:05"

// init initializes all the logger handlers during service startup.
// This will be executed only once at the start of service.
func init() {

	// initBootstrapLogger initializes the logger handle for logging
	// service bootstrap events.
	initBootstrapLogger()

	// initAccessLogger initializes the logger handle for logging
	// service access events.
	initAccessLogger()

	// initLogger initializes the logger handle for logging
	// all service events.
	initLogger()
}

// initBootstrapLogger initializes the logger handle for logging
// service bootstrap events.
func initBootstrapLogger() {
	BootstrapLogger = logrus.New()
	BootstrapLogger.Level = getLogLevel(config.
		SrvConfig.GetString("logging.logfile.bootstrap.loglevel"))
	BootstrapLogger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(
		config.SrvConfig.GetString("logging.logfile.bootstrap.path")+
			strings.ToLower(config.SrvConfig.GetString("logging.logfile.bootstrap.name")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		BootstrapLogger.SetOutput(os.Stdout)
	} else {
		BootstrapLogger.SetOutput(file)
	}
}

// initAccessLogger initializes the logger handle for logging
// service access events.
func initAccessLogger() {
	AccessLogger = logrus.New()
	AccessLogger.Level = getLogLevel(config.SrvConfig.GetString("logging.logfile.access.loglevel"))
	AccessLogger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(
		config.SrvConfig.GetString("logging.logfile.access.path")+
			strings.ToLower(config.SrvConfig.GetString("logging.logfile.access.name")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		AccessLogger.SetOutput(os.Stdout)
	} else {
		AccessLogger.SetOutput(file)
	}
}

// initLogger initializes the logger handle for logging
// all service events.
func initLogger() {
	Logger = logrus.New()
	Logger.Level = getLogLevel(config.SrvConfig.GetString("logging.logfile.service.loglevel"))
	Logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(
		config.SrvConfig.GetString("logging.logfile.service.path")+
			strings.ToLower(config.SrvConfig.GetString("logging.logfile.service.name")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		Logger.SetOutput(os.Stdout)
	} else {
		logwriter := io.MultiWriter(os.Stdout, file)
		Logger.SetOutput(logwriter)
	}
}

// getLogLevel returns the log level to initialize the logger handler.
func getLogLevel(loglevel string) logrus.Level {
	switch loglevel {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.ErrorLevel
	}
}

// LogDebug ..
func LogDebug(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Debug(obj)
}

// LogInfo ..
func LogInfo(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Info(obj)
}

// LogWarning ..
func LogWarning(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Warn(obj)
}

// LogError ..
func LogError(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Error(obj)
}

// LogFatal ..
func LogFatal(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Fatal(obj)
}

// LogPanic ..
func LogPanic(obj interface{}, fields map[string]interface{}) {
	Logger.WithFields(fields).Panic(obj)
}
