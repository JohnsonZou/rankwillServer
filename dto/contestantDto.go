package dto

import "rankwillServer/model"

type QueryPageDto struct {
	Contestantname string  `json:"contestantname"`
	Username       string  `json:"username"`
	Rank           int     `json:"rank"`
	Oldrating      float64 `json:"oldrating"`
	Newrating      float64 `json:"newrating"`
	Deltarating    float64 `json:"deltarating"`
	Dataregion     string  `json:"dataregion"`
}

func ToQueryPageDto(f []model.Contestant) []QueryPageDto {
	var res []QueryPageDto
	for _, c := range f {
		res = append(res, QueryPageDto{
			Contestantname: c.Contestname,
			Username:       c.Username,
			Rank:           c.Rank,
			Oldrating:      c.Rating,
			Newrating:      c.PredictedRating,
			Deltarating:    c.PredictedRating - c.Rating,
			Dataregion:     c.Data_region,
		})
	}
	return res
}
func ToQueryByNameDto(c model.Contestant) QueryPageDto {
	return QueryPageDto{
		Contestantname: c.Contestname,
		Username:       c.Username,
		Rank:           c.Rank,
		Oldrating:      c.Rating,
		Newrating:      c.PredictedRating,
		Deltarating:    c.PredictedRating - c.Rating,
		Dataregion:     c.Data_region,
	}
}
