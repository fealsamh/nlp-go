package aibot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// DefaultServiceURL ...
const DefaultServiceURL = "https://ai-bot-srv.nw.r.appspot.com"

// Client ...
type Client struct {
	ServiceURL string
	SecretKey  string
}

func (c *Client) callServiceGet(path string, out any) (int, error) {
	return c.callService(http.MethodGet, path, nil, out)
}

func (c *Client) callServicePost(path string, in, out any) (int, error) {
	return c.callService(http.MethodPost, path, in, out)
}

func (c *Client) callService(method, path string, in, out any) (int, error) {
	var rd io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return 0, err
		}
		rd = bytes.NewReader(b)
	}
	cl := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, c.ServiceURL+path, rd)
	if err != nil {
		return 0, err
	}
	if c.SecretKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.SecretKey)
	}
	r, err := cl.Do(req)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return r.StatusCode, c.httpError(r)
	}
	if out != nil {
		if err := json.NewDecoder(r.Body).Decode(out); err != nil {
			return r.StatusCode, err
		}
	}
	return r.StatusCode, nil
}

func (c *Client) httpError(r *http.Response) error {
	msg := "failed to read error message"
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		msg = strings.TrimSpace(string(b))
	} else {
		msg += " (" + err.Error() + ")"
	}
	return fmt.Errorf("%s: %s", r.Status, msg)
}
