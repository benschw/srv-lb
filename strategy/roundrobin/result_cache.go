package roundrobin

import (
	"fmt"
	"net"
	"sort"
	"sync"

	"github.com/benschw/srv-lb/dns"
)

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
