package requests

import "time"

type TokenSaveRequest struct {
	AuthUUID   string    `json:"expires_at"`
	UserId     uint64    `json:"user_id"`
	ExpiryDate time.Time `json:"expiry_date"`
}
