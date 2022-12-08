package requests

type VerifyEmailRequest struct {
	Email             string
	Verification_code string
	Password          string
}
