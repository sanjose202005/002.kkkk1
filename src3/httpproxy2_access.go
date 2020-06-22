// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    //"flag"
    //    "io"
    "log"
    //    "net"
    "net/http"
    //    "strings"
    "fmt"
    //"strconv"
    //"os"
    //"sync"
    //"sync/atomic"
    "time"
)

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
//        if ___p3._vTS_cfg._vHttps {
//            atomic . AddUint64(&___p3._vAccessF , 1)
//        } else {
//            atomic . AddUint64(&___p3._vAccessS , 1)
            // 处理http
            ___p3._http_deal_with(___rw3, ___req3)
//        }
    } else {
//        if ___p3._vTS_cfg._vHttps == false{
//            atomic . AddUint64(&___p3._vAccessF , 1)
//        } else {
//            atomic . AddUint64(&___p3._vAccessS , 1)
            // 处理https
            // 直通模式不做任何中间处理
            ___p3._httpS_deal_with(___rw3, ___req3)
//        }
    }

}

