package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

const componentName = "log"

var (
	Log *logrus.Logger
)

func init() {

	logger := logrus.New()
	logger.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	Log = logger

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://ea3b7e2dc92a4d989cdb17e3e098c140@sentry.io/1778424",
	}); err != nil {
		LogWarn(componentName, "Sentry initialization failed: ", err)
		//LogWithComponent("sentry").Warning("Sentry initialization failed: %v\n", err)
	}

}

func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func LogError(componentName string, msg string, err error) {
	LogWithComponent(componentName).Error(msg, err)
	LogToSentry(err)
}

func LogPanic(componentName string, msg string, err error) {
	LogWithComponent(componentName).Panic(msg, err)
	LogToSentry(err)
}

func LogWarn(componentName string, msg string, args ...interface{}) {

	LogWithComponent(componentName).Warn(msg, args)

}

func LogInfo(componentName string, msg string) {
	LogWithComponent(componentName).Infoln(msg)
}

func LogFatal(componentName string, msg string, err error) {
	LogWithComponent(componentName).Fatalln(msg, err)
}

func LogWithComponent(componentName string) *logrus.Entry {

	return Log.WithField("component", componentName)

}

func LogToSentry(err error) {
	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}
