package lb

import (
	"fmt"
	"sort"

	"github.com/benschw/srv-lb/dns"
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

// Example of using a generic load balancer
func ExampleNewGeneric() {
	srvName := "foo.service.fligl.io"
	lb := NewGeneric(DefaultConfig())

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
