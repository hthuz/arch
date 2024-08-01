package registry

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {

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
