package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストの詳細をログに出力
		log.Printf("Received request: %s %s from %s", r.Method, r.URL, r.RemoteAddr)
		_, err := w.Write([]byte("Hello, HTTP/3 world!"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	server := http3.Server{
		Addr:      "0.0.0.0:8080",
		Handler:   handler,
		TLSConfig: generateTLSConfig(),
	}

	log.Println("Starting HTTP/3 server on :8080")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func generateTLSConfig() *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{loadCertificate()},
		NextProtos:   []string{"h3"},
	}
}

func loadCertificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal("Failed to load certificate: ", err)
	}
	return cert
}
