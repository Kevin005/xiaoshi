package response

import "xiaoshi/model"

/**
返回对象
*/
type RespUser struct {
	*RespModel
	Data interface{} `json:"data"`
}

type RespUserData struct {
	User model.Users `json:"user"`
}

