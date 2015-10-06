package lb

import (
	"fmt"
	"sort"
	"testing"

	"github.com/benschw/srv-lb/dns"
	"github.com/stretchr/testify/assert"
)

type ByConnectionString []dns.Address

func (a ByConnectionString) Len() int      { return len(a) }
func (a ByConnectionString) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByConnectionString) Less(i, j int) bool {
	if a[i].Address == a[j].Address {
		return a[i].Port < a[j].Port
	}
	return a[i].Address < a[j].Address
}

// Example load balancer with defaults
func ExampleNew() {
	lb := New(DefaultConfig(), "foo.service.fligl.io")

	add1, err := lb.Next()
	if err != nil {
		panic(err)
	}
	add2, err := lb.Next()
	if err != nil {
		panic(err)
	}

	adds := []dns.Address{add1, add2}
	sort.Sort(ByConnectionString(adds))

	fmt.Printf("%s\n%s", adds[0], adds[1])
	// Output:
	// 0.1.2.3:8001
	// 4.5.6.7:8002
}

// Example of using a generic load balancer with custom configuration
func ExampleNewGeneric() {
	srvName := "foo.service.fligl.io"
	lb := NewGeneric(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	})

	add1, err := lb.Next(srvName)
	if err != nil {
		panic(err)
	}
	add2, err := lb.Next(srvName)
	if err != nil {
		panic(err)
	}

	adds := []dns.Address{add1, add2}
	sort.Sort(ByConnectionString(adds))

	fmt.Printf("%s\n%s", adds[0], adds[1])
	// Output:
	// 0.1.2.3:8001
	// 4.5.6.7:8002
}

func TestRoundRobinFacade(t *testing.T) {
	//given
	c := NewGeneric(&Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: RoundRobin,
	})

	// when
	srvName := "foo.service.fligl.io"
	add1, err1 := c.Next(srvName)
	add2, err2 := c.Next(srvName)

	// then
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	adds := []dns.Address{add1, add2}
	sort.Sort(ByConnectionString(adds))

	expected := []dns.Address{dns.Address{Address: "0.1.2.3", Port: 8001}, dns.Address{Address: "4.5.6.7", Port: 8002}}
	assert.Equal(t, expected, adds, "unexpected results")
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
	assert.Nil(t, err)
}
