// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    "flag"
    //    "io"
    "log"
    //    "net"
    "net/http"
    //    "strings"
    "fmt"
    "strconv"
    "os"
    "sync"
    "time"
)

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
func ListenAndServe(addr string, handler Handler) error
*/

type _TS_proxy struct {
    _vTS_cfg    _TS_cfg
    _vListen    string // exampel : 0.0.0.0:22220
    _vR01new    uint64
    _vR02deal   uint64
    _vW01new    uint64
    _vW02deal   uint64
}

// 设置type
type _TS_cfg struct {
    _vAddr        string   // 监听地址
    _vPort        string   // 监听端口
    _vIsAnonymous bool     // 高匿名模式
    _vDebug       bool     // 调试模式
    _vHttps       bool     // https mode -> true ; http mdoe -> false
}

var _vHandle01 =  _TS_proxy{
    _vTS_cfg: _TS_cfg{
        _vAddr:        "0.0.0.0",
        _vPort:        "22221",
        _vIsAnonymous: true,
        _vDebug:       false,
        _vHttps:       false,
    },
};

var _vHandle02 =  _TS_proxy{}

var _vWait      sync.WaitGroup

func main() {

    // 参数
    __faddr := flag.String("addr","0.0.0.0","监听地址，默认0.0.0.0")
    __fprot := flag.String("port","22220","监听端口，默认22220")
    __fanonymous :=  flag.Bool("anonymous",true,"高匿名，默认高匿名")
    __fdebug :=  flag.Bool("debug",false,"调试模式显示更多信息，默认关闭")
    flag.Parse()

    _vHandle01._vTS_cfg._vAddr = *__faddr
    _vHandle01._vTS_cfg._vPort = *__fprot
    _vHandle01._vTS_cfg._vIsAnonymous = *__fanonymous
    _vHandle01._vTS_cfg._vDebug = *__fdebug

    _vHandle02 = _vHandle01

    __port , __err := strconv.Atoi( _vHandle01._vTS_cfg._vPort )
    if __err != nil { fmt.Println(" error of port : " , _vHandle01._vTS_cfg._vPort ) ; os.Exit(3); }
    __port += 10
    _vHandle02._vTS_cfg._vPort = strconv.Itoa( __port )

    _vHandle01._vListen = _vHandle01._vTS_cfg._vAddr + ":" + _vHandle01._vTS_cfg._vPort
    _vHandle02._vListen = _vHandle02._vTS_cfg._vAddr + ":" + _vHandle02._vTS_cfg._vPort

    fmt.Println("_vHandle01", _vHandle01)
    fmt.Println("_vHandle02", _vHandle02)


    _vWait.Add(1) ; go _Run( &_vHandle01 )
    _vWait.Add(1) ; go _Run( &_vHandle02 )

    _vWait.Wait()
}

func _Run(___proxy *_TS_proxy) {

    time.Sleep(1 * time.Second)
    fmt.Println(" wait until return 21 ", ___proxy )

    log.Fatalln(http.ListenAndServe(___proxy._vListen ,  ___proxy ))

    fmt.Println(" wait until return 25 ", ___proxy )
    _vWait.Done()
}


// 运行代理服务 xxx
func (___p3 *_TS_proxy) ServeHTTP(___rw3 http.ResponseWriter, ___req3 *http.Request) {
    // debug
    if ___p3._vTS_cfg._vDebug {
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

