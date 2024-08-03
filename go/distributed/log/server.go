package log

import (
	"fmt"
	"io"
	stlog "log"
	"net/http"
	"os"
)

var logger *stlog.Logger

// Start or initialize
func Run(destination string) {
	logger = stlog.New(Logfile(destination), "mylogger:", stlog.LstdFlags)
}

type Logfile string

func (f Logfile) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(f), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Every service defines a function to register itself as a handler
func RegisterHandlers() {
	http.HandleFunc("/log", logHandler)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log recv")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	msg, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Println(string(msg))
}
