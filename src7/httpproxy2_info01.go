// https://gist.github.com/wwek/d48d49ccbb568d9411935d00bc6e0395/raw/4a7130fd49c16cfeddbf12fd7d2da57426077747/httpproxy.go
package main

import (
    //"flag"
    //    "io"
    //"log"
    //    "net"
    //"net/http"
    //    "strings"
    "fmt"
    //"strconv"
    //"os"
    //"sync"
    "sync/atomic"
    "time"
)


func _printInfo01(){
    var __cnt uint64 = 0
    for {
        time.Sleep(10 * time.Second)
        fmt.Println( 
            "errorAgent" , 
            atomic . LoadUint64(&_vAccessAgentE ),
            "          ",
            "httpF" , 
            atomic . LoadUint64(&_vAccessHttpF ),
            "httpS" , 
            atomic . LoadUint64(&_vAccessHttpS ),
            "          ",
            "SslF" , 
            atomic . LoadUint64(&_vAccessSslF ),
            "SslS" , 
            atomic . LoadUint64(&_vAccessSslS ),
            "          ",
            __cnt,
            "\n",
        )
        __cnt += 1
    }
    _vWait.Done()
}

