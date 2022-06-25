package utils

import (
	"fmt"
	"net"
)

func GetLocalIP() []string {
	ips := make([]string, 0, 4)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrList, _ := netInterfaces[i].Addrs()
			for _, address := range addrList {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ips = append(ips, ipNet.IP.String())
					}
				}
			}
		}
	}
	return ips
}
