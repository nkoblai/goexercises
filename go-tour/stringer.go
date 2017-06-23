package stringer

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr type represents ip-adress.
type IPAddr [4]byte

var hosts = map[string]IPAddr{
	"loopback":  {127, 0, 0, 1},
	"googleDNS": {8, 8, 8, 8},
}

// String returns address as a dotted quad.
func (ip IPAddr) String() string {
	str := fmt.Sprintf("%v", ip[0])
	for _, b := range ip[1:] {
		str += fmt.Sprintf(".%v", b)
	}
	return str
}

// Convert returns address as a dotted quad.
func (ip IPAddr) Convert() string {
	s := make([]string, len(ip))
	for i := range ip {
		s[i] = strconv.Itoa(int(ip[i]))
	}
	return strings.Join(s, ".")
}

// TODO: Add a "String() string" method to IPAddr.
func main() {
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		fmt.Printf("%v: %v\n", name, IPAddr.Convert(ip))
	}
}
