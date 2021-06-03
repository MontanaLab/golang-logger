package logger

type LogLevel string

const (
	EMERGENCY LogLevel = "emergency"
	ALERT     LogLevel = "alert"
	CRITICAL  LogLevel = "critical"
	ERROR     LogLevel = "error"
	WARNING   LogLevel = "warning"
	NOTICE    LogLevel = "notice"
	INFO      LogLevel = "info"
	DEBUG     LogLevel = "debug"
)

func (l LogLevel) GetValue() int {
	var value int
	switch l {
	case EMERGENCY:
		value = 7
		break

	case ALERT:
		value = 6
		break

	case CRITICAL:
		value = 5
		break

	case ERROR:
		value = 4
		break

	case WARNING:
		value = 3
		break

	case NOTICE:
		value = 2
		break

	case INFO:
		value = 1
		break

	case DEBUG:
		value = 0
		break

	default:
		value = -1
		break
	}

	return value
}
