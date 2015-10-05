package clb

import "github.com/benschw/dns-clb/dns"

type LoadBalancer interface {
	Next() (dns.Address, error)
}

func New(address string) LoadBalancer {
	return &SRVLoadBalancer{
		Lb:      NewDriver(DefaultConfig()),
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
