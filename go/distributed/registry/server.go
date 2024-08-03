package registry

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/registry", registryHandler)
}

func registryHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var reg Registration
		msg, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		err = json.Unmarshal(msg, &reg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		registryInst.Add(reg)
		log.Printf("Adding service %v with URL %s\n", reg.ServiceName, reg.ServiceURL)
	case http.MethodDelete:
		msg, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		serviceURL := string(msg)
		if err = registryInst.Remove(serviceURL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Printf("Removing service with URL %s\n", serviceURL)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
