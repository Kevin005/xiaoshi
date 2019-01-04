package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/conf"
	"xiaoshi/model"
	"xiaoshi/model/response"
)

func CreateFeedback(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respFeedback := response.RespFeedback{}
	if hadToken, _ := checkToken(db, token); hadToken {
		feedback := model.Feedbacks{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&feedback); err != nil {
			respondError(w, conf.STATUS_BAD_REQUEST, err.Error())
		}
		defer r.Body.Close()
		if err := db.Save(&feedback).Error; err != nil {
			respondError(w, conf.STATUS_INTERNAL_SERVER_ERROR, err.Error())
			return
		}
		respFeedback.Data.Feedback = feedback
		respFeedback.Message = "pass"
		respFeedback.Success = "0"
		respondJSON(w, conf.STATUS_CREATED, respFeedback)
	} else {
		respFeedback.Message = "reject"
		respFeedback.Success = "1"
		respondJSON(w, conf.STATUS_INTERNAL_SERVER_ERROR, respFeedback)
	}
}

func GetAllFeedback(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respFeedbacks := response.RespFeedbacks{}
	if hadToken, _ := checkToken(db, token); hadToken {
		feedbacs := []model.Feedbacks{}
		db.First(&feedbacs, "token = ?", token)
		respFeedbacksData := response.RespFeedbackDatas{}
		respFeedbacksData.Feedbacks = feedbacs
		respFeedbacks.Data = respFeedbacksData
		respFeedbacks.Message = "pass"
		respFeedbacks.Success = "0"
		respondJSON(w, conf.STATUS_CREATED, respFeedbacks)
	} else {
		respFeedbacks.Message = "reject"
		respFeedbacks.Success = "1"
		respFeedbacks.Data = "user not found"
		respondJSON(w, conf.STATUS_CREATED, respFeedbacks)
	}
}
