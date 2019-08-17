package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func getNowTimeStamp() string {
	return time.Now().Format("2006-01-02T15:04:05.999999Z")
}

func main() {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			log.Printf("[%s] GetConn: %s\n", getNowTimeStamp(), hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("[%s] GotConn: %+v\n", getNowTimeStamp(), connInfo)
		},
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			log.Printf("[%s] DNSStart: %+v\n", getNowTimeStamp(), dnsInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			log.Printf("[%s] DNSDone: %+v\n", getNowTimeStamp(), dnsInfo)
		},
		GotFirstResponseByte: func() {
			log.Printf("[%s] GotFirstResponseByte\n", getNowTimeStamp())
		},
		TLSHandshakeStart: func() {
			log.Printf("[%s] TLSHandshakeStart\n", getNowTimeStamp())
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				log.Printf("[%s] error TLSHandshakeDone: %s\n", getNowTimeStamp(), err.Error())
			}
			log.Printf("[%s] TLSHandshakeDone: %+v\n", getNowTimeStamp(), state)
		},
		WroteHeaderField: func(key string, value []string) {
			log.Printf("[%s] WroteHeaderField: %s %+v\n", getNowTimeStamp(), key, value)
		},
		WroteRequest: func(reqInfo httptrace.WroteRequestInfo) {
			log.Printf("[%s] WroteRequest: %+v\n", getNowTimeStamp(), reqInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
