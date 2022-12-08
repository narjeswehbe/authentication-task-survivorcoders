package services

import (
	"auth_microservice/config"
	"auth_microservice/entity"
	"auth_microservice/requests"
	"auth_microservice/response"
)

func SaveToken(request requests.TokenSaveRequest) response.BaseResponse {
	//validate the save token request
	t := entity.Token{Token: request.Token, Expires_at: request.Expires_at, UserId: request.UserId}

	err := config.Db.Save(&t)
	if (err != nil) {
		return response.BaseResponse{Code: 500, Message: "unable to save token"}
	}
	return response.BaseResponse{Code: 500, Message: "token saved"}

}