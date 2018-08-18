package response

import "xiaoshi/app/model"

/**
返回对象
 */
type RespFeedback struct {
	RespModel
	Data respFeedbackData `json:"data"`
}

type respFeedbackData struct {
	Feedback model.Feedbacks `json:"feedback"`
}

/**
返回数组
 */
type RespFeedbacks struct {
	RespModel
	Data interface{} `json:"data"`
}

type RespFeedbackDatas struct {
	Feedbacks []model.Feedbacks `json:"feedback"`
}
