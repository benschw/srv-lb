package clb

import (
	"github.com/benschw/dns-clb/dns"
	"github.com/benschw/dns-clb/randomclb"
	"github.com/benschw/dns-clb/roundrobinclb"
)

type SRVLoadBalancerDriver interface {
	Next(name string) (dns.Address, error)
}

func NewDriver(cfg *Config) SRVLoadBalancerDriver {
	switch cfg.Strategy {
	case RoundRobin:
		return roundrobinclb.New(cfg.Dns)
	case Random:
		return randomclb.New(cfg.Dns)
	}
	panic("Unknown load balancer strategy")
}
