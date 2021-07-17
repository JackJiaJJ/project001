package rest

type Method string

const (
        Get    Method = "GET"
        Post   Method = "POST"
        Put    Method = "PUT"
        Patch  Method = "PATCH"
        Delete Method = "DELETE"
)

type Request struct {
        Method    Method
        BaseURL   string
        Headers   map[string]string
        AccountID map[string]string
        Token     map[string]string
}

type Response struct {
        StatusCode int
        Body       string
        Headers    map[string]string
}
