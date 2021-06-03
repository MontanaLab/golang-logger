package logger

type NullLogger struct{}

func NewLogger() Logger {
	return &NullLogger{}
}

func (l *NullLogger) Log(level LogLevel, message string, context map[string]interface{}) {
	return // Do nothing
}

func (l *NullLogger) Emergency(message string, context map[string]interface{}) {
	l.Log(EMERGENCY, message, context)
}

func (l *NullLogger) Alert(message string, context map[string]interface{}) {
	l.Log(ALERT, message, context)
}

func (l *NullLogger) Critical(message string, context map[string]interface{}) {
	l.Log(CRITICAL, message, context)
}

func (l *NullLogger) Error(message string, context map[string]interface{}) {
	l.Log(ERROR, message, context)
}

func (l *NullLogger) Warning(message string, context map[string]interface{}) {
	l.Log(WARNING, message, context)
}

func (l *NullLogger) Notice(message string, context map[string]interface{}) {
	l.Log(NOTICE, message, context)
}

func (l *NullLogger) Info(message string, context map[string]interface{}) {
	l.Log(INFO, message, context)
}

func (l *NullLogger) Debug(message string, context map[string]interface{}) {
	l.Log(DEBUG, message, context)
}
