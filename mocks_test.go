package gokaf

import (
	"testing"
)

func TestGetMockLogs(t *testing.T) {
	// Clear the buffer before running the test
	mockLoggerBuff.Reset()

	// Log some messages using the mock logger
	mockLogger.Warn("Message 1")
	mockLogger.Warn("Message 2")
	mockLogger.Warn("Message 3")

	// Get the logs using the getMockLogs function
	logs := getMockLogs()

	// Define expected log messages
	expectedLogs := []string{"Message 1", "Message 2", "Message 3"}

	// Check if the number of logs matches
	if len(logs) != len(expectedLogs) {
		t.Errorf("Expected %d logs, but got %d", len(expectedLogs), len(logs))
	}

	// Check if each log message is as expected
	for i, expectedLog := range expectedLogs {
		if logs[i] != expectedLog {
			t.Errorf("Expected log: %s, but got: %s", expectedLog, logs[i])
		}
	}
}

func TestMockLogsRegex(t *testing.T) {
	// Define log message with multiple fields
	logMessage := `time=2023-11-15T23:34:25.522+04:00 level=WARN msg="Engine Shutting Down"`

	// Find submatches using the mockLogsRegex
	matches := mockLogsRegex.FindStringSubmatch(logMessage)

	// Check if the 'msg' field was found
	if len(matches) < 2 {
		t.Errorf("Unable to extract 'msg' field from log message")
	}

	// Check if the extracted 'msg' value is as expected
	expectedMsg := "Engine Shutting Down"
	if matches[1] != expectedMsg {
		t.Errorf("Expected 'msg' value: %s, but got: %s", expectedMsg, matches[1])
	}
}
