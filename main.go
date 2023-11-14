package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func worker(id int, jobs <-chan string, ipWG *sync.WaitGroup) {
	defer ipWG.Done()
	for input := range jobs {
		// Remove prefixes like "http://", "https://", and "www."
		hostname := removePrefixes(input)
		ips, err := net.LookupIP(hostname)
		if err != nil {
			fmt.Printf("Error looking up IP for %s: %v\n", hostname, err)
			continue
		}
		for _, ip := range ips {
			if hostnameFlag {
				fmt.Println(hostname)
			}
			// only display IPv4 addresses
			if ip.To4() != nil {
				fmt.Println(ip)
			}
		}
	}
}

func removePrefixes(input string) string {
	// Remove "http://", "https://", and "www." prefixes
	prefixes := []string{"http://", "https://", "www."}
	for _, prefix := range prefixes {
		if strings.HasPrefix(input, prefix) {
			input = strings.TrimPrefix(input, prefix)
		}
	}
	return input
}

var hostnameFlag bool

func main() {
	flag.BoolVar(&hostnameFlag, "H", false, "Display hostname")
	concurrency := flag.Int("c", 20, "Number of concurrent workers")
	flag.Parse()

	jobs := make(chan string)
	var ipWG sync.WaitGroup

	for w := 1; w <= *concurrency; w++ {
		ipWG.Add(1)
		go worker(w, jobs, &ipWG)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)

	ipWG.Wait()
}
