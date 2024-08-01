package registry

import (
	"errors"
	"fmt"
)

type Registration struct {
	ServiceName string
	ServiceHost string
	ServicePort string
}

func (r *Registration) GetURL() string {
	return fmt.Sprintf("http://%s:%s", r.ServiceHost, r.ServicePort)
}

type registry struct {
	registrations []Registration
}

var registryInst = registry{
	registrations: make([]Registration, 0),
}

func (r *registry) Add(reg Registration) {
	r.registrations = append(r.registrations, reg)
}
func (r *registry) Remove(url string) error {
	for i := range r.registrations {
		if r.registrations[i].GetURL() == url {
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			return nil
		}
	}
	return errors.New("serviceURL not found")
}
