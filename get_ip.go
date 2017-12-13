package main

import (
    "net"
    "os"
    "fmt"
)

func main(){
    domain := os.Args[1]
    if domain == "" {
        fmt.Println("1 argument expected")
    }
    addr, err := net.LookupIP(domain)
    if err != nil {
        fmt.Println("An error occured")
    } else {
        fmt.Println("Discovered IP Adresses: ")
        for _, ip := range addr {
        fmt.Println(ip)
    }
    }
}
