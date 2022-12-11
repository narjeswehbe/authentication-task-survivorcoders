package services

import (
	"auth_microservice/config"
	"auth_microservice/entity"
	"auth_microservice/requests"
	"auth_microservice/response"
	"fmt"
	"github.com/procyon-projects/chrono"
	"golang.org/x/net/context"
	"time"
)

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
}

func SaveToken(request requests.TokenSaveRequest) response.BaseResponse {
	//validate the save token request
	t := entity.BlackList{UserID: request.UserId, AuthUUID: request.AuthUUID, ExpiryDate: request.ExpiryDate}

	err := config.Db.Save(&t)
	if err != nil {
		return response.BaseResponse{Code: 500, Message: "unable to save token"}
	}
	return response.BaseResponse{Code: 500, Message: "token saved"}

}

// this function will check for expired tokens in the blacklist every 30 min
func CleanBlackList() {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		var tokens []entity.BlackList

		config.Db.First(&tokens)
		for i := 0; i < len(tokens); i++ {
			if tokens[i].ExpiryDate.Before(time.Now()) {
				//remove the token
				config.Db.Delete(&tokens[i])
			}
			fmt.Print("Fixed Rate of 5 seconds")
		}
	}, 5*time.Second)

	if err == nil {
		fmt.Print("Task has been scheduled successfully.")
	}
}
