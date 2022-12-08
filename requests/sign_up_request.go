package requests

type SighUpRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	//optional fields
	Phone_number string `json:"number"`
	Nationality  string `json:"nationality"`
}
