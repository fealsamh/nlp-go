package aibot

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
	var u *User
	if err := c.callService("/api/v1/whoami", &whoamiRequest{
		SecretKey: c.SecretKey,
	}, &u); err != nil {
		return nil, err
	}
	return u, nil
}
