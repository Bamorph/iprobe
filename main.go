package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		hostname := scanner.Text()

		// Resolve the hostname to an IP address
		ipAddresses, err := net.LookupHost(hostname)
		if err != nil {
			// Handle errors, and consider security implications
			fmt.Printf("Error resolving %s: %v\n", hostname, err)
			continue
		}

		// Print the IP addresses to the standard output
		for _, ip := range ipAddresses {
			fmt.Println(ip)
		}
	}

	if err := scanner.Err(); err != nil {
		// Handle any errors related to reading input
		fmt.Printf("Error reading input: %v\n", err)
	}
}
