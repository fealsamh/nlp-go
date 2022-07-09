package aibot

type signinRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signinResponse struct {
	SecretKey string `json:"secret_key"`
	OrgName   string `json:"org_name"`
}

func (c *Client) Signin(username, password string) (string, string, error) {
	var out *signinResponse
	if err := c.callService("/api/v1/signin", &signinRequest{
		Username: username,
		Password: password,
	}, &out); err != nil {
		return "", "", err
	}
	return out.SecretKey, out.OrgName, nil
}
