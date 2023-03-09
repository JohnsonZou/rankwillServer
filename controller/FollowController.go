package controller

import (
	"github.com/gin-gonic/gin"
	"rankwillServer/common"
	"rankwillServer/dto"
	"rankwillServer/model"
	"rankwillServer/response"
)

func Follow(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	lcusername := c.PostForm("username")
	if lcusername == "" {
		response.Fail(c, nil, "Empty leetcode username")
		return
	}
	db := common.GetDB()
	if isFollowExisted(db, email, lcusername) {
		response.Fail(c, nil, "Duplicated follow")
		return
	}
	f := model.Following{
		Email:      email,
		Lcusername: lcusername,
	}
	db.Create(&f)
	response.Success(c, nil, "Successfully follow")
}
func Unfollow(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	lcusername := c.PostForm("username")
	if lcusername == "" {
		response.Fail(c, nil, "Empty leetcode username")
		return
	}
	db := common.GetDB()
	if isFollowExisted(db, email, lcusername) {
		db.Where("email=?", email).Where("lcusername=?", lcusername).Delete(&model.Following{})
		response.Success(c, nil, "Successfully unfollow")
		return
	}
	response.Fail(c, nil, "Leetcode user not exist")
}
func GetFollowList(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	db := common.GetDB()
	var fol []model.Following
	db.Where("email=?", email).Find(&fol)
	response.Success(c, gin.H{"result": dto.ToFollowDto(fol)}, "Successfully get followlist")

}
func GetFollowing(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	contestname := c.PostForm("contestname")
	if contestname == "" {
		response.Fail(c, nil, "Empty contest name")
		return
	}
	db := common.GetDB()
	var res []model.Contestant
	var fol []model.Following
	db.Where("email=?", email).Find(&fol)
	for _, v := range fol {
		curname := v.Lcusername
		db.Where("contestname=?", contestname).Where("username=?", curname).Find(&res)
	}

	response.Success(c, gin.H{"result": dto.ToQueryPageDto(res)}, "Successfully get following")
}
