package native

import "github.com/CyCoreSystems/ari"

// Config provides the ARI Configuration accessors for a native client
type Config struct {
	client *Client
}

// Get gets a lazy handle to a configuration object
func (c *Config) Get(configClass string, objectType string, id string) *ari.ConfigHandle {
	return ari.NewConfigHandle(configClass, objectType, id, c)
}

// Data retrieves the state of a configuration object
func (c *Config) Data(configClass string, objectType string, id string) (cd ari.ConfigData, err error) {
	cd.ID = id
	cd.Class = configClass
	cd.Type = objectType
	err = c.client.conn.Get("/asterisk/config/dynamic/"+configClass+"/"+objectType+"/"+id, &cd.Fields)
	return
}

// Update updates the given configuration object
func (c *Config) Update(configClass string, objectType string, id string, tuples []ari.ConfigTuple) (err error) {
	err = c.client.conn.Put("/asterisk/config/dynamic/"+configClass+"/"+objectType+"/"+id, nil, &tuples)
	return
}

// Delete deletes the configuration object
func (c *Config) Delete(configClass string, objectType string, id string) (err error) {
	err = c.client.conn.Delete("/asterisk/config/dynamic/"+configClass+"/"+objectType+"/"+id, nil, "")
	return
}
