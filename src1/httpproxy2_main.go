// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    "flag"
    "io"
    "log"
    "net"
    "net/http"
    "strings"
)

type _TS_proxy struct {
    _TS_cfg _TS_cfg
}

// 设置type
type _TS_cfg struct {
    _Addr        string   // 监听地址
    _Port        string   // 监听端口
    _IsAnonymous bool     // 高匿名模式
    _Debug       bool     // 调试模式
}

func main() {

    // 参数
    faddr := flag.String("addr","0.0.0.0","监听地址，默认0.0.0.0")
    fprot := flag.String("port","22220","监听端口，默认22220")
    fanonymous :=  flag.Bool("anonymous",true,"高匿名，默认高匿名")
    fdebug :=  flag.Bool("debug",false,"调试模式显示更多信息，默认关闭")
    flag.Parse()

    cfg := &_TS_cfg{}
    cfg._Addr = *faddr
    cfg._Port = *fprot
    cfg._IsAnonymous = *fanonymous
    cfg._Debug = *fdebug
    // fmt.Println(cfg)
    Run(cfg)
}

func Run(cfg *_TS_cfg) {
    pxy := NewPxy()
    pxy.SetPxyCfg(cfg)
    log.Printf("HttpPxoy is runing on %s:%s \n", cfg._Addr, cfg._Port)
    // http.Handle("/", pxy)
    bindAddr := cfg._Addr + ":" + cfg._Port
    log.Fatalln(http.ListenAndServe(bindAddr, pxy))
}


// 实例化
func NewPxy() *_TS_proxy {
    return &_TS_proxy{
        _TS_cfg: _TS_cfg{
            _Addr:        "",
            _Port:        "22221",
            _IsAnonymous: true,
            _Debug:       false,
        },
    }
}

// 配置参数
func (___p1 *_TS_proxy) SetPxyCfg(cfg *_TS_cfg) {
    if cfg._Addr != "" {
        ___p1._TS_cfg._Addr = cfg._Addr
    }
    if cfg._Port != "" {
        ___p1._TS_cfg._Port = cfg._Port
    }
    if cfg._IsAnonymous != ___p1._TS_cfg._IsAnonymous {
        ___p1._TS_cfg._IsAnonymous = cfg._IsAnonymous
    }
    if cfg._Debug != ___p1._TS_cfg._Debug {
        ___p1._TS_cfg._Debug = cfg._Debug
    }

}

// 运行代理服务
func (___p2 *_TS_proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    // debug
    if ___p2._TS_cfg._Debug {
        log.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
        // fmt.Println(req)
    }

    // http && https
    if req.Method != "CONNECT" {
        // 处理http
        ___p2.http_deal_with(rw, req)
    } else {
        // 处理https
        // 直通模式不做任何中间处理
        ___p2.httpS_deal_with(rw, req)
    }

}

// http
func (___p3 *_TS_proxy) http_deal_with(rw http.ResponseWriter, req *http.Request) {

    transport := http.DefaultTransport

    // 新建一个请求outReq
    outReq := new(http.Request)
    // 复制客户端请求到outReq上
    *outReq = *req // 复制请求

    //  处理匿名代理
    if ___p3._TS_cfg._IsAnonymous == false {
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


// https
func (___p4 *_TS_proxy) httpS_deal_with(rw http.ResponseWriter, req *http.Request) {

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

