package log

import (
	"bytes"
	"fmt"
	stlog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientServiceName string) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientServiceName))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	buf := bytes.NewBuffer(data)
	resp, err := http.Post(cl.url, "text/plain", buf)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to log message, get status code %v", resp.StatusCode)
	}
	return len(data), nil
}
