// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    "flag"
    //    "io"
    "log"
    //    "net"
    "net/http"
    //    "strings"
)

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
func ListenAndServe(addr string, handler Handler) error
*/

type _TS_proxy struct {
    _vTS_cfg _TS_cfg
}

// 设置type
type _TS_cfg struct {
    _Addr        string   // 监听地址
    _Port        string   // 监听端口
    _IsAnonymous bool     // 高匿名模式
    _Debug       bool     // 调试模式
}

var _vHandle01 =  _TS_proxy{
    _vTS_cfg: _TS_cfg{
        _Addr:        "",
        _Port:        "22221",
        _IsAnonymous: true,
        _Debug:       false,
    },
};

func main() {

    // 参数
    __faddr := flag.String("addr","0.0.0.0","监听地址，默认0.0.0.0")
    __fprot := flag.String("port","22220","监听端口，默认22220")
    __fanonymous :=  flag.Bool("anonymous",true,"高匿名，默认高匿名")
    __fdebug :=  flag.Bool("debug",false,"调试模式显示更多信息，默认关闭")
    flag.Parse()

    __cfg := &_TS_cfg{}

    //_vHandle01 .
    
    __cfg._Addr = *__faddr
    __cfg._Port = *__fprot
    __cfg._IsAnonymous = *__fanonymous
    __cfg._Debug = *__fdebug
    // fmt.Println(__cfg)
    

    _Run(__cfg)
}

func _Run(___cfg1 *_TS_cfg) {
    __pxy := &_vHandle01 
    __pxy._SetPxyCfg(___cfg1)
    log.Printf("HttpPxoy is runing on %s:%s \n", ___cfg1._Addr, ___cfg1._Port)
    // http.Handle("/", __pxy)
    __bindAddr := ___cfg1._Addr + ":" + ___cfg1._Port
    log.Fatalln(http.ListenAndServe(__bindAddr, __pxy))
}


// 配置参数
func (___p1 *_TS_proxy) _SetPxyCfg(___cfg2 *_TS_cfg) {
    if ___cfg2._Addr != "" {
        ___p1._vTS_cfg._Addr = ___cfg2._Addr
    }
    if ___cfg2._Port != "" {
        ___p1._vTS_cfg._Port = ___cfg2._Port
    }
    if ___cfg2._IsAnonymous != ___p1._vTS_cfg._IsAnonymous {
        ___p1._vTS_cfg._IsAnonymous = ___cfg2._IsAnonymous
    }
    if ___cfg2._Debug != ___p1._vTS_cfg._Debug {
        ___p1._vTS_cfg._Debug = ___cfg2._Debug
    }

}

// 运行代理服务 xxx
func (___p3 *_TS_proxy) ServeHTTP(___rw3 http.ResponseWriter, ___req3 *http.Request) {
    // debug
    if ___p3._vTS_cfg._Debug {
        log.Printf("Received request %s %s %s\n", ___req3.Method, ___req3.Host, ___req3.RemoteAddr)
        // fmt.Println(___req3)
    }

    // http && https
    if ___req3.Method != "CONNECT" {
        // 处理http
        ___p3._http_deal_with(___rw3, ___req3)
    } else {
        // 处理https
        // 直通模式不做任何中间处理
        ___p3._httpS_deal_with(___rw3, ___req3)
    }

}

