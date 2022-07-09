package aibot

type signupRequest struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"org_name"`
}

func (c *Client) Signup(username, fullname, email, password, orgName string) error {
	if err := c.callServicePost("/api/v1/signup", &signupRequest{
		Username: username,
		Fullname: fullname,
		Email:    email,
		Password: password,
		OrgName:  orgName,
	}, nil); err != nil {
		return err
	}
	return nil
}
