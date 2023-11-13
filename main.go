package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
    "sync"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        ips, _ := net.LookupIP(j)
        for _, ip := range ips {
            fmt.Println(ip)
        }
    }
}

func main() {
    concurrency := flag.Int("c", 20, "Number of concurrent workers")
    flag.Parse()

    jobs := make(chan string)
    var wg sync.WaitGroup

    for w := 1; w <= *concurrency; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        jobs <- scanner.Text()
    }
    close(jobs)

    wg.Wait()
}
