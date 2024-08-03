package registry

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Registration struct {
	ServiceName      string
	ServiceHost      string
	ServicePort      string
	ServiceURL       string
	RequiredServices []string
	ServiceUpdateURL string
	HeartbeatURL     string
}

type patchEntry struct {
	Name string
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}

type registry struct {
	registrations []Registration
}

var registryInst = registry{
	registrations: make([]Registration, 0),
}

func (r *registry) Add(reg Registration) error {
	r.registrations = append(r.registrations, reg)
	err := r.sendRequiredServices(reg)
	p := patch{
		Added: []patchEntry{{
			Name: reg.ServiceName,
			URL:  reg.ServiceURL,
		}},
	}
	r.notify(p)
	return err
}
func (r *registry) Remove(url string) error {
	fmt.Println(r.GetRegisteredServiceName())
	for i := range r.registrations {
		if r.registrations[i].ServiceURL == url {
			p := patch{
				Removed: []patchEntry{{
					Name: r.registrations[i].ServiceName,
					URL:  url,
				}},
			}
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.notify(p)
			return nil
		}
	}
	return errors.New("serviceURL not found")
}
func (r *registry) sendRequiredServices(reg Registration) error {
	var p patch
	// For all regisitred service
	for _, registration := range r.registrations {
		// For all reg's required services
		for _, requiredSerivce := range reg.RequiredServices {
			// If the reg's required service is already registered
			if requiredSerivce == registration.ServiceName {
				p.Added = append(p.Added, patchEntry{
					Name: registration.ServiceName,
					URL:  registration.ServiceURL,
				})
			}
		}
	}
	err := r.sendPatch(p, reg.ServiceUpdateURL)
	if err != nil {
		return err
	}
	return nil

}

func (r *registry) sendPatch(p patch, url string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to send patch, get code %v", resp.StatusCode)
	}
	return nil
}

// When service at registry is updated, notify other services, which may
// depend on these updated services.
func (r *registry) notify(fullpatch patch) {
	for _, reg := range r.registrations {
		// Only send relevant services in patches
		go func(reg Registration) {
			var sentPatch patch
			for _, requiredService := range reg.RequiredServices {
				for _, addedService := range fullpatch.Added {
					if requiredService == addedService.Name {
						sentPatch.Added = append(sentPatch.Added, addedService)
					}
				}
				for _, removedServce := range fullpatch.Removed {
					if requiredService == removedServce.Name {
						sentPatch.Removed = append(sentPatch.Removed, removedServce)
					}
				}
			}
			if len(sentPatch.Added) != 0 || len(sentPatch.Removed) != 0 {
				err := r.sendPatch(sentPatch, reg.ServiceUpdateURL)
				if err != nil {
					log.Println(err)
				}
			}

		}(reg)
	}

}

func (r *registry) GetRegisteredServiceName() []string {
	services := []string{}
	for _, reg := range r.registrations {
		services = append(services, reg.ServiceName)
	}
	return services
}

func (r *registry) heartbeat(freq time.Duration) {

	for {
		var wg sync.WaitGroup
		for _, reg := range r.registrations {
			wg.Add(1)
			go func(reg Registration) {
				defer wg.Done()
				for attempt := 0; attempt < 3; attempt++ {
					resp, err := http.Get(reg.HeartbeatURL)
					if err == nil && resp.StatusCode == http.StatusOK {
						log.Printf("heartbeat acked for %v\n", reg.ServiceName)
						return
					}
					time.Sleep(1 * time.Second)
				}
				// heartbeat fail case
				r.Remove(reg.ServiceURL)
				log.Printf("heartbeat lost for %v\n", reg.ServiceName)
			}(reg)
		}
		wg.Wait()
		time.Sleep(freq)
	}

}

var once sync.Once

func SetupRegistryService() {
	once.Do(func() {
		go registryInst.heartbeat(3 * time.Second)
	})
}
