package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	ip1, ip2, ma := "", "", ""
	fmt.Scanf("%s%s%s", &ip1, &ip2, &ma)
	checkIp(ip1, ip2, ma)
}

func checkIp(ips ...string) {
	var ip1, ip2, ma uint32 = 0, 0, 0
	ip1 = getIp(ips[0])
	ip2 = getIp(ips[1])
	ma = getIp(ips[2])
	if ip1&ma == ip2&ma {
		fmt.Println("1 ", getIpStr(ip1&ma))
	} else {
		fmt.Println("0 ", getIpStr(ip1&ma))
	}
}

func getIp(str string) uint32 {
	tmp := strings.Split(str, ".")
	i := 0
	var num uint32
	for j := len(tmp) - 1; j >= 0; j-- {
		n, _ := strconv.Atoi(tmp[j])
		num += uint32(n << uint(i*8))
		i++
	}
	return num
}

func getIpStr(i uint32) string {
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)
	return ip.String()
}
