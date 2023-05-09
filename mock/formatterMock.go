package mock

import (
	"fmt"
	"strings"
	"time"

	logger "github.com/Dharitri-org/me-core-logger-go"
)

const messageFixedLength = 40
const formatPlainString = "%s[%s] %s %s %s %s\n"

type FormatterMock struct {
}

// Output converts the provided LogLineHandler into a slice of bytes ready for output
func (mock *FormatterMock) Output(line logger.LogLineHandler) []byte {
	if line == nil {
		return nil
	}

	level := logger.LogLevel(line.GetLogLevel())
	timestamp := displayTime(line.GetTimestamp())
	loggerName := ""
	correlation := ""
	message := formatMessage(line.GetMessage())
	args := formatArgsNoAnsi(line.GetArgs()...)

	return []byte(
		fmt.Sprintf(formatPlainString,
			level,
			timestamp, loggerName, correlation,
			message, args,
		),
	)
}

// IsInterfaceNil returns true if there is no value under the interface
func (mock *FormatterMock) IsInterfaceNil() bool {
	return mock == nil
}

func displayTime(timestamp int64) string {
	t := time.Unix(0, timestamp)
	return t.Format("2006-01-02 15:04:05.000")
}

func formatArgsNoAnsi(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	argString := ""
	for index := 1; index < len(args); index += 2 {
		argString += fmt.Sprintf("%s = %s ", args[index-1], args[index])
	}

	return argString
}

func formatMessage(msg string) string {
	return padRight(msg, messageFixedLength)
}

func padRight(str string, maxLength int) string {
	paddingLength := maxLength - len(str)

	if paddingLength > 0 {
		return str + strings.Repeat(" ", paddingLength)
	}

	return str
}
