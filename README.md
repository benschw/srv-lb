[![Build Status](https://drone.io/github.com/benschw/dns-clb-go/status.png)](https://drone.io/github.com/benschw/dns-clb-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/dns-clb?status.png)](http://godoc.org/github.com/benschw/dns-clb-go)

# DNS Client Load Balancer for Go

Selects a `SRV` record answer according to specified load balancer algorithm, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


## Example:
	
	srvName := "foo.service.fligl.io"
	lb := New(srvName)

	address, err := lb.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

- Uses dns server configured in `/etc/resolv.conf`
- Uses round robin strategy

### or configure explicitely

	srvName := "foo.service.fligl.io"
	lbDriver := NewDriver(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	})
	lb := &SRVLoadBalancer{Lb: lbDriver, Address: srvName}

	address, err := lb.Next()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001



## Development
tests are run against some fixture dns entries I set up on fligl.io (`dig foo.service.fligl.io SRV`).

	go get -u -t
	go test ./...

	


