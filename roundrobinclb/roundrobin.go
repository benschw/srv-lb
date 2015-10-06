package roundrobinclb

import "github.com/benschw/srv-lb/dns"

func New(lib dns.Lookup) *RoundRobinClb {
	return &RoundRobinClb{
		dnsLib: lib,
		cache:  NewCache(),
	}
}

type RoundRobinClb struct {
	dnsLib dns.Lookup
	cache  *ResultCache
}

func (lb *RoundRobinClb) Next(name string) (dns.Address, error) {
	var add dns.Address

	srvs, err := lb.dnsLib.LookupSRV(name)
	if err != nil {
		return add, err
	}

	srv, err := lb.cache.Next(name, srvs)
	if err != nil {
		return add, err
	}

	ip, err := lb.dnsLib.LookupA(srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}
