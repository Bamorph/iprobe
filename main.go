package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
    "sync"
)


func worker(id int, jobs <-chan string, ipWG *sync.WaitGroup) {
    defer ipWG.Done()
    for hostname := range jobs {
        ips, _ := net.LookupIP(hostname)
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

func main() {
    var hostnameFlag bool
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
