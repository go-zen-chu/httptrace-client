package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptrace"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			log.Printf("GetConn: %s\n", hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("GotConn: %+v\n", connInfo)
		},
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			log.Printf("DNSStart: %+v\n", dnsInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			log.Printf("DNSDone: %+v\n", dnsInfo)
		},
		GotFirstResponseByte: func() {
			log.Println("GotFirstResponseByte")
		},
		TLSHandshakeStart: func() {
			log.Println("TLSHandshakeStart")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				log.Printf("error TLSHandshakeDone: %s\n", err.Error())
			}
			log.Printf("TLSHandshakeDone: %+v\n", state)
		},
		WroteHeaderField: func(key string, value []string) {
			log.Printf("WroteHeaderField: %s %+v\n", key, value)
		},
		WroteRequest: func(reqInfo httptrace.WroteRequestInfo) {
			log.Printf("WroteRequest: %+v\n", reqInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
