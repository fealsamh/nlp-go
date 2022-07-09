package aibot

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
	if err := c.callServiceGet("/api/v1/whoami", &u); err != nil {
		return nil, err
	}
	return u, nil
}
