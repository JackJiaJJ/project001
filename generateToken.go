package rest

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "strings"
)

func GenerateToken(request Request) {
        client := &http.Client{}
        var (
                reqBody string
                resBody []byte
        )

        for id, key := range request.AccountID {
                reqBody = "grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=" + key
                req, err := http.NewRequest(string(request.Method), request.BaseURL, strings.NewReader(reqBody))
                //res,err := client.Post(request.BaseURL, request.BodyType, strings.NewReader(reqBody))

                ErrorCheck("make request", err)

                if len(request.Headers) != 0 {
                        for key, val := range request.Headers {
                                req.Header.Set(key, val)
                        }
                }

                res, err := client.Do(req)
                ErrorCheck("post request", err)
                defer res.Body.Close()
                if res.StatusCode != 200 {
                        fmt.Printf("failed to post request, got res.StatusCode: %d, error message: %s\n", res.StatusCode, http.StatusText(res.StatusCode))
                        panic("please check why post failed")
                }

                resBody, err = ioutil.ReadAll(res.Body)
                ErrorCheck("read response body", err)

                token := strings.Split(string(resBody), "\"")[3]

                request.Token[id] = token
        }
}
