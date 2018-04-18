package network

import (
	"github.com/shirou/gopsutil/net"
)

// ServiceGetInterfaces gets the network interaces.
func ServiceGetInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

// ServiceGetConnections gets the network connections.
func ServiceGetConnections() ([]net.ConnectionStat, error) {
	return net.Connections("")
}
