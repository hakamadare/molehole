package ip

import (
	"net"

	"github.com/rdegges/go-ipify"
)

var (
	ValidAddrs []net.IP = []net.IP{
		net.ParseIP("52.55.7.50"),
	}
)

func GetIP() (string, error) {
	return ipify.GetIp()
}
