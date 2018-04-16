package network

import (
	"github.com/shirou/gopsutil/net"
)

// GetInterfaces .
func GetInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

// GetConnections .
func GetConnections() ([]net.ConnectionStat, error) {
	return net.Connections("")
}
