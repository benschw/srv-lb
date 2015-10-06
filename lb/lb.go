package lb

import "github.com/benschw/srv-lb/dns"

type LoadBalancer interface {
	Next() (dns.Address, error)
}

func New(cfg *Config, address string) LoadBalancer {
	return &SRVLoadBalancer{
		Lb:      NewGeneric(cfg),
		Address: address,
	}
}

type SRVLoadBalancer struct {
	Lb      SRVLoadBalancerDriver
	Address string
}

func (s *SRVLoadBalancer) Next() (dns.Address, error) {
	return s.Lb.Next(s.Address)
}

type StaticLoadBalancer struct {
	Address dns.Address
}

func (s *StaticLoadBalancer) Next() (dns.Address, error) {
	return s.Address, nil
}
