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
    var __VipListOkmapCnt int = 0 ;
    var __vPrintMap bool = false

    for {
        time.Sleep(10 * time.Second)
        __vPrintMap = false

        if __cnt % 10 == 9 {
            __vPrintMap = true
        }
        if _VipListErrmap.Count() != __VipListErrmapCnt {
            __vPrintMap = true
            __VipListErrmapCnt = _VipListErrmap.Count() 
        }
        if _VipListOKmap.Count() != __VipListOkmapCnt {
            __vPrintMap = true
            __VipListOkmapCnt = _VipListOKmap.Count() 
        }

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
            ", okIP:", __VipListOkmapCnt,
            ", erIP:", __VipListErrmapCnt,
            "\n",
        )
        __cnt += 1

        //__vPrintMap = true
        if __vPrintMap {
            if __VipListErrmapCnt != 0 {
                fmt.Println( 
                    "_VipListErrmap : " ,
                    _VipListErrmap . Items() ,
                )
            }
            if __VipListOkmapCnt != 0 {
                fmt.Println( 
                    "_VipListOKmap : " ,
                    _VipListOKmap . Items() ,
                )
            }
        }
    }
    _vWait.Done()
}

