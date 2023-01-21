package models

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type CreateUserResponse struct {
	Id int `json:"id"`
}

type MakeFriendsRequest struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

type MakeFriendsResponse struct {
	Message string `json:"msg"`
}

type DeleteUserRequest struct {
	TargetId int `json:"target_id"`
}

type DeleteUserResponse struct {
	UserName string `json:"name"`
}

type GetFriendsResponse struct {
	Friends []string `json:"result"`
}

type UpdateUserResponse struct {
	Message string `json:"msg"`
}
