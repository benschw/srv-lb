[![Build Status](https://img.shields.io/codeship/b556c2e0-4dc7-0133-eaf7-524cf6105349.svg)](https://codeship.com/projects/106694)
[![GoDoc](http://godoc.org/github.com/benschw/srv-lb?status.png)](http://godoc.org/github.com/benschw/srv-lb/lb)


# SRV Record Load Balancer library for Go

`SRV-Lb` is a load balancer designed for use with service discovery solutions
that expose a discovery interface of DNS SRV records
(e.g. [consul](https://consul.io/) or [skyDNS](https://github.com/skynetservices/skydns))


The library selects a `SRV` record answer according to specified load balancer algorithm,
resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


To select a DNS server you can us the value from your system's `resolv.conf` (the default),
specify it explicitely when configuring the library,
or set it as an ENV variable (e.g. `SRVLB_HOST=127.0.0.1:8600` to connect to a local consul agent) at run time.


The library defaults to use a "Round Robin" algorithm, but you can specify another or build your own (see below).


## Example:
### Default Load Balancer

	srvName := "foo.service.fligl.io"
	l := lb.New(lb.DefaultConfig(), srvName)

	address, err := l.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

- Uses dns server configured in `/etc/resolv.conf`
- Uses round robin strategy


### or build a generic load balancer

	srvName := "foo.service.fligl.io"
	l := lb.NewGeneric(lb.DefaultConfig())

	address, err := l.Next(srvName)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

### or configure explicitly

	srvName := "foo.service.fligl.io"
	cfg := &lb.Config{
		Dns:      dns.NewLookupLib("127.0.0.1:8600"),
		Strategy: random.RandomStrategy,
	}
	l := lb.New(cfg, srvName)

	address, err := l.Next()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001



## Development
tests are run against some fixture dns entries I set up on fligl.io (`dig foo.service.fligl.io SRV`).

	go get -u -t ./...
	go test ./...

	
## Build your own load balancing strategy

`srv-lb` leverages go's `init()` function to allow you to use your own
load balancer strategy without forking the core library. Below is a walkthrough
of how to create your own "FancyLB" strategy. For a complete example,
[see how the "random" strategy is implemented](https://github.com/benschw/srv-lb/blob/master/strategy/random/random.go).

_The default strategy, `RoundRobin`, is registered slightly differently to avoid import cycles, so avoid using it as an example_


Give your strategy a unique identifier

	const FancyStrategy lb.StrategyType = "fancy"

Create a factory (and your implementation of `GenericLoadBalancer`)

	func New(lib dns.Lookup) lb.GenericLoadBalancer {
		return &FancyLB{Dns: lib}
	}

Register it with the load balancer

	func init() {
		lb.RegisterStrategy(FancyStrategy, New)
	}


And then specify it when constructing your load balancer

	cfg := lb.DefaultConfig()
	cfg.Strategy = fancy.FancyStrategy
	
	l := lb.New(cfg, srvName)

