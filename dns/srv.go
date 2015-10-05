package dns

import "net"

type ByTarget []net.SRV

func (a ByTarget) Len() int           { return len(a) }
func (a ByTarget) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTarget) Less(i, j int) bool { return a[i].Target < a[j].Target }
