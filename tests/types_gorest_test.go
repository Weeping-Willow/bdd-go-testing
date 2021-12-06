package tests

type userListResponse struct {
	Users []user `json:"data"`
}

type userCreateResponse struct {
	User user `json:"data"`
}

type user struct {
	Id     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Status string `json:"status,omitempty"`
}

type getUserRequest struct {
	Email string `json:"email"`
}

func getDefaultUser() *user {
	return &user{
		Name:   "John Doe",
		Email:  "randomemailever@kmail.lv",
		Gender: "male",
		Status: "active",
	}
}

func (u *user) userWithId(id int) *user {
	u.Id = id
	return u
}

func (u *user) userWithEmail(email string) *user {
	u.Email = email
	return u
}
