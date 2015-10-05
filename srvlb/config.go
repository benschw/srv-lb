package srvlb

import (
	"os"

	"github.com/benschw/srv-lb/dns"
)

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
	var dnsLib dns.Lookup

	if dnsHost := os.Getenv("SRVLB_HOST"); dnsHost != "" {
		dnsLib = dns.NewLookupLib(dnsHost)
	} else {
		dnsLib = dns.NewDefaultLookupLib()
	}

	return &Config{
		Dns:      dnsLib,
		Strategy: RoundRobin,
	}
}
