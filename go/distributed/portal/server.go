package portal

import (
	"distributed/grades"
	"distributed/registry"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.Handle("/", http.RedirectHandler("/students", http.StatusPermanentRedirect))
	http.HandleFunc("/students", portalHandler)

}

func portalHandler(w http.ResponseWriter, r *http.Request) {
	renderStudents(w, r)
}

func renderStudents(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("render students err: ", err)
		}
	}()

	serviceURL, err := registry.GetProvider("Grade Service")
	if err != nil {
		return
	}
	resp, err := http.Get(serviceURL + "/students")
	if err != nil {
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var students grades.Students
	err = json.Unmarshal(data, &students)
	if err != nil {
		return
	}
	fmt.Printf("RECV: +v\n", students)
}
