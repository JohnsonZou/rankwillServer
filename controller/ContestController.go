package controller

import (
	"github.com/gin-gonic/gin"
	"rankwillServer/common"
	"rankwillServer/dto"
	"rankwillServer/model"
	"rankwillServer/response"
	"sort"
	"strconv"
)

type allContest []model.Contest

func (a allContest) Len() int {
	return len(a)
}
func (a allContest) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a allContest) Less(i, j int) bool {
	return a[i].StartTime > a[j].StartTime
}
func GetContest(c *gin.Context) {
	page, err := strconv.Atoi(c.PostForm("page"))
	if page <= 0 || err != nil {
		response.Fail(c, nil, "page err!")
		return
	}
	var tot []model.Contest
	db := common.GetDB()
	_ = db.Find(&tot)
	sort.Sort(allContest(tot))
	total := len(tot)

	const pagesize = 10

	if ((page - 1) * pagesize) > total {
		response.Success(c, gin.H{"totnum": total, "data": nil}, "there is no such page")
	} else {
		response.Success(c, gin.H{"totnum": total, "data": dto.ToContestDto(tot[(page-1)*pagesize : common.Minint(page*pagesize, total)])}, "query successfully")
	}

}
