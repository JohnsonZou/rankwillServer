package dto

import "rankwillServer/model"

type UserDto struct {
	Email string `json:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Email: user.Email,
	}
}

//func ToQueryPageDto() {
//
//}
