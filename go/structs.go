package forum

type UserData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
}

type LoginData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
