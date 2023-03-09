package dto

import "rankwillServer/model"

type FollowDto struct {
	Lcusername string `json:"lcusername"`
}

func ToFollowDto(f []model.Following) []FollowDto {
	var res []FollowDto
	for _, k := range f {
		res = append(res, FollowDto{
			Lcusername: k.Lcusername,
		})
	}
	return res
}
