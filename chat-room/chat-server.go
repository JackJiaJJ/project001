package main

import (
        "fmt"
        "net"
        "strings"
        "time"
)

func ErrorCheck(component string, err error) {
        if err != nil {
                fmt.Printf("%v: %v\n", component, err)
                return
        }
}

var connMap = make(map[string]net.Conn)

func main() {
        t := time.Now()
        fmt.Printf("begin to start chat server at %v\n", t)

        ln, err := net.Listen("tcp", "127.0.0.1:8085")
        ErrorCheck("net.Listen", err)
        fmt.Println("server is running ...")

        defer ln.Close()

        for {
                conn, err := ln.Accept()
                ErrorCheck("ln.Accept()", err)
                fmt.Println("server is running ..")

                go handle(conn)
        }
}

func handle(conn net.Conn) {
        for {
                data := make([]byte, 512)
                n, err := conn.Read(data)
                if n == 0 || err != nil {
                        continue
                }

                msg_str := strings.Split(string(data[0:n]), "|")

                switch msg_str[0] {
                case "nick":
                        for k, v := range connMap {
                                if k != msg_str[1] {
                                        v.Write([]byte("[" + msg_str[1] + "]: joining..."))
                                }
                        }
                        connMap[msg_str[1]] = conn
                        fmt.Println(connMap)
                case "say":
                        for k, v := range connMap {
                                if k != msg_str[1] {
                                        fmt.Println("send " + msg_str[2] + " to " + k)
                                        v.Write([]byte("[" + msg_str[1] + "]: " + msg_str[2]))
                                }
                        }
                case "quit", "exit":
                        for k, v := range connMap {
                                if k != msg_str[1] {
                                        v.Write([]byte("[" + msg_str[1] + "]: quit"))
                                }
                        }
                        delete(connMap, msg_str[1])
                }
        }
}
