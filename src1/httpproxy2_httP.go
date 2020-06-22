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
func (___p3 *_TS_proxy) _http_deal_with(___rw3 http.ResponseWriter, ___req3 *http.Request) {

    __transport := http.DefaultTransport

    // 新建一个请求outReq
    __outReq := new(http.Request)
    // 复制客户端请求到outReq上
    *__outReq = *___req3 // 复制请求

    //  处理匿名代理
    if ___p3._vTS_cfg._vIsAnonymous == false {
        if clientIP, _, err := net.SplitHostPort(___req3.RemoteAddr); err == nil {
            if prior, ok := __outReq.Header["X-Forwarded-For"]; ok {
                clientIP = strings.Join(prior, ", ") + ", " + clientIP
            }
            __outReq.Header.Set("X-Forwarded-For", clientIP)
        }
    }

    // outReq请求放到传送上
    __res, err := __transport.RoundTrip(__outReq)
    if err != nil {
        ___rw3.WriteHeader(http.StatusBadGateway)
        ___rw3.Write([]byte(err.Error()))
        return
    }

    // 回写http头
    for key, value := range __res.Header {
        for _, v := range value {
            ___rw3.Header().Add(key, v)
        }
    }
    // 回写状态码
    ___rw3.WriteHeader(__res.StatusCode)
    // 回写body
    io.Copy(___rw3, __res.Body)
    __res.Body.Close()
}


