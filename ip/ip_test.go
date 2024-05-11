package ip

import (
	"fmt"
	"testing"
)

func TestIp2long(t *testing.T) {
	ip := "192.168.1.1"
	fmt.Println(Ip2long(ip))
	fmt.Println(Ipv4ToLong(ip))
	fmt.Println(LongToIpv4(3232235777))
}
