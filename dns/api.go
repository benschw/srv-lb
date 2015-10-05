package dns

import (
	"fmt"
	"net"
)

type ByTarget []net.SRV

func (a ByTarget) Len() int           { return len(a) }
func (a ByTarget) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTarget) Less(i, j int) bool { return a[i].Target < a[j].Target }

type Address struct {
	Address string
	Port    uint16
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%d", a.Address, a.Port)
}
