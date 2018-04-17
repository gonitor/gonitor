package network

import (
	"github.com/shirou/gopsutil/net"
)

// ServiceGetInterfaces .
func ServiceGetInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

// ServiceGetConnections .
func ServiceGetConnections() ([]net.ConnectionStat, error) {
	return net.Connections("")
}
