package dto

import "rankwillServer/model"

type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Email:    user.Email,
	}
}

//func ToQueryPageDto() {
//
//}
