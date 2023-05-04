package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// GetFreePort 获取可用端口
func GetFreePort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	defer listener.Close()
	if err != nil {
		return 0, err
	}
	_, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(port)
}

// GetFreeIp 获取可用ip
func GetFreeIp() (string, error) {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrList {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && strings.Contains(ipnet.IP.String(), ".") {
			return ipnet.IP.String(), nil
		}
	}
	return "", errors.New("free ip not found")
}

func main() {
	fmt.Println(GetFreeIp())
}
