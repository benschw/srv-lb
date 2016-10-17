package lb

import (
	"os"

	"github.com/benschw/srv-lb/dns"
)

type Config struct {
	Dns      dns.Lookup
	Strategy StrategyType
}

func DefaultConfig() (*Config, error) {
	var dnsLib dns.Lookup

	if dnsHost := os.Getenv("SRVLB_HOST"); dnsHost != "" {
		dnsLib = dns.NewLookupLib(dnsHost)
	} else {
		lib, err := dns.NewDefaultLookupLib()
		if err != nil {
			return nil, err
		}
		dnsLib = lib
	}

	return &Config{
		Dns:      dnsLib,
		Strategy: RoundRobinStrategy,
	}, nil
}
