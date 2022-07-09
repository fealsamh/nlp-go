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

func (c *Client) callService(path string, in, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	cl := &http.Client{Timeout: 5 * time.Second}
	r, err := cl.Post(serviceUrl+path, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return c.httpError(r)
	}
	if out != nil {
		if err := json.NewDecoder(r.Body).Decode(out); err != nil {
			return err
		}
	}
	return nil
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
