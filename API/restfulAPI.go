package API

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "strings"
)

func AddQueryParams(baseURL string, queryParams map[string]string) string {
        baseURL += "?"
        params := url.Values{}

        for key, value := range queryParams {
                params.Add(key, value)
        }

        return baseURL + params.Encode()
}

func BuildRequest(request Request) (*http.Request, error) {
        if len(request.QueryParams) != 0 {
                request.BaseURL = AddQueryParams(request.BaseURL, request.QueryParams)
        }

        req, err := http.NewRequest(string(request.Method), request.BaseURL, strings.NewReader(request.Body))
        if err != nil {
                fmt.Printf("failed to build request %v, error is: %v\n", err)
        }

        if len(request.Headers) != 0 {
                for key, value := range request.Headers {
                        req.Header.Set(key, value)
                }
        }

        return req, err
}

func MakeRquest(req *http.Request) (*http.Response, error) {
        client := http.Client{}
        return client.Do(req)
}

func OutputResponse(res *http.Response) string {
        if res.StatusCode != 200 {
                fmt.Printf("failed to execute this invoke, response code is %d, error message is %s\n", res.StatusCode, http.StatusText(res.StatusCode))
        }
        data, err := ioutil.ReadAll(res.Body)
        if err != nil {
                fmt.Println("failed to get response, error:", err)
        }

        return string(data)
}
