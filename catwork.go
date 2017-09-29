package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "syscall"
    "strings"
    "net/http"
    "golang.org/x/crypto/ssh/terminal"
    "flag"
)


func main (){
    if terminal.IsTerminal(syscall.Stdin) {
        // Execute: go run main.go
        fmt.Print("Type something then press the enter key: ")
        var comment string
        fmt.Scan(&comment)
        fmt.Printf("Result: %s\n", comment)
        return
    }

    // Execute: echo "foo" | go run main.go
    Text_Area, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }

    // Generated by curl-to-Go: https://mholt.github.io/curl-to-go
    text_body := "body="
    text_body += string("[code]"+string(Text_Area)+"[/code]")
    body := strings.NewReader(text_body)
    const DefaultRoom = ""
    var room string
    flag.StringVar(&room,"r",DefaultRoom,"room ID to use")
    flag.Parse()
    req, err := http.NewRequest("POST", os.ExpandEnv("https://api.chatwork.com/v2/rooms/"+ room +"/messages"), body)
    if err != nil {
        // handle err
    }
    const Api_key = ""
    var Key string
    flag.StringVar(&Key,"k",Api_key,"API Key to use")
    flag.Parse()
    req.Header.Set("X-Chatworktoken", os.ExpandEnv(Key))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
     // handle err
    }
    defer resp.Body.Close()

    //fmt.Printf("Result: %s\n", string(Text_Area))
}
