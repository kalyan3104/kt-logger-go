package pipes

import (
	"os"
	"testing"

	"github.com/Dharitri-org/me-core/marshal"
	"github.com/stretchr/testify/require"
)

func TestParentMessenger_ReadLogLine(t *testing.T) {
	logsReader, logsWriter, err := os.Pipe()
	require.Nil(t, err)
	profileReader, profileWriter, err := os.Pipe()
	require.Nil(t, err)

	parentMessenger := NewParentMessenger(logsReader, profileWriter, &marshal.JsonMarshalizer{})
	childMessenger := NewChildMessenger(profileReader, logsWriter)

	_, _ = childMessenger.SendLogLine([]byte("{}"))
	logLine, err := parentMessenger.ReadLogLine()
	require.Nil(t, err)
	require.NotNil(t, logLine)

	_, _ = childMessenger.SendLogLine([]byte(`{"LoggerName": "foo", "Message": "bar"}`))
	logLine, err = parentMessenger.ReadLogLine()
	require.Nil(t, err)
	require.NotNil(t, logLine)
	require.Equal(t, logLine.LoggerName, "foo")
	require.Equal(t, logLine.Message, "bar")
}

func TestParentMessenger_ReadLogLine_BadJsonShouldErrWithActualJson(t *testing.T) {
	logsReader, logsWriter, err := os.Pipe()
	require.Nil(t, err)
	profileReader, profileWriter, err := os.Pipe()
	require.Nil(t, err)

	parentMessenger := NewParentMessenger(logsReader, profileWriter, &marshal.JsonMarshalizer{})
	childMessenger := NewChildMessenger(profileReader, logsWriter)

	_, _ = childMessenger.SendLogLine([]byte("bad json"))
	logLine, err := parentMessenger.ReadLogLine()
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "bad json")
	require.Nil(t, logLine)
}
