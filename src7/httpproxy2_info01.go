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
    for {
        time.Sleep(10 * time.Second)
        fmt.Println( 
            "http" , 
            atomic . LoadUint64(&_vAccessHttpF ),
            atomic . LoadUint64(&_vAccessHttpS ),
            "          ",
            "http" , 

            atomic . LoadUint64(&_vAccessSslF ),
            atomic . LoadUint64(&_vAccessSslS ),
            "\n",
        )
    }
    _vWait.Done()
}

