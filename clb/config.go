package clb

import "github.com/benschw/dns-clb/dns"

type LoadBalancerStrategy int

const (
	Random     LoadBalancerStrategy = iota
	RoundRobin LoadBalancerStrategy = iota
)

type Config struct {
	Dns      dns.Lookup
	Strategy LoadBalancerStrategy
}

func DefaultConfig() *Config {
	return &Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	}
}