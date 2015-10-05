package clb

import (
	"github.com/benschw/dns-clb/dns"
	"github.com/benschw/dns-clb/randomclb"
	"github.com/benschw/dns-clb/roundrobinclb"
)

type LoadBalancer interface {
	GetAddress(name string) (dns.Address, error)
}

func New(cfg *Config) (LoadBalancer, error) {
	return buildClb(cfg.Dns, cfg.LoadBalancerType)
}

func buildClb(lib dns.Lookup, lbType LoadBalancerType) (LoadBalancer, error) {
	switch lbType {
	case RoundRobin:
		return roundrobinclb.New(lib), nil
	case Random:
		return randomclb.New(lib), nil
	}
	panic("Unknown load balancer strategy")
}
