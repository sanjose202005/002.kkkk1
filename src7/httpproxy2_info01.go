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
    var __VipListErrmapCnt int = 0 ;
    var __VipListOKCnt int = 0 ;
    var __vPrintErrmap bool = false

    for {
        time.Sleep(10 * time.Second)
        __vPrintErrmap = false
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

        if __cnt % 10 == 9 {
            __vPrintErrmap = true
        }
        if _VipListErrmap.Count() != __VipListErrmapCnt {
            __vPrintErrmap = true
            __VipListErrmapCnt = _VipListErrmap.Count() 
        }
        if _VipListOKmap.Count() != __VipListOKCnt {
            __vPrintErrmap = true
            __VipListOKCnt = _VipListOKmap.Count() 
        }

        __vPrintErrmap = true
        if __vPrintErrmap && __VipListErrmapCnt != 0 {
            fmt.Println( 
                "_VipListErrmap : " ,
                _VipListErrmap . Items() ,
                "\n _VipListOKmap : " ,
                _VipListOKmap . Items() ,
            )
        }
    }
    _vWait.Done()
}

