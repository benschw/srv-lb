package dns

import (
	"net"
	"sort"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

func TestLookupShouldFailWithBadNS(t *testing.T) {
	lib := NewLookupLib("foo:9999")

	_, err := lib.LookupA("foo")

	assert.NotNil(t, err)
}

func TestLookupShouldFailWithBadHost(t *testing.T) {
	lib := NewLookupLib("8.8.8.8:53")

	_, err := lib.LookupA("foo")

	assert.NotNil(t, err)
}

func TestLookupShouldResolveARecord(t *testing.T) {
	lib := NewLookupLib("8.8.8.8:53")

	address, err := lib.LookupA("github.com")

	assert.Nil(t, err)

	ip := net.ParseIP(address)
	assert.NotNil(t, ip.To4())
}

func TestClientConfigLookupA(t *testing.T) {
	lib, err := NewClientConfigLookupLib(&predictableClientConfig{})
	assert.Nil(t, err)

	address, err := lib.LookupA("github.com")

	assert.Nil(t, err)

	ip := net.ParseIP(address)
	assert.NotNil(t, ip.To4())
}

func TestDefaultLookupSRV(t *testing.T) {
	lib, err := NewClientConfigLookupLib(&predictableClientConfig{})
	assert.Nil(t, err)

	addresses, err := lib.LookupSRV("foo.service.fligl.io")

	assert.Nil(t, err)

	sort.Sort(ByTarget(addresses))

	assert.Equal(t, 2, len(addresses), "should be two results")
	assert.Equal(t, "foo1.fligl.io.", addresses[0].Target, "Unexpected Result")
	assert.Equal(t, "foo2.fligl.io.", addresses[1].Target, "Unexpected Result")
}

func TestLookUpShouldSkipBadNS(t *testing.T) {
	lib, err := NewClientConfigLookupLib(&firstNSBadClientConfig{})

	assert.Nil(t, err)

	address, err := lib.LookupA("github.com")

	assert.Nil(t, err)

	ip := net.ParseIP(address)
	assert.NotNil(t, ip.To4())
}

type predictableClientConfig struct{}

func (c *predictableClientConfig) Get() (*dns.ClientConfig, error) {
	return &dns.ClientConfig{
		Servers: []string{"8.8.8.8", "8.8.4.4"},
		Port:    "53",
	}, nil
}

type firstNSBadClientConfig struct{}

func (c *firstNSBadClientConfig) Get() (*dns.ClientConfig, error) {
	return &dns.ClientConfig{
		Servers: []string{"foo", "8.8.8.8"},
		Port:    "53",
	}, nil
}
