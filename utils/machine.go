package utils

import "net"

func GetMacAddr() string {
	ifas, err := net.Interfaces()
	if err != nil {
		return "-"
	}

	macAddress := ""
	for _, ifa := range ifas {
		addr := ifa.HardwareAddr.String()
		if addr != "" {
			macAddress = macAddress + addr
		}
	}

	return macAddress
}

func GetLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			return ip.String()
		}
	}
	return ""
}
