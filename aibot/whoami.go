package aibot

// User ...
type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Fullname string   `json:"fullname"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	OrgID    int      `json:"org_id"`
	OrgName  string   `json:"org_name"`
}

// Whoami ...
func (c *Client) Whoami() (*User, error) {
	var u *User
	if _, err := c.callServiceGet("/api/v1/whoami", &u); err != nil {
		return nil, err
	}
	return u, nil
}
