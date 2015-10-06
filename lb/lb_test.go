package lb

import (
	"fmt"
	"testing"

	"github.com/benschw/srv-lb/dns"
)

// Example load balancer with defaults
func ExampleNew() {
	lb := New(DefaultConfig(), "foo.service.fligl.io")

	address, err := lb.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001
}

// Example of using a generic load balancer with custom configuration
func ExampleNewGeneric() {
	srvName := "foo.service.fligl.io"
	lb := NewGeneric(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	})

	address, err := lb.Next(srvName)
	if err != nil {
		fmt.Print(err)
	}

	if address.Port == 8001 {
		fmt.Printf("%s", address)
	} else {
		address2, err := lb.Next(srvName)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s", address2)
	}
	// Output: 0.1.2.3:8001
}

func TestRoundRobinFacade(t *testing.T) {
	//given
	c := NewGeneric(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	})

	// when
	srvName := "foo.service.fligl.io"
	_, err := c.Next(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}

func TestRandomFacade(t *testing.T) {
	//given
	c := NewGeneric(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: Random,
	})

	// when
	srvName := "foo.service.fligl.io"
	_, err := c.Next(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}
