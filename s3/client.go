package s3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JackJiaJJ/project001/log"
)

type Options struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Endpoint        string
}
type Client struct {
	options Options
}

type Bucket struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

func New(opts *Options) (*Client, error) {
	if opts.AccessKeyID == "" || opts.SecretAccessKey == "" {
		return nil, fmt.Errorf("both accessKeyID and secretAccessKey must be provided")
	}

	if opts.Region == "" {
		log.Info("[New] [Region is not set, using default region - US South now]")
		opts.Region = "us-south-1"
	}

	if opts.Endpoint == "" {
		log.Info("[New] [Endpoint is not set, using default endpoint - http://127.0.0.1:9091/buckets/getAll now]")
		opts.Endpoint = "http://127.0.0.1:9091/buckets/getAll"
	}

	return &Client{options: *opts}, nil
}

func (c *Client) ExecuteHttpRequest(method string) ([]byte, error) {
	r, err := http.NewRequest(method, c.options.Endpoint, nil)
	if err != nil {
		log.Errorf("[ExecuteGetRequest] [Failed to execute Http Request, error is - %v]", err)
		return nil, err
	}

	hc := &http.Client{}
	resp, err := hc.Do(r)
	if err != nil {
		log.Errorf("[ExecuteHttpRequest] [Failed to do HTTP request, error is - %v]", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("[ExecuteHttpRequest] [Failed to read response body, error is - %v]", err)
		return nil, err
	}
	return body, nil
}

func (c *Client) GetBucket(ctx context.Context, bucketName string) (Bucket, error) {
	var buckets []Bucket
	var bucket Bucket

	select {
	case <-ctx.Done():
		return bucket, ctx.Err()

	default:
		log.Infof("[GetBucket] [Trying to get bucket - %v]", bucketName)
	}

	data, err := c.ExecuteHttpRequest(http.MethodGet)
	if err != nil {
		log.Errorf("[GetBucket] [Failed to execute HTTP request, error is - %v]", err)
		return bucket, err
	}
	err = json.Unmarshal(data, &buckets)
	if err != nil {
		log.Errorf("[GetBucket] [Failed to unmarshal JSON data, error is - %v]", err)
		return bucket, err
	}

	for _, b := range buckets {
		if b.Name == bucketName {
			return b, nil
		}
	}
	return bucket, fmt.Errorf("bucket %s not found", bucketName)
}
