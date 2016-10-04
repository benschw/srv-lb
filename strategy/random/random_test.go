package random

import (
	"fmt"
	"log"
	"testing"

	"github.com/benschw/srv-lb/dns"
	"github.com/stretchr/testify/assert"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

//strconv.FormatInt(int64(srv.Port), 10)
func TestRandomLookup(t *testing.T) {
	lib, err := dns.NewDefaultLookupLib()
	assert.Nil(t, err)

	// given
	srvName := "foo.service.fligl.io"
	c := New(lib)

	// when
	address, err := c.Next(srvName)

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
