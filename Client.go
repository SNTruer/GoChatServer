package main

import (
    "bufio"
    "net"
    "fmt"
    "os"
    //"./Queue"
)

var reader *bufio.Reader
var writer *bufio.Writer

func read() {
    for{
        message, _ := reader.ReadString('\n')
        fmt.Println(message)
    }
}

func write() {
    stringWriter := bufio.NewReader(os.Stdin)
    for{
        message, _ := stringWriter.ReadString('\n')
        writer.WriteString(message)
        writer.Flush()
        fmt.Println("Send!")
    }
}

func main() {

    fmt.Println("Logging..")

    join, error := net.Dial("tcp", "127.0.0.1:8081")

    if error == nil{
        fmt.Println("Success")
        fmt.Println(join)
        fmt.Println(error)

        reader = bufio.NewReader(join)
        writer = bufio.NewWriter(join)

        go read()
        go write()
    } else {
        fmt.Println("Fail")
        fmt.Println(join)
        fmt.Println(error)
    }

    for{

    }
}