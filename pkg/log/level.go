package log

// A log level or log severity tells how important a given log message is.
type level int

const (
	//The DEBUG log level should be used for information that may be needed for
	//diagnosing issues and troubleshooting or when running application in the
	//test environment for the purpose of making sure everything is running correctly.
	DEBUG level = iota

	//The INFO log level indicates that something happened, application entered a
	//certain state, etc. The information logged using the INFO log level should be
	//purely informative and not looking into them on a regular basis should not
	//result in missing any important information.
	INFO

	//The WARN log level indicates that something unexpected happened in the application,
	//a problem, or a situation that might disturb one of the processes, but that does not
	//mean that the application failed. The WARN level should be used in situations that
	//are unexpected, but the code can continue to work.
	WARN

	//The ERROR log level should be used when the application hits an issue preventing
	//one or more functionalities from functioning properly.
	ERROR

	//The FATAL log level tells that the application encountered an event or entered
	//a state in which one of the crucial business functionality is no longer working.
	FATAL
)

// The String method returns a string corresponding to log level.
func (l level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	normal = 37
)

// The color method returns the formatting color associated with the level.
func (l level) color() uint {
	switch l {
	case DEBUG:
		return normal
	case ERROR, FATAL:
		return red
	case WARN:
		return yellow
	case INFO:
		return blue
	default:
		return normal
	}
}
