package srvlb

import (
	"github.com/benschw/srv-lb/dns"
	"github.com/benschw/srv-lb/randomclb"
	"github.com/benschw/srv-lb/roundrobinclb"
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
