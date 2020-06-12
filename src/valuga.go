package main

import (
	"io"
	"net"
	"net/http"
	"time"
    "fmt"

	"golang.org/x/net/proxy"
)

func handleHTTP(w http.ResponseWriter, req *http.Request, dialer proxy.Dialer) {
	tp := http.Transport{
		Dial: dialer.Dial,
	}
	resp, err := tp.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func handleTunnel(w http.ResponseWriter, req *http.Request, dialer proxy.Dialer) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	srcConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	dstConn, err := dialer.Dial("tcp", req.Host)
	if err != nil {
		srcConn.Close()
		return
	}

	srcConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	go transfer(dstConn, srcConn)
	go transfer(srcConn, dstConn)
}

func transfer(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()

	io.Copy(dst, src)
}

func serveHTTP(w http.ResponseWriter, req *http.Request) {
    //upProxyStr := "134.175.102.70:1080" ;
    upProxyStr := "216.144.228.130:15378"  ;
	d := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	//dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, d)
	//dialer, _ := proxy.SOCKS5("tcp", "114.99.2.150:38801", nil, d)
	//dialer, _ := proxy.SOCKS5("tcp", "134.175.102.70:1080" , nil, d)
    fmt.Printf( "\n\nUsing up-proxy on : %s\n\n" , upProxyStr );
	dialer, _ := proxy.SOCKS5("tcp", upProxyStr , nil, d)

	if req.Method == "CONNECT" {
		handleTunnel(w, req, dialer)
	} else {
		handleHTTP(w, req, dialer)
	}
}

func main() {
    listenStr := "127.0.0.1:22226" ;
    fmt.Printf( "\n\nListing on : %s\n\n" , listenStr );
	http.ListenAndServe( listenStr , http.HandlerFunc(serveHTTP))
}
