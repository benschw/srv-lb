package lb

import "github.com/benschw/srv-lb/dns"

type StrategyType string
type StrategyFactory func(dns.Lookup) GenericLoadBalancer

var strategyReg map[StrategyType]StrategyFactory

func RegisterStrategy(strategy StrategyType, factory StrategyFactory) {
	strategyReg[strategy] = factory
}

// Load balancer that can service request for any SRV record address
type GenericLoadBalancer interface {
	Next(name string) (dns.Address, error)
}

func NewGeneric(cfg *Config) GenericLoadBalancer {
	if _, ok := strategyReg[cfg.Strategy]; !ok {
		panic("Unknown load balancer strategy")
	}
	return strategyReg[cfg.Strategy](cfg.Dns)
}

func init() {
	strategyReg = make(map[StrategyType]StrategyFactory)
	RegisterStrategy(RoundRobinStrategy, NewRoundRobinStrategy)
}
