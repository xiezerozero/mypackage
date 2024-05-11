package ip

import (
	"encoding/binary"
	"math"
	"net"
)

func Ip2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

func Ipv4ToLong(ipAddress string) uint {
	p := net.ParseIP(ipAddress).To4()
	if p == nil {
		return 0
	}
	return uint(p[0])<<24 | uint(p[1])<<16 | uint(p[2])<<8 | uint(p[3])
}

func LongToIpv4(i uint) string {
	if i > math.MaxUint32 {
		return ""
	}
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)
	return ip.String()
}
