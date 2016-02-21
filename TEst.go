package main

import (
    "bufio"
    "fmt"
    "os"
)

func Test(){
    for{
        fmt.Println("12")
    }
}

func main() {

    reader := bufio.NewReader(os.Stdin)

    for{
        text, _ := reader.ReadString('\n')

        fmt.Println(text)
    }
}