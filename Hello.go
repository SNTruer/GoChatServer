package main

import (
    "bufio"
    "net"
    "fmt"
    //"os"
    "./Queue"
)

var messageQueue *Queue.Queue
const MAX_QUEUE_SIZE = 100
var users []*User

type User struct {
    //getText chan string
    reader *bufio.Reader
    writer *bufio.Writer
}

func makeServer(){

    messageQueue = Queue.NewQueue(MAX_QUEUE_SIZE)

    users = make([]*User, 0)
}

func acceptUsers(){

    in, _ := net.Listen("tcp", ":8081")

    for{
        fmt.Println("Logging")
        conn, error := in.Accept()
        fmt.Println(conn)

        //reader := bufio.NewReader(os.Stdin)

        if error == nil{
            user := User{
                reader : bufio.NewReader(conn),
                writer : bufio.NewWriter(conn),
            }
            users = append(users, &user)
            go getMessage(&user)
        }else {
            fmt.Println(error)
        }
    }

    //text, _ := reader.ReadString('\n')
}

func getMessage(user *User){
    fmt.Println("Get message start")
    fmt.Println(user)
    for{
        message, _ := user.reader.ReadString('\n')
        fmt.Println("Message Get")
        messageQueue.Push(&Queue.Node{message})
    }
    fmt.Println("Get message End")
}

func showMessage(){
    for {
        node := messageQueue.Pop()
        if node != nil {
            //fmt.Println(node.Value)
            broadcast(node.Value)
        }
    }
}

func broadcast(value string){
    for _, user := range users {
        user.writer.WriteString(value)
        user.writer.Flush()
    }
}

func main() {
    makeServer()

    go acceptUsers()

    go showMessage()

    for {
    }
}
