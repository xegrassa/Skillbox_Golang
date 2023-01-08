package models

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type MakeFriendsRequest struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

type DeleteUserRequest struct {
	TargetId int `json:"target_id"`
}
