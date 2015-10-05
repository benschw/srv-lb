package roundrobinclb

import (
	"fmt"
	"log"
	"testing"

	"github.com/benschw/dns-clb/dns"
	"github.com/stretchr/testify/assert"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestRoundRobinLookup(t *testing.T) {
	// given
	srvName := "foo.service.fligl.io"

	lib := dns.NewDefaultLookupLib()
	c := New(lib)

	// when
	address, err := c.GetAddress(srvName)

	// then
	assert.Nil(t, err)

	if address.Port == 8001 && address.Address == "0.1.2.3" {
		return
	} else if address.Port == 8002 && address.Address == "4.5.6.7" {
		return
	} else {
		t.Errorf("port '%d' not expected with address: '%s'", address.Port, address.Address)
	}

}
