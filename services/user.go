package services

import (
	"auth_microservice/config"
	"auth_microservice/entity"
	"auth_microservice/myMiddleware"
	"auth_microservice/requests"
	"auth_microservice/response"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func SignUp(request requests.SighUpRequest) response.BaseResponse {
	code := GenerateCode()
	// validation
	if validEmail(request.Email) == false {
		return response.BaseResponse{Code: 400, Message: "Please enter a valid email"}
	}
	if len(request.Username) == 0 {
		return response.BaseResponse{Code: 400, Message: "username cannot be empty"}
	}
	if len(request.Firstname) == 0 {
		return response.BaseResponse{Code: 400, Message: "First name cannot be empty"}
	}
	if len(request.Lastname) == 0 {
		return response.BaseResponse{Code: 400, Message: "Last name cannot be empty"}
	}
	if len(request.Email) == 0 {
		return response.BaseResponse{Code: 400, Message: "Email cannot be empty"}
	}

	user := entity.User{Firstname: request.Firstname, Lastname: request.Lastname, Username: request.Username, Email: request.Email, Verification_code: strconv.FormatInt(int64(code), 10),
		Phone_number: request.Phone_number, Nationality: request.Nationality}

	var foundUser entity.User
	err := config.Db.Where("email = ?", request.Email).First(&foundUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		config.Db.Create(&user)
		ok2 := Send(code, request.Email)
		if ok2 == false {
			return response.BaseResponse{Code: 500, Message: "we could not send an email , please re-check it"}

		}
		return response.BaseResponse{Code: 200, Message: "Verification code sent"}
	} else { //the user is in the database!!
		// if the user requested to verify his email but he is already in the database , techincal problems for example
		if foundUser.Verified == true { //No need 2a3mello she , za3abne
			return response.BaseResponse{Code: 400, Message: "your account  already exits"}
		} else {
			ok := Send(code, request.Email)
			if ok == false {
				return response.BaseResponse{Code: 500, Message: "we could not send an email , please re-check it"}

			}
			foundUser.Verification_code = strconv.FormatInt(int64(code), 10)
			config.Db.Save(&foundUser)
			return response.BaseResponse{Code: 200, Message: "a new verification code is  sent"}
		}
	}
}

func VerifyEmail(request requests.VerifyEmailRequest) response.BaseResponse {

	var user entity.User
	if !validPassword(request.Password) {
		return response.BaseResponse{Code: 500, Message: "password should include 1 uppercase , 1 lowercase  at least 8 characters long"}
	}
	err := config.Db.Where("email = ?", request.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.BaseResponse{Code: 500, Message: "Email not found !! If you don't have an account , please create one"}
	}
	config.Db.Where("email = ?", request.Email).First(&user)
	if user.ID == 0 {
		return response.BaseResponse{Code: 500, Message: "Account not found"}
	}
	if user.Verified == true {
		return response.BaseResponse{Code: 200, Message: "Account already verified"}
	}
	if user.Verification_code == request.Verification_code {
		user.Verified = true

	} else {
		return response.BaseResponse{Code: 500, Message: "wrong verification code!!"}
	}

	user.Password = HashPassword(request.Password)
	config.Db.Save(&user)
	//fmt.Println(user)
	return response.BaseResponse{Code: 500, Message: "Your account is verified successfully"}

}

func Login(request requests.LoginRequest) response.BaseResponse {
	var tokenstring string
	var user *entity.User
	//validate request
	if len(request.Email) == 0 {
		return response.BaseResponse{500, "Email cannot be empty"}
	}
	if !validEmail(request.Email) {
		return response.BaseResponse{500, "Enter a Valid email"}
	}
	if len(request.Password) == 0 {
		return response.BaseResponse{500, "Password cannot be empty"}
	}
	//find the user
	config.Db.Where("email = ? OR username= ?", request.Email, request.Email).Find(&user)

	if user.ID == 0 {
		return response.BaseResponse{500, "User not found"}
	}
	if VerifyPassword(request.Password, user.Password) == false {
		return response.BaseResponse{500, "your credentials don't match "}

	} else {
		//generate token :
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":   user.ID,
			"auth_uuid": uuid.NewV4().String(),
			"exp":       time.Now().Add(time.Minute * 60).Unix(),
		})
		fmt.Println(token)
		var err error
		tokenstring, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			return response.BaseResponse{500, "failed to create token"}
		}
	}
	return response.BaseResponse{200, tokenstring}
}

func Logout(token string) response.BaseResponse {
	//get the token from header , it is passed in the param

	//get the claims
	claims, ok := myMiddleware.ExtractClaims(token)
	var authToken entity.BlackList
	auth_uuid := fmt.Sprint(claims["auth_uuid"])
	user_id := fmt.Sprint(claims["user_id"])
	final, _ := strconv.ParseInt(user_id, 10, 64)
	authToken.AuthUUID = auth_uuid
	authToken.UserID = uint64(final)
	authToken.ExpiryDate = time.Unix(int64(claims["exp"].(float64)), 0)

	config.Db.Save(&authToken)
	if !ok {
		return response.BaseResponse{500, "failed to log out"}
	}
	return response.BaseResponse{200, "logged out"}

}
