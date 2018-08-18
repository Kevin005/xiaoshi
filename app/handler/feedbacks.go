package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"xiaoshi/app/model"
	"encoding/json"
	"xiaoshi/app/model/response"
)

func CreateFeedback(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respFeedback := response.RespFeedback{}
	if hadToken, _ := checkToken(db, token); hadToken {
		feedback := model.Feedbacks{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&feedback); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		defer r.Body.Close()
		if err := db.Save(&feedback).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respFeedback.Data.Feedback = feedback
		respFeedback.Message = "pass"
		respFeedback.Success = "0"
		respondJSON(w, http.StatusCreated, respFeedback)
	} else {
		respFeedback.Message = "reject"
		respFeedback.Success = "1"
		respondJSON(w, http.StatusInternalServerError, respFeedback)
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
		respondJSON(w, http.StatusCreated, respFeedbacks)
	} else {
		respFeedbacks.Message = "reject"
		respFeedbacks.Success = "1"
		respFeedbacks.Data = "user not found"
		respondJSON(w, http.StatusInternalServerError, respFeedbacks)
	}
}
