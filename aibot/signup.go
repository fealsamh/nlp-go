package aibot

import (
	"bytes"
	"encoding/json"
)

type signupRequest struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"org_name"`
}

func (c *Client) Signup(username, fullname, email, password, orgName string) error {
	b, err := json.Marshal(&signupRequest{
		Username: username,
		Fullname: fullname,
		Email:    email,
		Password: password,
		OrgName:  orgName,
	})
	if err != nil {
		return err
	}
	r, err := c.httpClient().Post(serviceUrl+"/api/v1/signup", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return c.httpError(r)
	}
	return nil
}
