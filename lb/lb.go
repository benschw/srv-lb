package lb

import (
	"github.com/benschw/srv-lb/dns"
	"github.com/benschw/srv-lb/strategy/random"
	"github.com/benschw/srv-lb/strategy/roundrobin"
)

// Load balancer that can service request for any SRV record address
type GenericLoadBalancer interface {
	Next(name string) (dns.Address, error)
}

func NewGeneric(cfg *Config) GenericLoadBalancer {
	switch cfg.Strategy {
	case RoundRobin:
		return roundrobin.New(cfg.Dns)
	case Random:
		return random.New(cfg.Dns)
	}
	panic("Unknown load balancer strategy")
}

// Load balancer that can service request for a configured SRV record address
type LoadBalancer interface {
	Next() (dns.Address, error)
}

func New(cfg *Config, address string) LoadBalancer {
	return &SRVLoadBalancer{
		Lb:      NewGeneric(cfg),
		Address: address,
	}
}

// Default implementation for load balancing on a SRV record
type SRVLoadBalancer struct {
	Lb      GenericLoadBalancer
	Address string
}

func (s *SRVLoadBalancer) Next() (dns.Address, error) {
	return s.Lb.Next(s.Address)
}

// Specify an address to always return (good for compatibility and in test)
type StaticLoadBalancer struct {
	Address dns.Address
}

func (s *StaticLoadBalancer) Next() (dns.Address, error) {
	return s.Address, nil
}
