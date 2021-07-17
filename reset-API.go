package rest

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "errors"
)

var component string

func ErrorCheck(component string, err error) {
        if err != nil {
                fmt.Printf("%s: %v", component, err)
        }
}

func BuildRequest(request Request) []byte {
        client := &http.Client{}
        req, err := http.NewRequest(string(request.Method), request.BaseURL, nil)
        /*if err != nil {
           return req, err
        }*/
        ErrorCheck("make request", err)
        //fmt.Println("req is: ",*req)

        //check headers
        if len(request.Headers) != 0 {
                for key, value := range request.Headers {
                        req.Header.Set(key, value)
                }
        }

        //fmt.Println(req.Header)

        res, err := client.Do(req)

        res.Header.Set("Content-Type","application/json")
        //fmt.Println(res.Header)
        ErrorCheck("client.Do(req)", err)
        if res.StatusCode != 200 {
           fmt.Printf("failed to run GET API, get response.StatusCode: %d, statusText: %s\n",res.StatusCode,http.StatusText(res.StatusCode))
           fmt.Printf("request url is %s\n",request.BaseURL)
           panic(errors.New("need to check why"))
        }

        defer res.Body.Close()
        //got response
        data, _ := ioutil.ReadAll(res.Body)
        //fmt.Printf("%s\n", data)
        return data
}
