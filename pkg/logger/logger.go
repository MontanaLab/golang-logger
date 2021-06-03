package logger

type Logger interface {
	Emergency(message string, context map[string]interface{})
	Alert(message string, context map[string]interface{})
	Critical(message string, context map[string]interface{})
	Error(message string, context map[string]interface{})
	Warning(message string, context map[string]interface{})
	Notice(message string, context map[string]interface{})
	Info(message string, context map[string]interface{})
	Debug(message string, context map[string]interface{})
	Log(level LogLevel, message string, context map[string]interface{})
}
