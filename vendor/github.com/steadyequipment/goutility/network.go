package goutility

import (
	"net"
)

// GetFirstNonLoopbackIP get first non-loopback ip address of the current machine
func GetFirstNonLoopbackIP() (string, error) {

	// https://gist.github.com/mowings/017c80c188d1024ba3e7
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {

		ip := GetAvailableIPFromAddress(addr)

		// TODO: clarify function name, we're only looking for IPv4s here
		// TODO: clarify function name, we're returning a string, not a net.IP
		ipString := retreiveNonLoopbackV4IPString(ip)
		if ipString != nil {
			return *ipString, nil
		}
	}
	return "", MakeUnableToFindExternalIPAddressError()
}

func retreiveNonLoopbackV4IPString(ip net.IP) *string {

	if ip == nil {
		return nil
	}

	if ip.IsLoopback() {
		return nil
	}

	ipv4 := ip.To4()
	if ipv4 == nil {
		return nil
	}

	ipString := ipv4.String()
	return &ipString
}

// GetAvailableIPFromAddress retrieves the IP from a net.Addr if one is available
func GetAvailableIPFromAddress(address net.Addr) net.IP {

	var ip net.IP
	switch v := address.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	return ip
}
