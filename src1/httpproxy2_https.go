// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
//    "flag"
    "io"
    "log"
    "net"
    "net/http"
//    "strings"
)



// https
func (___p4 *_TS_proxy) _httpS_deal_with(rw http.ResponseWriter, req *http.Request) {

    // 拿出host
    host := req.URL.Host
    hij, ok := rw.(http.Hijacker)
    if !ok {
        log.Printf("HTTP Server does not support hijacking")
    }

    client, _, err := hij.Hijack()
    if err != nil {
        return
    }

    // 连接远程
    server, err := net.Dial("tcp", host)
    if err != nil {
        return
    }
    client.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

    // 直通双向复制
    go io.Copy(server, client)
    go io.Copy(client, server)
}

