package services

import (
	"auth_microservice/config"
	"auth_microservice/entity"
	"auth_microservice/myMiddleware"
	"auth_microservice/requests"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func SignUp(request requests.SighUpRequest) {
	code := GenerateCode()
	user := entity.User{Firstname: request.Firstname, Lastname: request.Lastname, Username: request.Username, Email: request.Email, Verification_code: strconv.FormatInt(int64(code), 10),
		Phone_number: request.Phone_number, Nationality: request.Nationality}

	// na2es validation
	var foundUser entity.User
	err := config.Db.Where("email = ?", request.Email).First(&foundUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		config.Db.Create(&user)
		Send(code, request.Email)
		return
	} else { //the user is in the database!!
		// if the user requested to verify his email but he is already in the database , techincal problems for example
		if foundUser.Verified == true { //No need 2a3mello she , za3abne
			return
		} else {
			Send(code, request.Email)
			foundUser.Verification_code = strconv.FormatInt(int64(code), 10)
			config.Db.Save(&foundUser)
			return
		}
	}
}

func VerifyEmail(request requests.VerifyEmailRequest) string {

	var user entity.User
	err := config.Db.Where("email = ?", request.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "Email not found !! If you don't have an account , please create one"
	}
	config.Db.Where("email = ?", request.Email).First(&user)
	fmt.Println("found user")
	fmt.Println(user)
	if user.Verified == true {
		return "Account is already verified"
	}
	if user.Verification_code == request.Verification_code {
		user.Verified = true
	} else {
		return "wrong verification code"
	}
	//	p, _ := GenerateFromPassword(user.Password)
	//pass 8 didgitd
	user.Password = HashPassword(request.Password)
	config.Db.Save(&user)
	//fmt.Println(user)
	return "Account verified"

}

func Login(request requests.LoginRequest) string {
	var user *entity.User
	//validate request
	//find the user
	config.Db.Where("email = ? OR username= ?", request.Email, request.Email).Find(&user)
	fmt.Println(user)
	if user.ID == 0 {
		return "User not found"
	}
	if VerifyPassword(request.Password, user.Password) == true {
		fmt.Println("Logged in ")

	} else {
		return "your credentials do not match ! "
	}
	//generate token :
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"auth_uuid": uuid.NewV4().String(),
		"exp":       time.Now().Add(time.Minute * 60).Unix(),
	})
	fmt.Println(token)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "failed to create token"
	}
	return tokenString
}

func Logout(token string) bool {
	//get the token from header , it is passed in the param
	fmt.Println("inside services")
	//get the claims
	claims, ok := myMiddleware.ExtractClaims(token)
	var authToken entity.BlackList
	auth_uuid := fmt.Sprint(claims["auth_uuid"])
	user_id := fmt.Sprint(claims["user_id"])
	final, _ := strconv.ParseInt(user_id, 10, 64)
	authToken.AuthUUID = auth_uuid
	authToken.UserID = uint64(final)
	authToken.ExpiryDate = time.Unix(int64(claims["exp"].(float64)), 0)

	//save token to black list !!!

	config.Db.Save(&authToken)
	return ok
	//save to blacklist

}
