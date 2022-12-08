package requests

type TokenSaveRequest struct {
	Expires_at int64
	UserId     uint
	Token      string
}
