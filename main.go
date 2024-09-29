package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// SOCKS5 proxy URL to connect through Tor
	torProxy := "socks5://127.0.0.1:9150"

	// Configure proxy settings
	proxyURL, err := url.Parse(torProxy)
	if err != nil {
		fmt.Println("Error creating proxy URL:", err)
		return
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // 30 seconds timeout
	}

	// Send HTTP request through the Tor network
	resp, err := client.Get("http://check.torproject.org")
	if err != nil {
		fmt.Println("Error during request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)
}
