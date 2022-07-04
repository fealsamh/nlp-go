package aibot

import (
	"bytes"
	"encoding/json"
)

type whoamiRequest struct {
	SecretKey string `json:"secret_key"`
}

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Fullname string   `json:"fullname"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	OrgID    int      `json:"org_id"`
	OrgName  string   `json:"org_name"`
}

func (c *Client) Whoami() (*User, error) {
	b, err := json.Marshal(&whoamiRequest{
		SecretKey: c.SecretKey,
	})
	if err != nil {
		return nil, err
	}
	r, err := c.httpClient().Post(serviceUrl+"/api/v1/whoami", "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return nil, c.httpError(r)
	}
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
