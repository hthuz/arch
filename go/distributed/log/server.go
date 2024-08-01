package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

var logger *stlog.Logger

// Start or initialize
func Run(destination string) {
	logger = stlog.New(os.Stdout, "mylogger:", stlog.LstdFlags)
}

// Every service defines a function to register itself as a handler
func RegisterHandlers() {
	http.HandleFunc("/log", logHandler)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
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
