package main

import (
    "strings"
    "net"
    "fmt"
)

func InternalIP() string {
    inters, err := net.Interfaces()
    if err != nil {
        return ""
    }
    for _, inter := range inters {
        fmt.Println(inter.Name)
        if !strings.HasPrefix(inter.Name, "lo") {
            addrs, err := inter.Addrs()
            if err != nil {
                continue
            }
            fmt.Println(addrs)
            for _, addr := range addrs {
                if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                    if ipnet.IP.To4() != nil {
                        return ipnet.IP.String()
                    }
                }
            }
        }
    }
    return ""
}

func main()  {
    fmt.Println(InternalIP())
}