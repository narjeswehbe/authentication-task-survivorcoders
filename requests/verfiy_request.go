package requests

type VerifyEmailRequest struct {
	Email             string `json:"email"`
	Verification_code string `json:"code"`
	Password          string `json:"password"`
}
