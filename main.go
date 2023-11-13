package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	ipv6Flag := flag.Bool("v6", false, "Resolve IPv6 addresses")
	concurrencyFlag := flag.Int("c", 20, "Set the concurrency level")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	// Create a worker pool with the specified concurrency
	workerPool := make(chan struct{}, *concurrencyFlag)
	var wg sync.WaitGroup

	for i := 0; i < *concurrencyFlag; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				// Read a hostname from the scanner
				hostname, err := readHostname(scanner)
				if err != nil {
					// End of input, no more hostnames to process
					return
				}

				// Limit concurrency by adding to the worker pool
				workerPool <- struct{}{}

				go func(hostname string) {
					defer func() {
						<-workerPool // Release a worker when done
					}()

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
						return
					}

					// Filter and print the IP addresses
					for _, ip := range ipAddresses {
						// Filter for IPv4 or IPv6 based on the flag
						if !*ipv6Flag && ip.To4() == nil {
							continue // Skip IPv6 if not requested
						}
						fmt.Println(ip)
					}
				}(hostname)
			}
		}()
	}

	wg.Wait()
}

// readHostname reads a hostname from the scanner and handles errors.
func readHostname(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	return "", fmt.Errorf("End of input")
}
