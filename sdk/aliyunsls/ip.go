package aliyunsls

import "net"

//LocalIP 获取本机网卡IP
func LocalIP() (ipv4 string) {
	var (
		err   error
		addrs []net.Addr
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr := range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet := addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() == nil {
				continue
			}
			ipv4 = ipNet.IP.String() // 192.168.1.1
			return
		}
	}

	if err != nil {
		ipv4 = "127.0.0.1"
	}
	return
}
