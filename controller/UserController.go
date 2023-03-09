package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rankwillServer/common"
	"rankwillServer/dto"
	"rankwillServer/model"
	"rankwillServer/response"
	"regexp"
)

func isEmailExisted(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email=?", email).First(&user)
	return user.ID != 0
}
func isFollowExisted(db *gorm.DB, uname string, lcusername string) bool {
	var fl model.Following
	db.Where("username=?", uname).Where("lcusername=?", lcusername).First(&fl)
	return fl.ID != 0
}
func getUserByEmail(db *gorm.DB, email string) model.User {
	var user model.User
	db.Where("email=?", email).First(&user)
	return user
}
func validEmail(email string) (bool, error) {
	regex := "^([a-z0-9A-Z]+[-|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$"
	return regexp.MatchString(regex, email)
}
func Register(c *gin.Context) {
	db := common.GetDB()
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	log.Println(username, email, password)

	if res, matchErr := validEmail(email); res == false || matchErr != nil {
		if matchErr != nil {
			response.Response(c, http.StatusInternalServerError, 500, nil, "Email matching fail")
		}
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Email invalid")
		return
	}
	if isEmailExisted(db, email) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Register failed,email existed.")
		return
	}
	newUser := model.User{
		Email:    email,
		Password: password,
	}
	db.Create(&newUser)
	response.Success(c, nil, "Successfully register")
}
func Login(c *gin.Context) {
	db := common.GetDB()
	email := c.PostForm("email")
	password := c.PostForm("password")
	log.Println(email, password)
	loginUser := getUserByEmail(db, email)
	if loginUser.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Login failed,email not exist")
		return
	}
	if loginUser.Password != password {
		response.Fail(c, nil, "Wrong password")
		return
	}
	token, tokenGenErr := common.ReleaseToken(loginUser)
	if tokenGenErr != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token generation failed")
		log.Println("token generation failed", tokenGenErr.Error())
		return
	}
	response.Success(c, gin.H{"token": token}, "Successfully login")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}}, "UserInfo request successfully")
}

//func GetFollowing
