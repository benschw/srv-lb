package first

import "github.com/benschw/srv-lb/lb"

func Config() (*lb.Config, error) {
	config, err := lb.DefaultConfig()

	if err == nil {
		config.Strategy = FirstStrategy
	}

	return config, err
}
