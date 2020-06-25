// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    //"flag"
    //    "io"
    "log"
    "net"
    "net/http"
    //    "strings"
    "fmt"
    //"strconv"
    //"os"
    //"sync"
    "sync/atomic"
    "time"
    //"github.com/streamrail/concurrent-map"
    "github.com/orcaman/concurrent-map"
    // https://pkg.go.dev/github.com/orcaman/concurrent-map?tab=doc
)

var _VipListOKmap  = cmap.New()
var _VipListErrmap = cmap.New()

func _Run(___proxy *_TS_proxy) {

    time.Sleep(1 * time.Second)
    fmt.Println(" wait until return 21 ", ___proxy )

    log.Fatalln(http.ListenAndServe(___proxy._vListen ,  ___proxy ))

    fmt.Println(" wait until return 25 ", ___proxy )
    _vWait.Done()
}

var _myAgent = map[string]bool{
    "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0": true, // for my linux firefox
    "Mozilla/5.0 (Android 10; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0": true, // for my pixel # 10
    "Mozilla/5.0 (Android 5.1.1; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0": true, // for test01 in China
}

// 运行代理服务 xxx
func (___p3 *_TS_proxy) ServeHTTP(___rw3 http.ResponseWriter, ___req3 *http.Request) {
    // debug
    if ___p3._vTS_cfg._vDebug {
        log.Printf("Received request %s %s %s\n", ___req3.Method, ___req3.Host, ___req3.RemoteAddr)
        // fmt.Println(___req3)
    }

    __vAgent := ___req3.UserAgent();

    if 2 == 3 { // show  all agent
        log.Printf(
            "\n ok, 810182382 , len %d : agent [%s]\n\n",
            len(__vAgent),
            __vAgent,
        )
    }

    __vRemoteIP , _, __err0:= net.SplitHostPort( ___req3 . RemoteAddr ) 
    if nil != __err0 {
        fmt.Println(" 8390184848 err met ", __err0 )
        return
    }
    fmt.Println(" 8438181 get __vRemoteIP ", __vRemoteIP )

    if _ , __ok := _myAgent[__vAgent] ; !__ok {
        log.Printf(
            "\n different 810182381 found : Method %s , Host %s , RemoteAddr %s \nURL %s\n Header: %s\n\n",
            ___req3.Method, ___req3.Host, ___req3.RemoteAddr,
            ___req3.URL,
            ___req3.Header,
        )
        atomic . AddUint64(&_vAccessAgentE , 1)

        if __agentIPcnt2 , __ok2 := _VipListErrmap.Get( __vRemoteIP ); __ok2 {
            _VipListErrmap . Set( __vRemoteIP , __agentIPcnt2.(int) + 1 )
        } else {
            _VipListErrmap . Set( __vRemoteIP , 1 )
        }

        return;
    }

    if __agentIPcnt3 , __ok3 := _VipListOKmap.Get( __vRemoteIP ); __ok3 {
        _VipListOKmap . Set( __vRemoteIP , __agentIPcnt3.(int) + 1 )
    } else {
        _VipListOKmap . Set( __vRemoteIP , 1 )
    }

    // http && https
    if ___req3.Method != "CONNECT" { // http
        if ___p3._vTS_cfg._vHttps == true {
            atomic . AddUint64(&_vAccessHttpF , 1)
        } else {
            atomic . AddUint64(&(_vAccessHttpS) , 1)
            // 处理http
            ___p3._http_deal_with(___rw3, ___req3)
        }
    } else {
        if ___p3._vTS_cfg._vHttps == false{
            atomic . AddUint64(&_vAccessSslF , 1)
        } else {
            atomic . AddUint64(&_vAccessSslS , 1)
            // 处理https
            // 直通模式不做任何中间处理
            ___p3._httpS_deal_with(___rw3, ___req3)
        }
    }


}

