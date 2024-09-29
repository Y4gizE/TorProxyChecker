# Tor Proxy Connection Checker

This Go application checks the connection through the Tor network using a SOCKS5 proxy. It sends an HTTP request to `http://check.torproject.org` and returns the HTTP status code to confirm if the connection is routed through Tor.

## Requirements

- [Tor](https://www.torproject.org/download/) must be installed and running on your machine.
- You must configure the Tor `torcc` file to enable the SOCKS5 proxy.

## Usage

1. **Tor Configuration (torcc)**:
   
   Make sure your `torcc` file contains the following lines:

   ```bash
   SocksPort 0.0.0.0:9150
   DataDirectory B:\Tor Browser\Browser\TorBrowser\Data\Tor
