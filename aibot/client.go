package aibot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const serviceUrl = "https://ai-bot-srv.nw.r.appspot.com"

type Client struct {
	SecretKey string
}

func (c *Client) callServiceGet(path string, out interface{}) (int, error) {
	cl := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, serviceUrl+path, nil)
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

func (c *Client) callServicePost(path string, in, out interface{}) (int, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	cl := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodPost, serviceUrl+path, bytes.NewReader(b))
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
