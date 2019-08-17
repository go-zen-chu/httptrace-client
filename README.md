# httptrace-client

httptrace sample in golang

```
$ go run main.go

2019/08/17 18:33:23 [2019-08-17T18:33:23.986281Z] GetConn: google.com:443
2019/08/17 18:33:23 [2019-08-17T18:33:23.986669Z] DNSStart: {Host:google.com}
2019/08/17 18:33:23 [2019-08-17T18:33:23.989765Z] DNSDone: {Addrs:[{IP:216.58.196.238 Zone:} {IP:2404:6800:400a:806::200e Zone:}] Err:<nil> Coalesced:false}
2019/08/17 18:33:24 [2019-08-17T18:33:24.039766Z] TLSHandshakeStart
2019/08/17 18:33:24 [2019-08-17T18:33:24.360351Z] TLSHandshakeDone: {Version:771 HandshakeComplete:true DidResume:false CipherSuite:49195 NegotiatedProtocol:h2 NegotiatedProtocolIsMutual:true ServerName: PeerCertificates:[0xc000096b00 0xc000097080] VerifiedChains:[[0xc000096b00 0xc000097080 0xc0002b5b80]] SignedCertificateTimestamps:[[0 164 185 9 144 ...]}
2019/08/17 18:33:24 [2019-08-17T18:33:24.361127Z] GotConn: {Conn:0xc0000a6a80 Reused:false WasIdle:false IdleTime:0s}
2019/08/17 18:33:24 [2019-08-17T18:33:24.361289Z] WroteHeaderField: :authority [google.com]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361335Z] WroteHeaderField: :method [GET]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361344Z] WroteHeaderField: :path [/]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361349Z] WroteHeaderField: :scheme [https]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361356Z] WroteHeaderField: accept-encoding [gzip]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361364Z] WroteHeaderField: user-agent [Go-http-client/2.0]
2019/08/17 18:33:24 [2019-08-17T18:33:24.361423Z] WroteRequest: {Err:<nil>}
2019/08/17 18:33:24 [2019-08-17T18:33:24.460487Z] GotFirstResponseByte
```

You'll know that https request is performed as follows:

1. create connection
1. dns resolve, get ip
1. TLS handshake start
1. get connection
1. write http header
1. finish writing http header
1. recieve response

Here is a capture of tcp packets.

From client ip (172.20.10.2) to server ip (172.217.25.110), packet sent on TLSv1.2.

![](./tcp.png)

