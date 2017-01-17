package messageHandler

import (
	"strconv"
	"strings"
	"errors"
	"net"
	"log"
)

func GetID() (byte, error) {
	ip, err := GetLocalIP()
	if err != nil {
		return byte(0), err
	} else {
		ipString := strings.Split(ip, ".")
		id, err := strconv.Atoi(ipString[3])
		if err != nil {
			log.Println(err)
		}
		return byte(id), nil 	
	}
}	

func GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
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
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Cannot resolve IP address. Please check your network connection.")
}