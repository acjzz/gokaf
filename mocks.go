package gokaf

import (
	"bytes"
	"log/slog"
	"regexp"
	"strings"
)

var (
	mockLoggerBuff        bytes.Buffer
	mockLoggerTextHandler = slog.NewTextHandler(&mockLoggerBuff, nil)
	mockLogger            = slog.New(mockLoggerTextHandler)
	mockLogsRegex         = regexp.MustCompile(`msg="([^"]+)"`)
)

func getMockLogs() []string {
	var logs []string

	for _, message := range strings.Split(mockLoggerBuff.String(), "\n") {
		matches := mockLogsRegex.FindStringSubmatch(message)
		if len(matches) >= 2 {
			logs = append(logs, matches[1])
		}
	}
	return logs
}
