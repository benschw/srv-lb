package first

import (
	"github.com/benschw/srv-lb/dns"
	"github.com/benschw/srv-lb/lb"
)

const FirstStrategy lb.StrategyType = "first"

func New(lib dns.Lookup) lb.GenericLoadBalancer {
	return &FirstClb{lib}
}

type FirstClb struct {
	dnsLib dns.Lookup
}

func (lb *FirstClb) Next(name string) (dns.Address, error) {
	var add dns.Address

	srvs, err := lb.dnsLib.LookupSRV(name)
	if err != nil {
		return add, err
	}

	ip, err := lb.dnsLib.LookupA(srvs[0].Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srvs[0].Port}, nil
}

func init() {
	lb.RegisterStrategy(FirstStrategy, New)
}
