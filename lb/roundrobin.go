package lb

import (
	"fmt"
	"net"
	"sort"
	"sync"

	"github.com/benschw/srv-lb/dns"
)

const RoundRobinStrategy StrategyType = "round-robin"

func NewRoundRobinStrategy(lib dns.Lookup) GenericLoadBalancer {
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

type AddressResult struct {
	srvs []net.SRV
	i    int
}

type ResultCache struct {
	m map[string]*AddressResult
	sync.RWMutex
}

func NewCache() *ResultCache {
	return &ResultCache{
		m: make(map[string]*AddressResult),
	}
}

func (s *ResultCache) Next(key string, srvs []net.SRV) (net.SRV, error) {
	s.Lock()
	defer s.Unlock()

	s.applyResults(key, srvs)

	return s.getNextResult(key)
}

func (s *ResultCache) applyResults(key string, srvs []net.SRV) {
	old, ok := s.m[key]
	if !ok {
		old = &AddressResult{
			srvs: make([]net.SRV, 0),
			i:    0,
		}
		s.m[key] = old
	}
	ordered, changed := s.mergeResults(old.srvs, srvs)
	if changed {
		s.m[key].srvs = ordered
		s.m[key].i = 0
	}
}

func (s *ResultCache) mergeResults(oldResults []net.SRV, newResults []net.SRV) ([]net.SRV, bool) {
	sort.Sort(dns.ByTarget(newResults))

	if len(oldResults) != len(newResults) {
		return newResults, true
	}

	for i, val := range oldResults {
		if newResults[i] != val {
			return newResults, true
		}
	}

	return newResults, false
}

func (s *ResultCache) getNextResult(key string) (net.SRV, error) {
	var srv net.SRV
	srvs := s.m[key].srvs
	i := s.m[key].i

	if len(srvs) == 0 {
		return srv, fmt.Errorf("no results found for '%s'", key)
	}

	if len(srvs) <= i {
		i = 0
	}

	srv = srvs[i]
	s.m[key].i += 1

	return srv, nil
}
