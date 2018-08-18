package handler

import (
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"xiaoshi/app/model"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
	//todo 可以统一返回json {"success":0}格式
}

func respondReject() {
	//todo 可以统一返回json {"success":1}格式
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
	//todo 可以统一返回json {"success":0}格式
}

func checkToken(db *gorm.DB, rToken string) (bool, model.Users) {
	user := model.Users{}
	db.First(&user, "token = ?", rToken)
	if user.Token != "" {
		return true, user
	}
	return false, user
}

func checkPhoneNumber(db *gorm.DB, phoneNum string) (bool, model.Users) {
	user := model.Users{}
	db.First(&user, "phone_number = ?", phoneNum)
	if user.PhoneNumber != "" {
		return true, user
	}
	return false, user
}
