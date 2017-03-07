package native

import (
	"net/url"

	"github.com/CyCoreSystems/ari"
)

type nativeSound struct {
	client *Client
}

// Get returns a managed handle to a SoundData
func (s *nativeSound) Get(name string) *ari.SoundHandle {
	return ari.NewSoundHandle(name, s)
}

// Data returns the details of a given ARI Sound
// Equivalent to GET /sounds/{name}
func (s *nativeSound) Data(name string) (sd ari.SoundData, err error) {
	err = s.client.conn.Get("/sounds/"+name, &sd)
	return sd, err
}

// List returns available sounds limited by the provided filters.
// valid filters are "lang", "format", and nil (no filter)
// An empty filter returns all available sounds
func (s *nativeSound) List(filters map[string]string) (sh []*ari.SoundHandle, err error) {

	var sounds = []struct {
		Name string `json:"name"`
	}{}

	uri := "/sounds"
	if len(filters) > 0 {
		v := url.Values{}
		for key, val := range filters {
			v.Set(key, val)
		}
		uri += "?" + v.Encode()
	}

	err = s.client.conn.Get(uri, &sounds)

	// Store whatever we received, even if incomplete or error
	for _, i := range sounds {
		sh = append(sh, s.Get(i.Name))
	}

	return sh, err
}
