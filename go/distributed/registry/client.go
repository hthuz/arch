package registry

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var prov = provider{
	services: make(map[string]string, 0),
}

func RegisterService(r Registration) error {

	// Listen on Heartbeat url
	heartbeatURL, err := url.Parse(r.HeartbeatURL)
	if err != nil {
		return err
	}
	http.HandleFunc(heartbeatURL.Path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Listen on serviceUpdate url
	serviceUpdateURL, err := url.Parse(r.ServiceUpdateURL)
	if err != nil {
		return err
	}
	http.HandleFunc(serviceUpdateURL.Path, serviceUpdateHandler)

	buf := new(bytes.Buffer)
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = buf.Write(data)
	if err != nil {
		return fmt.Errorf("buf write error: %w", err)
	}

	resp, err := http.Post("http://localhost:5000/registry", "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail to register")
	}

	return nil
}

func UnregisterService(url string) error {

	buf := new(bytes.Buffer)
	_, err := buf.Write([]byte(url))
	if err != nil {
		return fmt.Errorf("buf write error: %w", err)
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/registry", buf)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("fail to unregister")
	}
	return nil
}

func serviceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var p patch
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prov.ApplyPatch(p)
}

type provider struct {
	services map[string]string // Map of serviceName to serviceURL
}

func (p *provider) ApplyPatch(pat patch) {
	for _, entry := range pat.Added {
		p.services[entry.Name] = entry.URL
	}
	for _, entry := range pat.Removed {
		delete(p.services, entry.Name)
	}
}

// Get the URL of provider service
func GetProvider(serviceName string) (string, error) {
	url, ok := prov.services[serviceName]
	if !ok {
		return "", fmt.Errorf("No providers found for %v", serviceName)
	}
	return url, nil
}
