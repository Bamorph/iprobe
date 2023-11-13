package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	ipv6Flag := flag.Bool("ipv6", false, "Resolve IPv6 addresses")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		hostname := scanner.Text()

		// Resolve the hostname to IP addresses based on the flag
		var ipAddresses []net.IP
		var err error

		if *ipv6Flag {
			ipAddresses, err = net.LookupIP(hostname)
		} else {
			ipAddresses, err = net.LookupIP(hostname)
		}

		if err != nil {
			// Handle errors and consider security implications
			fmt.Printf("Error resolving %s: %v\n", hostname, err)
			continue
		}

		// Filter and print the IP addresses
		for _, ip := range ipAddresses {
			// Filter for IPv4 or IPv6 based on the flag
			if !*ipv6Flag && ip.To4() == nil {
				continue // Skip IPv6 if not requested
			}
			fmt.Println(ip)
		}
	}

	if err := scanner.Err(); err != nil {
		// Handle any errors related to reading input
		fmt.Printf("Error reading input: %v\n", err)
	}
}
