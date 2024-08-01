package grades

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	http.HandleFunc("/students", gradeHandler)
	http.HandleFunc("/students/", gradeHandler)
}

// /students return all students
// /students/{id} return a certan student
// /students/{id}/gpa
func gradeHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) == 2 {
		data, err := json.Marshal(students)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
		return
	}
	if len(segs) == 3 {
		id, err := strconv.Atoi(segs[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		student, err := students.GetStudentByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}
		data, err := json.Marshal(student)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
		return
	}
	if len(segs) == 4 {
		id, err := strconv.Atoi(segs[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if segs[3] != "gpa" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		student, err := students.GetStudentByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}
		data, err := json.Marshal(student.GetGPA())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
