package dto

import "rankwillServer/model"

type ContestDto struct {
	ContestName   string `json:"contestName"`
	UpdateTime    int64  `json:"updateTime"`
	ContestantNum int    `json:"contestantNum"`
}

func ToContestDto(c []model.Contest) []ContestDto {
	var res []ContestDto
	for _, v := range c {
		res = append(res, ContestDto{
			ContestName:   v.TitleSlug,
			UpdateTime:    v.StartTime,
			ContestantNum: v.ContestantNum,
		})
	}
	return res
}
