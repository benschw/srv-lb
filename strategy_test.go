package main

import (
	"sort"
	"testing"

	"github.com/benschw/srv-lb/dns"
	"github.com/benschw/srv-lb/lb"
	"github.com/benschw/srv-lb/strategy/random"
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

func TestRoundRobinStrategy(t *testing.T) {
	//given
	c := lb.NewGeneric(&lb.Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: lb.RoundRobinStrategy,
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

func TestRandomStrategy(t *testing.T) {
	//given
	c := lb.NewGeneric(&lb.Config{
		Dns:      dns.NewDefaultLookupLib(),
		Strategy: random.RandomStrategy,
	})

	// when
	srvName := "foo.service.fligl.io"
	_, err := c.Next(srvName)

	// then
	assert.Nil(t, err)
}
