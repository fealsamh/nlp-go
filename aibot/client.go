package aibot

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const serviceUrl = "https://ai-bot-srv.nw.r.appspot.com"

type Client struct {
	SecretKey string
}

func (c *Client) httpClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}

func (c *Client) httpError(r *http.Response) error {
	msg := "failed to read error message"
	if b, err := io.ReadAll(r.Body); err == nil {
		msg = strings.TrimSpace(string(b))
	} else {
		msg += " (" + err.Error() + ")"
	}
	return fmt.Errorf("%s: %s", r.Status, msg)
}
