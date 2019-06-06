package config

import "github.com/kelseyhightower/envconfig"

//Config represents configuration for
//this app
type Config struct {
	NumWorkers int32  `envconfig:"num_workers" default:"2"`
	QueueName  string `envconfig:"queue_name" default:"hello"`
}

// Get is used to get configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
