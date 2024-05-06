package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run scanner.go <ip_address/subnet>")
		return
	}

	target := os.Args[1]
	ips, err := getAllIPs(target)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	openPorts := make(map[string][]int)

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			ports := scanPorts(ip)
			openPorts[ip] = ports
		}(ip)
	}

	wg.Wait()

	for ip, ports := range openPorts {
		fmt.Printf("%s: Open ports: %v\n", ip, ports)
	}
}

func getAllIPs(target string) ([]string, error) {
	ip := net.ParseIP(target)
	if ip != nil {
		return []string{target}, nil
	}

	_, ipnet, err := net.ParseCIDR(target)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func scanPorts(ip string) []int {
	var openPorts []int
	for port := 1; port <= 65535; port++ {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	sort.Ints(openPorts)
	return openPorts
}
