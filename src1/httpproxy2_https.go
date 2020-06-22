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
func (___p4 *_TS_proxy) _httpS_deal_with(___rw4 http.ResponseWriter, ___req4 *http.Request) {

    // 拿出host
    __host := ___req4.URL.Host
    __hij, __ok := ___rw4.(http.Hijacker)
    if !__ok {
        log.Printf("HTTP Server does not support hijacking")
    }

    __client, _, __err41 := __hij.Hijack()
    if __err41 != nil {
        return
    }

    // 连接远程
    __server, __err42 := net.Dial("tcp", __host)
    if __err42 != nil {
        return
    }
    __client.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

    // 直通双向复制
    go io.Copy(__server, __client)
    go io.Copy(__client, __server)
}

