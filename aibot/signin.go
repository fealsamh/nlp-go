package aibot

import (
	"bytes"
	"encoding/json"
)

type signinRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signinResponse struct {
	SecretKey string `json:"secret_key"`
	OrgName   string `json:"org_name"`
}

func (c *Client) Signin(username, password string) (string, string, error) {
	b, err := json.Marshal(&signinRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", "", err
	}
	r, err := c.httpClient().Post(serviceUrl+"/api/v1/signin", "application/json", bytes.NewReader(b))
	if err != nil {
		return "", "", err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return "", "", c.httpError(r)
	}
	var out signinResponse
	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return "", "", err
	}
	return out.SecretKey, out.OrgName, nil
}
