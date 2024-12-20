package main

/*import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {
        server := flag.String("s", "127.0.0.1:8085", "-s server:port")
        name := flag.String("u", "jack", "-u username")
        flag.Parse()

        if len(os.Args) != 5 {
                fmt.Printf("Usage: %v -s server:port -u username\n", os.Args[0])
                os.Exit(1)
        }

        conn, err := net.Dial("tcp", *server)
        if err != nil {
                fmt.Println("net.Dial: ", err)
                return
        }

        defer conn.Close()

        go handle(conn)

        fmt.Println("Welcome", *name)
        conn.Write([]byte("nick|" + *name))

        for {
                //var msg string
                //fmt.Scanln(&msg)

                inputReader := bufio.NewReader(os.Stdin)
                msg, _ := inputReader.ReadString('\n')
                msg = strings.TrimSpace(msg)
                if msg == "quit" || msg == "exit" {
                        conn.Write([]byte("quit|" + *name))
                  break
                }
                conn.Write([]byte("say|" + *name + "|" + msg))
        }
}

func handle(conn net.Conn) {
        for {
                data := make([]byte, 512)
                n, err := conn.Read(data)
                if n == 0 || err != nil {
                        break
                }
                fmt.Println(string(data[:]))
        }
}*/
