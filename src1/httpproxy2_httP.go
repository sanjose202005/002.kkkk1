// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
//    "flag"
    "io"
//    "log"
    "net"
    "net/http"
    "strings"
)


// http
func (___p3 *_TS_proxy) _http_deal_with(rw http.ResponseWriter, req *http.Request) {

    transport := http.DefaultTransport

    // 新建一个请求outReq
    outReq := new(http.Request)
    // 复制客户端请求到outReq上
    *outReq = *req // 复制请求

    //  处理匿名代理
    if ___p3._vTS_cfg._vIsAnonymous == false {
        if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
            if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
                clientIP = strings.Join(prior, ", ") + ", " + clientIP
            }
            outReq.Header.Set("X-Forwarded-For", clientIP)
        }
    }

    // outReq请求放到传送上
    res, err := transport.RoundTrip(outReq)
    if err != nil {
        rw.WriteHeader(http.StatusBadGateway)
        rw.Write([]byte(err.Error()))
        return
    }

    // 回写http头
    for key, value := range res.Header {
        for _, v := range value {
            rw.Header().Add(key, v)
        }
    }
    // 回写状态码
    rw.WriteHeader(res.StatusCode)
    // 回写body
    io.Copy(rw, res.Body)
    res.Body.Close()
}


