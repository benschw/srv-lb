package random

import (
	"fmt"
	"math/rand"

	"github.com/benschw/srv-lb/dns"
	"github.com/benschw/srv-lb/lb"
)

const RandomStrategy lb.StrategyType = "random"

func New(lib dns.Lookup) lb.GenericLoadBalancer {
	lb := new(RandomClb)
	lb.dnsLib = lib
	return lb
}

type RandomClb struct {
	dnsLib dns.Lookup
}

func (lb *RandomClb) Next(name string) (dns.Address, error) {
	add := dns.Address{}

	srvs, err := lb.dnsLib.LookupSRV(name)
	if err != nil {
		return add, err
	}
	if len(srvs) == 0 {
		return add, fmt.Errorf("no SRV records found")
	}

	//	log.Printf("%+v", srvs)
	srv := srvs[rand.Intn(len(srvs))]

	ip, err := lb.dnsLib.LookupA(srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}

func init() {
	lb.RegisterStrategy(RandomStrategy, New)
}
