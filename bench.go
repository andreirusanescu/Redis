package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

func sendCommand(conn net.Conn, cmd string) (string, error) {
	_, err := conn.Write([]byte(cmd))
	if err != nil {
		return "", err
	}
	reply, err := bufio.NewReader(conn).ReadString('\n')
	return reply, err
}

func main() {
	addr := "localhost:6379" // change if needed
	connections := 35000     // number of concurrent clients
	requests := 1            // number of requests per client

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < connections; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Println("Dial error:", err)
				return
			}
			defer conn.Close()

			for j := 0; j < requests; j++ {
				// Example: send PING
				_, err := sendCommand(conn, "*1\r\n$4\r\nPING\r\n")
				if err != nil {
					fmt.Println("Send error:", err)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	totalRequests := connections * requests
	fmt.Printf("Completed %d requests in %s (%.2f req/sec)\n",
		totalRequests, elapsed, float64(totalRequests)/elapsed.Seconds())
}
