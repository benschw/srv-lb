package lb

import (
	"os"

	"github.com/benschw/srv-lb/dns"
)

type Config struct {
	Dns      dns.Lookup
	Strategy StrategyType
}

func DefaultConfig() *Config {
	var dnsLib dns.Lookup

	if dnsHost := os.Getenv("SRVLB_HOST"); dnsHost != "" {
		dnsLib = dns.NewLookupLib(dnsHost)
	} else {
		dnsLib = dns.NewDefaultLookupLib()
	}

	return &Config{
		Dns:      dnsLib,
		Strategy: RoundRobinStrategy,
	}
}
