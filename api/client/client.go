package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, token string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{},
	}
}


func (client *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, client.requestPath(path), &body)
	if err != nil {
		return nil, err
	}


	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", client.authToken))
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}
	f, err := os.Create("a.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	f.WriteString(req.Header.Get("Content-Type"))
	f.Close()
	
	f, err = os.Create("b.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	f.WriteString(req.Header.Get("Authorization"))
	f.Close()
	
	f, err = os.Create("c.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	f.WriteString(req.URL.String())
	f.Close()
	
	f, err = os.Create("d.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	f.WriteString(client.requestPath(path))
	f.Close()

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}