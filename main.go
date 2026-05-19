package main

import (
	"log"
	"os"
	"github.com/things-go/go-socks5"
)

func main() {
	port := os.Getenv("PROXY_PORT")
	if port == "" {
		port = "443"
	}

	username := os.Getenv("PROXY_USERNAME")
	password := os.Getenv("PROXY_PASSWORD")

	var server *socks5.Server
	if username != "" && password != "" {
		log.Printf("Starting SOCKS5 server on port %s with username/password authentication", port)
		creds := socks5.StaticCredentials{
			username: password,
		}
		server = socks5.NewServer(
			socks5.WithCredential(creds),
			socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		)
	} else {
		log.Printf("Starting SOCKS5 server on port %s without authentication (public access)", port)
		server = socks5.NewServer(
			socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		)
	}

	addr := ":" + port
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Fatalf("Failed to start SOCKS5 server: %v", err)
	}
}
