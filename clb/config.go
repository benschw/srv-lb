package clb

import "github.com/benschw/dns-clb/dns"

type LoadBalancerType int

const (
	Random     LoadBalancerType = iota
	RoundRobin LoadBalancerType = iota
)

type Config struct {
	Dns              dns.Lookup
	LoadBalancerType LoadBalancerType
}

func DefaultConfig() *Config {
	return &Config{
		Dns:              dns.NewDefaultLookupLib(),
		LoadBalancerType: RoundRobin,
	}
}
